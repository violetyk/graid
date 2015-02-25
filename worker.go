package main

import (
	"bytes"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
)

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

// http://localhost:8080/xx/yy/zz/hogehoge.png:e?hoge=fuga&k=v#f
// http://localhost:8080/hogehoge.png
// http://localhost:8080/path/to/hogehoge.png:w50:w100
// http://localhost:8080/path/to/hogehoge.png:c100,200,10,50
// http://localhost:8080/http://example.com/hogehoge.png
// http://localhost:8080/http://example.com/hogehoge.png:c100,200,10,50

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
		response, err := http.Get(worker.Query.SourceUrl)
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
	io.Copy(w, dst)
}
