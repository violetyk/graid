package main

import (
	"bytes"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
)

var httpClient *http.Client

func init() {
	httpClient = new(http.Client)
}

type Worker struct {
	Id        int
	Query     *Query
	Cache     *Cache
	Processor *Processor
	useCache  bool
}

func NewWorker(id int) *Worker {
	w := &Worker{
		Id:        id,
		Query:     NewQuery(),
		Processor: NewProcessor(),
	}

	config := LoadConfig()
	if config.Cache.Enable {
		w.useCache = true
		w.Cache = NewCache()
	}

	return w
}

func (worker *Worker) Execute(w http.ResponseWriter, r *http.Request) {

	// parse query
	if !worker.Query.Parse(r.URL.String()) {
		errors.New("TODO: return 404")
	}

	var data []byte
	var err error

	beCached := false

	// ready image data
	if worker.useCache && worker.Cache.Exists(worker.Query) {
		data, err = worker.Cache.Read(worker.Query)
	} else {
		response, err := httpClient.Get(worker.Query.SourceUrl)
		if err == nil {
			data, err = ioutil.ReadAll(response.Body)
		}
		defer response.Body.Close()
		beCached = true
	}
	if err != nil {
		errors.New("TODO: return 404")
	}

	// image
	src, err := NewImage(data)
	if err != nil {
		errors.New("TODO: return 404")
	}

	// process
	dst := new(bytes.Buffer)
	worker.Processor.Execute(src, dst, worker.Query)

	// cache
	if worker.useCache && beCached {
		go worker.Cache.Write(worker.Query, dst.Bytes())
	}

	// response
	w.Header().Set("Content-Type", http.DetectContentType(dst.Bytes()))
	w.Header().Set("Content-Length", strconv.Itoa(dst.Len()))
	io.Copy(w, dst)
}
