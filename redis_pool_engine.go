package main

import (
	"time"

	"github.com/garyburd/redigo/redis"
)

var (
	pool *redis.Pool
)

func init() {
	config := LoadConfig().Cache.Redis

	if config.Pool.Enable {
		pool = &redis.Pool{
			Dial: func() (redis.Conn, error) {
				c, err := redis.Dial("tcp", config.Host+config.Port)
				if err != nil {
					return nil, err
				}
				return c, err
			},
			MaxIdle:     config.Pool.MaxIdle,
			IdleTimeout: 240 * time.Second,
		}
	}
}

type RedisPoolEngine struct {
}

func NewRedisPoolEngine() *RedisPoolEngine {
	return &RedisPoolEngine{}
}

func (engine *RedisPoolEngine) Exists(key string) bool {
	c := pool.Get()
	defer c.Close()

	exists, err := redis.Bool(c.Do("EXISTS", key))
	if err != nil {
		return false
	}
	return exists
}

func (engine *RedisPoolEngine) Write(key string, data []byte) (err error) {
	c := pool.Get()
	defer c.Close()

	_, err = c.Do("SET", key, data)
	return err
}

func (engine *RedisPoolEngine) Read(key string) (data []byte, err error) {
	c := pool.Get()
	defer c.Close()

	data, err = redis.Bytes(c.Do("GET", key))
	return data, err
}
