package cache

import (
	"io"
	"os"
	"path/filepath"
	"strings"
)

type FileEngine struct {
	Path string
}

func NewFileEngine(path string) *FileEngine {
	if err := os.MkdirAll(path, os.ModePerm); err != nil {
		panic(err)
	}
	return &FileEngine{Path: path}
}

func (engine *FileEngine) Exists(keys []string) bool {
	if _, err := os.Stat(engine.getCachePath(keys)); err != nil {
		return false
	}
	return true
}

func (engine *FileEngine) Write(reader io.Reader, keys []string) (err error) {
	cache_path := engine.getCachePath(keys)
	if err := os.MkdirAll(filepath.Dir(cache_path), os.ModePerm); err != nil {
		return err
	}
	file, err := os.Create(cache_path)

	if err != nil {
		return err
	}
	defer file.Close()
	_, err = io.Copy(file, reader)
	return err
}

func (engine *FileEngine) getCachePath(keys []string) string {
	return engine.Path + string(os.PathSeparator) + strings.Join(keys, string(os.PathSeparator))
}
