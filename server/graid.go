package server

import (
	"net/http"

	"github.com/violetyk/graid/config"
	. "github.com/violetyk/graid/handler"
)

type Graid struct {
	server *http.Server
}

func NewGraid() *Graid {
	config := config.Load()

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
