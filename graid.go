package main

import (
	"net/http"
)

type Graid struct {
	server *http.Server
}

func NewGraid() *Graid {
	config := LoadConfig()

	serveMux := http.NewServeMux()
	serveMux.Handle("/", NewImageHandler())

	server := &http.Server{
		Addr:    config.Server.Port,
		Handler: serveMux,
	}

	return &Graid{
		server: server,
	}
}

func (graid *Graid) Start() error {
	return graid.server.ListenAndServe()
}
