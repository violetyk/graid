package main

import "net/url"

type RedisEngineAdapter struct {
}

func NewRedisEngineAdapter() *RedisEngineAdapter {
	return &RedisEngineAdapter{}
}

func (adapter *RedisEngineAdapter) CacheKey(query *Query) string {
	return url.QueryEscape(query.SourceUrl) + query.Stringify()
}
