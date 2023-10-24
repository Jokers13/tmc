package remotes

import (
	"fmt"
	"github.com/web-of-things-open-source/tm-catalog-cli/internal/model"
	"log/slog"
	url2 "net/url"
	"os"
	"path/filepath"
	"strings"
)

type FileRemote struct {
	root string
}

func NewFileRemote(config map[string]any) (*FileRemote, error) {
	urlString := config["url"].(string)
	rootUrl, err := url2.Parse(urlString)
	if err != nil {
		slog.Default().Error("could not parse root URL for file remote", "url", urlString, "error", err)
		return nil, fmt.Errorf("could not parse root URL %s for file remote: %w", urlString, err)
	}
	if rootUrl.Scheme != "file" {
		slog.Default().Error("root URL for file remote must begin with file:", "url", urlString)
		return nil, fmt.Errorf("root URL for file remote must begin with file: %s", urlString)
	}
	rootPath := rootUrl.Opaque
	if strings.HasPrefix(rootPath, "~") {
		home, err := os.UserHomeDir()
		if err != nil {
			slog.Default().Error("cannot expand user home directory", "error", err)
			return nil, fmt.Errorf("cannot expand user home directory: %w", err)
		}
		rootPath = strings.Replace(rootPath, "~", home, 1)
	}
	return &FileRemote{
		root: rootPath,
	}, nil
}

func (f *FileRemote) Push(model *model.ThingModel, filename string, raw []byte) error {
	manufacturer := model.Manufacturer.Name
	mpn := model.Mpn
	fullPath := filepath.Join(f.root, manufacturer, mpn, filename)
	dir := filepath.Dir(fullPath)
	err := os.MkdirAll(dir, os.ModePerm) //fixme: review permissions
	if err != nil {
		return fmt.Errorf("could not create directory %s: %w", dir, err)
	}
	err = os.WriteFile(fullPath, raw, os.ModePerm) //fixme: review permissions
	if err != nil {
		return fmt.Errorf("could not write TM to catalog: %v", err)
	}

	return nil
}
