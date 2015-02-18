package cache

import (
	"errors"
	"io"
	"strings"

	"github.com/violetyk/graid/config"
)

type CacheEngine interface {
	Exists(keys []string) bool
	Write(reader io.Reader, keys []string) error
	// Read()
	// Delete()
}

type Cache struct {
	engine CacheEngine
}

func NewCache() *Cache {
	config := config.Load()

	var e CacheEngine
	switch strings.ToLower(config.Cache.Engine) {
	case "file":
		e = NewFileEngine(config.Cache.File.Path)
		// case "redis":
		// e = NewRedisEngine()
	}

	_, ok := e.(CacheEngine)
	if !ok {
		panic(errors.New(`cache engine is not available.`))
	}

	return &Cache{engine: e}
}

func (c *Cache) Write(reader io.Reader, keys []string) error {
	return c.engine.Write(reader, keys)
}

func (c *Cache) Exists(keys []string) bool {
	return c.engine.Exists(keys)
}
