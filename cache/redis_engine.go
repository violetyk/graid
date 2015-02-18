package cache

import "log"

type RedisEngine struct {
}

func NewRedisEngine() *RedisEngine {
	return &RedisEngine{}
}

func (e *RedisEngine) Exists(key string) bool {
	return false
}

func (e *RedisEngine) Write() {
	log.Println("write cache to redis")
}
