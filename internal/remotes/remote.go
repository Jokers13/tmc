package remotes

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
	"github.com/web-of-things-open-source/tm-catalog-cli/internal/config"
	"github.com/web-of-things-open-source/tm-catalog-cli/internal/model"
)

const (
	KeyRemotes       = "remotes"
	KeyRemoteType    = "type"
	KeyRemoteUrl     = "url"
	KeyRemoteDefault = "default"

	RemoteTypeFile = "file"
)

type Config map[string]map[string]any

var ErrNoDefault = errors.New("no default remote config found")
var ErrRemoteNotFound = errors.New("named remote not found")
var ErrRemoteExists = errors.New("named remote already exists")
var SupportedTypes = []string{RemoteTypeFile}

type Remote interface {
	// Push writes the Thing Model file into the path under root that corresponds to id.
	// Returns ErrTMExists if the same file is already stored with a different timestamp
	Push(id model.TMID, raw []byte) error
	Fetch(id model.TMID) ([]byte, error)
	CreateToC() error
	List(filter string) (model.TOC, error)
	Versions(name string) (model.TOCEntry, error)
}

// Get returns the Remote built from config with the given name
// Empty name returns the default remote
func Get(name string) (Remote, error) {
	remotes, err := ReadConfig()
	if err != nil {
		return nil, err
	}
	rc, ok := remotes[name]
	if name == "" {
		if len(remotes) == 1 {
			for _, v := range remotes {
				rc = v
			}
		} else {
			found := false
			for _, v := range remotes {
				if def, ok := v[KeyRemoteDefault]; ok {
					if d, ok := def.(bool); ok && d {
						rc = v
						found = true
						break
					}
				}
			}
			if !found {
				return nil, ErrNoDefault
			}
		}
	} else {
		if !ok {
			return nil, ErrRemoteNotFound
		}
	}

	switch t := rc[KeyRemoteType]; t {
	case RemoteTypeFile:
		return NewFileRemote(rc)
	default:
		return nil, fmt.Errorf("unsupported remote type: %v. Supported types are %v", t, SupportedTypes)
	}

}

func ReadConfig() (Config, error) {
	remotesConfig := viper.Get(KeyRemotes)
	remotes, ok := remotesConfig.(map[string]any)
	if !ok {
		return nil, fmt.Errorf("invalid remotes contig")
	}
	cp := map[string]map[string]any{}
	for k, v := range remotes {
		if cfg, ok := v.(map[string]any); ok {
			cp[k] = cfg
		} else {
			return nil, fmt.Errorf("invalid remote config: %s", k)
		}
	}
	return cp, nil
}

func SetDefault(name string) error {
	conf, err := ReadConfig()
	if err != nil {
		return err
	}
	if _, ok := conf[name]; !ok {
		return ErrRemoteNotFound
	}
	for n, rc := range conf {
		if n == name {
			rc[KeyRemoteDefault] = true
		} else {
			delete(rc, KeyRemoteDefault)
		}
	}
	return saveConfig(conf)
}
func Remove(name string) error {
	conf, err := ReadConfig()
	if err != nil {
		return err
	}
	if _, ok := conf[name]; !ok {
		return ErrRemoteNotFound
	}
	delete(conf, name)
	return saveConfig(conf)
}

func Add(name, typ, confStr string, confFile []byte) error {
	_, err := Get(name)
	if err == nil || !errors.Is(err, ErrRemoteNotFound) {
		return ErrRemoteExists
	}

	var rc map[string]any
	switch typ {
	case RemoteTypeFile:
		rc, err = createFileRemoteConfig(confStr, confFile)
		if err != nil {
			return err
		}
	default:
		return fmt.Errorf("unsupported remote type: %v. Supported types are %v", typ, SupportedTypes)
	}

	conf, err := ReadConfig()
	if err != nil {
		return err
	}

	conf[name] = rc

	return saveConfig(conf)
}

func saveConfig(conf Config) error {
	dc := 0
	for _, rc := range conf {
		d := rc[KeyRemoteDefault]
		if b, ok := d.(bool); ok && b {
			dc++
		}
	}
	if dc > 1 {
		return fmt.Errorf("too many default remotes. can accept at most one")
	}

	viper.Set(KeyRemotes, conf)
	configFile := viper.ConfigFileUsed()
	if configFile == "" {
		configFile = filepath.Join(config.DefaultConfigDir, "config.json")
	}
	err := os.MkdirAll(config.DefaultConfigDir, 0700)
	if err != nil {
		return err
	}
	return viper.WriteConfigAs(configFile)
}
