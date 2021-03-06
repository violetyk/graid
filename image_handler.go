package main

import (
	"net/http"
)

type WorkerPool chan *Worker

type ImageHandler struct {
	workerPool WorkerPool
}

func NewImageHandler() *ImageHandler {

	config := LoadConfig()

	workerPool := make(WorkerPool, config.Server.WorkerPoolSize)
	for i := 1; i <= config.Server.WorkerPoolSize; i++ {
		workerPool <- NewWorker(i)
	}

	return &ImageHandler{
		workerPool: workerPool,
	}
}

func (handler *ImageHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/favicon.ico" {
		return
	}
	worker := <-handler.workerPool
	worker.Execute(w, r)
	handler.workerPool <- worker
}
