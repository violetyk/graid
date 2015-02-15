package cache

import "log"

type FileEngine struct {
}

func NewFileEngine() *FileEngine {
	return &FileEngine{}
}

func (e *FileEngine) Write() {
	log.Println("write cache to file")
}
