package main

import (
	"bytes"
	"errors"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	// "github.com/violetyk/graid/config"
)

type Worker struct {
	Id       int
	Query    *Query
	Cache    *Cache
	useCache bool
}

func NewWorker(id int) *Worker {
	w := &Worker{
		Id:    id,
		Query: NewQuery(),
	}

	config := LoadConfig()
	if config.Cache.Enable {
		w.useCache = true
		w.Cache = NewCache()
	}

	return w
}

func (worker *Worker) Execute(w http.ResponseWriter, r *http.Request) {

	// http://localhost:8080/xx/yy/zz/hogehoge.png:e?hoge=fuga&k=v#f
	// http://localhost:8080/hogehoge.png
	// http://localhost:8080/path/to/hogehoge.png:w50:w100
	// http://localhost:8080/path/to/hogehoge.png:c100,200,10,50
	// http://localhost:8080/http://example.com/hogehoge.png
	// http://localhost:8080/http://example.com/hogehoge.png:c100,200,10,50

	if !worker.Query.Parse(r.URL.String()) {
		errors.New("TODO: return 404")
	}

	if worker.useCache {
		if worker.Cache.Exists(worker.Query) {
			log.Println("cache exists")
		}
	}

	response, err := http.Get(worker.Query.SourceUrl)
	if err != nil {
		errors.New("TODO: return 404")
	}
	defer response.Body.Close()

	if worker.useCache {
		data, err := ioutil.ReadAll(response.Body)
		if err == nil {
			worker.Cache.Write(worker.Query, data)
		}
		io.Copy(w, bytes.NewReader(data))
	}

}

// func (worker *Worker) stringQuerySourceUrl() string {
// return url.QueryEscape(worker.Query.SourceUrl)
// }
