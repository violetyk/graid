package handler

import (
	"log"
	"net/http"
	"runtime"

	"github.com/violetyk/graid/config"
)

type job struct {
	response http.ResponseWriter
	request  *http.Request
}

type jobs chan *job

type ImageHandler struct {
	jobs jobs
}

func NewImageHandler() *ImageHandler {

	config := config.Load()

	jobs := make(jobs, config.Server.WorkerPoolSize)
	for i := 1; i <= config.Server.WorkerPoolSize; i++ {
		go worker(i, jobs)
	}

	return &ImageHandler{
		jobs: jobs,
	}
}

func (handler *ImageHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/favicon.ico" {
		return
	}
	handler.jobs <- &job{response: w, request: r}
}

func worker(id int, jobs jobs) {
	for job := range jobs {
		log.Println("worker", id, "GR", runtime.NumGoroutine(), "processing job", job.request.URL.String())
	}
}
