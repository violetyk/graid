package main

import (
	"errors"
	"io"
	"strings"
)

type CacheEngine interface {
	Exists(key string) bool
	Write(key string, reader io.Reader) error
	// Read()
	// Delete()
}

type CacheEngineAdapter interface {
	CacheKey(query *Query) string
}

type Cache struct {
	engine  CacheEngine
	adapter CacheEngineAdapter
}

func NewCache() *Cache {
	config := LoadConfig()

	var e CacheEngine
	var a CacheEngineAdapter

	switch strings.ToLower(config.Cache.Engine) {
	case "file":
		e = NewFileEngine()
		a = NewFileEngineAdapter()
		// case "redis":
		// e = NewRedisEngine()
	}

	var ok bool
	_, ok = e.(CacheEngine)
	if !ok {
		panic(errors.New(`cache engine is not available.`))
	}

	_, ok = a.(CacheEngineAdapter)
	if !ok {
		panic(errors.New(`cache engine adapter is not available.`))
	}

	return &Cache{engine: e, adapter: a}
}

func (cache *Cache) Write(query *Query, reader io.Reader) error {
	return cache.engine.Write(cache.adapter.CacheKey(query), reader)
}

func (cache *Cache) Exists(query *Query) bool {
	return cache.engine.Exists(cache.adapter.CacheKey(query))
}
