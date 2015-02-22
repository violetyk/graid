package main

import "github.com/garyburd/redigo/redis"

type RedisEngine struct {
	Host string
	Port string
}

func NewRedisEngine(host, port string) *RedisEngine {
	return &RedisEngine{
		Host: host,
		Port: port,
	}
}

func (engine *RedisEngine) Exists(key string) bool {
	c, err := redis.Dial("tcp", engine.Host+engine.Port)
	if err != nil {
		panic(err)
		return false
	}
	defer c.Close()

	exists, err := redis.Bool(c.Do("EXISTS", key))
	if err != nil {
		return false
	}
	return exists
}

func (engine *RedisEngine) Write(key string, data []byte) (err error) {
	c, err := redis.Dial("tcp", engine.Host+engine.Port)
	if err != nil {
		return err
	}
	defer c.Close()

	_, err = c.Do("SET", key, data)
	return err
}

func (engine *RedisEngine) Read(key string) (data []byte, err error) {
	c, err := redis.Dial("tcp", engine.Host+engine.Port)
	if err != nil {
		return data, err
	}
	defer c.Close()

	data, err = redis.Bytes(c.Do("GET", key))
	return data, err
}
