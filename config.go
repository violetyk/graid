package main

import "github.com/BurntSushi/toml"

// singleton pattern

type config struct {
	Server server
	Cache  cache
	Origin origin
}

type server struct {
	Port             string `toml:"port"`
	WorkerPoolSize   int    `toml:"worker_pool_size"`
	IdleConnsPerHost int    `toml:"idle_conns_per_host"`
}

type cache struct {
	Enable bool       `toml:"enable"`
	Engine string     `toml:"engine"`
	File   cacheFile  `toml:"file"`
	Redis  cacheRedis `toml:"redis"`
}

type cacheFile struct {
	Path string `toml:"path"`
}

type cacheRedis struct {
	Host string `toml:"host"`
	Port string `toml:"port"`
	Pool cacheRedisPool
}

type cacheRedisPool struct {
	Enable  bool `toml:"enable"`
	MaxIdle int  `toml:"max_idle"`
}

type origin struct {
	Url string `toml:server`
}

var instance *config

func LoadConfig() *config {
	if instance == nil {
		var c config
		_, err := toml.DecodeFile("graid.toml", &c)
		if err != nil {
			panic(err)
		}
		instance = &c
	}
	return instance
}
