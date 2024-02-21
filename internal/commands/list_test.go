package commands

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/web-of-things-open-source/tm-catalog-cli/internal/model"
	"github.com/web-of-things-open-source/tm-catalog-cli/internal/remotes"
)

func TestListCommand_List(t *testing.T) {
	t.Run("merged", func(t *testing.T) {
		rm := remotes.NewMockRemoteManager(t)
		r1 := remotes.NewMockRemote(t)
		r2 := remotes.NewMockRemote(t)
		rm.On("All").Return([]remotes.Remote{r1, r2}, nil)
		r1.On("List", &model.SearchParams{Query: "omnicorp"}).Return(model.SearchResult{
			Entries: []model.FoundEntry{
				{
					Name:         "omnicorp/senseall",
					Manufacturer: model.SchemaManufacturer{Name: "omnicorp"},
					Mpn:          "senseall",
					Author:       model.SchemaAuthor{Name: "omnicorp"},
				},
				{
					Name:         "omnicorp/lightall",
					Manufacturer: model.SchemaManufacturer{Name: "omnicorp"},
					Mpn:          "lightall",
					Author:       model.SchemaAuthor{Name: "omnicorp"},
				},
			},
		}, nil)
		r2.On("List", &model.SearchParams{Query: "omnicorp"}).Return(model.SearchResult{
			Entries: []model.FoundEntry{
				{
					Name:         "omnicorp/senseall",
					Manufacturer: model.SchemaManufacturer{Name: "omnicorp"},
					Mpn:          "senseall",
					Author:       model.SchemaAuthor{Name: "omnicorp"},
				},
				{
					Name:         "omnicorp/actall",
					Manufacturer: model.SchemaManufacturer{Name: "omnicorp"},
					Mpn:          "actall",
					Author:       model.SchemaAuthor{Name: "omnicorp"},
				},
			},
		}, nil)

		c := NewListCommand(rm)
		res, err, _ := c.List(remotes.EmptySpec, &model.SearchParams{Query: "omnicorp"})

		assert.NoError(t, err)
		assert.Len(t, res.Entries, 3)
		assert.Equal(t, "omnicorp/actall", res.Entries[0].Name)
		assert.Equal(t, "omnicorp/lightall", res.Entries[1].Name)
		assert.Equal(t, "omnicorp/senseall", res.Entries[2].Name)
	})
	t.Run("one error", func(t *testing.T) {
		rm := remotes.NewMockRemoteManager(t)
		r1 := remotes.NewMockRemote(t)
		r2 := remotes.NewMockRemote(t)
		r2.On("Spec").Return(remotes.NewRemoteSpec("r2"))
		rm.On("All").Return([]remotes.Remote{r1, r2}, nil)
		r1.On("List", &model.SearchParams{Query: "omnicorp"}).Return(model.SearchResult{
			Entries: []model.FoundEntry{
				{
					Name:         "omnicorp/senseall",
					Manufacturer: model.SchemaManufacturer{Name: "omnicorp"},
					Mpn:          "senseall",
					Author:       model.SchemaAuthor{Name: "omnicorp"},
				},
				{
					Name:         "omnicorp/lightall",
					Manufacturer: model.SchemaManufacturer{Name: "omnicorp"},
					Mpn:          "lightall",
					Author:       model.SchemaAuthor{Name: "omnicorp"},
				},
			},
		}, nil)
		r2.On("List", &model.SearchParams{Query: "omnicorp"}).Return(model.SearchResult{}, errors.New("unexpected error"))

		c := NewListCommand(rm)
		res, err, errs := c.List(remotes.EmptySpec, &model.SearchParams{Query: "omnicorp"})

		assert.NoError(t, err)
		if assert.Len(t, errs, 1) {
			assert.ErrorContains(t, errs[0], "unexpected error")
		}
		assert.Len(t, res.Entries, 2)
		assert.Equal(t, "omnicorp/lightall", res.Entries[0].Name)
		assert.Equal(t, "omnicorp/senseall", res.Entries[1].Name)
	})

}
