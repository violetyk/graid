package main

import (
	"log"
	"net/http"
	"runtime"

	"github.com/violetyk/graid/config"
)

type Job struct {
	Response http.ResponseWriter
	Request  *http.Request
}

type Jobs chan *Job

type ImageHandler struct {
	Jobs Jobs
}

func NewImageHandler() *ImageHandler {

	config := config.Load()

	jobs := make(Jobs, config.Server.WorkerPoolSize)
	for i := 1; i <= config.Server.WorkerPoolSize; i++ {
		go worker(i, jobs)
	}

	return &ImageHandler{
		Jobs: jobs,
	}
}

func (handler *ImageHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/favicon.ico" {
		return
	}
	handler.Jobs <- &Job{Response: w, Request: r}
}

func worker(id int, jobs Jobs) {
	for job := range jobs {
		log.Println("worker", id, "GR", runtime.NumGoroutine(), "processing job", job.Request.URL.String())
	}
}
