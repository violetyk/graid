package cache

import (
	"errors"
	"strings"

	"github.com/violetyk/graid/config"
)

type CacheEngine interface {
	// Exists(source string) bool
	Write()
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
		// TODO: give config, setup file engine
		e = NewFileEngine()
	case "redis":
		e = NewRedisEngine()
	}

	_, ok := e.(CacheEngine)
	if !ok {
		panic(errors.New(`cache engine is not available.`))
	}

	return &Cache{engine: e}
}

func (c *Cache) Write() {
	c.engine.Write()
}
