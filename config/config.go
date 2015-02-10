package config

import "github.com/BurntSushi/toml"

type config struct {
	Server server
	Origin origin
}

type server struct {
	Port           string `toml:"port"`
	WorkerPoolSize int    `toml:"worker_pool_size"`
}

type origin struct {
	Url string `toml:server`
}

var instance *config

func Load() *config {
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
