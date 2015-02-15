package cache

import "log"

type RedisEngine struct {
}

func NewRedisEngine() *RedisEngine {
	return &RedisEngine{}
}

func (e *RedisEngine) Write() {
	log.Println("write cache to redis")
}
