package config

import "github.com/BurntSushi/toml"

type config struct {
	Origin origin
}

type origin struct {
	Server string `toml:server`
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
