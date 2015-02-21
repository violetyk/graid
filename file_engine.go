package main

import (
	"os"
	"path/filepath"
)

type FileEngine struct {
}

func NewFileEngine() *FileEngine {
	return &FileEngine{}
}

func (engine *FileEngine) Exists(path string) bool {
	if _, err := os.Stat(path); err != nil {
		return false
	}
	return true
}

func (engine *FileEngine) Write(path string, data []byte) (err error) {
	if err := os.MkdirAll(filepath.Dir(path), os.ModePerm); err != nil {
		return err
	}

	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(data)
	return err
}
