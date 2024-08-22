package commands

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"slices"
	"strings"

	"github.com/Masterminds/semver/v3"
	"github.com/buger/jsonparser"
	"github.com/wot-oss/tmc/internal/model"
	"github.com/wot-oss/tmc/internal/repos"
	"github.com/wot-oss/tmc/internal/utils"
)

func FetchByTMIDOrName(ctx context.Context, spec model.RepoSpec, idOrName string, restoreId bool) (string, []byte, error, []*repos.RepoAccessError) {
	tmid, fn, err := model.ParseAsTMIDOrFetchName(idOrName)
	if err != nil {
		return "", nil, err, nil
	}
	if tmid != nil {
		return FetchByTMID(ctx, spec, idOrName, restoreId)
	}
	return FetchByName(ctx, spec, *fn, restoreId)
}

func FetchByTMID(ctx context.Context, spec model.RepoSpec, tmid string, restoreId bool) (string, []byte, error, []*repos.RepoAccessError) {
	rs, err := repos.GetUnion(spec)
	if err != nil {
		return "", nil, err, nil
	}

	fetch, bytes, err, accessErrors := rs.Fetch(ctx, tmid)
	if err == nil && restoreId {
		bytes = restoreExternalId(bytes)
	}
	return fetch, bytes, err, accessErrors
}

func restoreExternalId(raw []byte) []byte {
	linksValue, dataType, _, err := jsonparser.Get(raw, "links")
	if err != nil && dataType != jsonparser.NotExist {
		return raw
	}

	if dataType != jsonparser.Array {
		return raw
	}

	var originalId string
	var linksArray []map[string]any

	err = json.Unmarshal(linksValue, &linksArray)
	if err != nil {
		slog.Default().Error("error unmarshalling links", "error", err)
		return raw
	}
	var newLinks []map[string]any
	for _, eLink := range linksArray {
		rel, relOk := eLink["rel"]
		href := utils.JsGetString(eLink, "href")
		if relOk && rel == "original" && href != nil {
			originalId = *href
		} else {
			newLinks = append(newLinks, eLink)
		}
	}
	if len(linksArray) != len(newLinks) { // original id found
		var withLinks []byte
		if len(newLinks) > 0 {
			linksBytes, err := json.Marshal(newLinks)
			if err != nil {
				slog.Default().Error("unexpected marshal error", "error", err)
				return raw
			}
			withLinks, err = jsonparser.Set(raw, linksBytes, "links")
			if err != nil {
				slog.Default().Error("unexpected json set value error", "error", err)
				return raw
			}
		} else {
			withLinks = jsonparser.Delete(raw, "links")
		}
		idBytes, _ := json.Marshal(originalId)

		withId, err := jsonparser.Set(withLinks, idBytes, "id")
		if err != nil {
			slog.Default().Error("unexpected json set value error", "error", err)
			return raw
		}
		return withId
	}

	return raw

}

func FetchByName(ctx context.Context, spec model.RepoSpec, fn model.FetchName, restoreId bool) (string, []byte, error, []*repos.RepoAccessError) {
	log := slog.Default()
	id, foundIn, err, errs := ResolveFetchName(ctx, spec, fn)
	if err != nil {
		return "", nil, err, errs
	}

	log.Debug(fmt.Sprintf("fetching %v from %s", id, foundIn))
	tmid, bytes, err, _ := FetchByTMID(ctx, foundIn, id, restoreId)
	return tmid, bytes, err, errs
}

func ResolveFetchName(ctx context.Context, spec model.RepoSpec, fn model.FetchName) (string, model.RepoSpec, error, []*repos.RepoAccessError) {
	res, err, errs := NewVersionsCommand().ListVersions(ctx, spec, fn.Name)
	if err != nil {
		return "", model.RepoSpec{}, err, errs
	}
	versions := make([]model.FoundVersion, len(res))
	copy(versions, res)

	var id string
	var foundIn model.RepoSpec
	// Just the name specified: fetch most recent
	if len(fn.Semver) == 0 {
		id, foundIn, err = findMostRecentVersion(versions)
		if err != nil {
			return id, foundIn, err, errs
		}
	} else {
		if _, err := semver.NewVersion(fn.Semver); err == nil {
			id, foundIn, err = findMostRecentMatchingVersion(versions, fn.Semver)
			if err != nil {
				return id, foundIn, err, errs
			}
		} else {
			return id, foundIn, err, errs
		}
	}
	return id, foundIn, err, errs
}

func findMostRecentVersion(versions []model.FoundVersion) (string, model.RepoSpec, error) {
	log := slog.Default()
	if len(versions) == 0 {
		err := fmt.Errorf("%w: no versions found", model.ErrTMNameNotFound)
		log.Error(err.Error())
		return "", model.EmptySpec, err
	}

	sortFoundVersionsDesc(versions)

	v := versions[0]
	return v.TMID, model.NewSpecFromFoundSource(v.FoundIn), nil
}

func findMostRecentMatchingVersion(versions []model.FoundVersion, ver string) (id string, source model.RepoSpec, err error) {
	log := slog.Default()
	ver, _ = strings.CutPrefix(ver, "v")

	// figure out how to match versions with ver
	var matcher func(*semver.Version) bool
	dots := strings.Count(ver, ".")
	if dots == 2 { // ver contains major.minor.patch
		sv := semver.MustParse(ver)
		matcher = sv.Equal
	} else { // at least one semver part is missing in ver
		c, err := semver.NewConstraint(fmt.Sprintf("~%s", ver))
		if err != nil {
			log.Error("couldn't parse semver constraint", "error", err)
			return "", model.EmptySpec, err
		}
		matcher = c.Check
	}

	// delete versions not matching ver from the list
	versions = slices.DeleteFunc(versions, func(version model.FoundVersion) bool {
		semVersion, err := semver.NewVersion(version.Version.Model)
		if err != nil {
			log.Error(err.Error())
			return false
		}
		matches := matcher(semVersion)
		return !matches
	})

	// see if anything remained
	if len(versions) == 0 {
		err := fmt.Errorf("%w: no version %s found", model.ErrTMNotFound, ver)
		log.Error(err.Error())
		return "", model.EmptySpec, err
	}

	// sort the remaining by semver then timestamp in descending order
	sortFoundVersionsDesc(versions)

	// and here's our winner
	v := versions[0]
	return v.TMID, model.NewSpecFromFoundSource(v.FoundIn), nil
}

// sortFoundVersionsDesc sorts by semver then timestamp in descending order, ie. from newest to oldest
func sortFoundVersionsDesc(versions []model.FoundVersion) {
	slices.SortStableFunc(versions, func(a, b model.FoundVersion) int {
		aid, _ := model.ParseTMID(a.TMID)
		bid, _ := model.ParseTMID(b.TMID)
		return -aid.Version.Compare(bid.Version)
	})
}
