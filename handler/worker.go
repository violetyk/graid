package handler

import (
	"errors"
	"io"
	"net/http"

	"github.com/violetyk/graid/cache"
	"github.com/violetyk/graid/config"
)

type Worker struct {
	Id       int
	Query    *Query
	Cache    *cache.Cache
	useCache bool
}

func NewWorker(id int) *Worker {
	w := &Worker{
		Id:    id,
		Query: NewQuery(),
	}

	config := config.Load()
	if config.Cache.Enable {
		w.useCache = true
		w.Cache = cache.NewCache()
	}

	return w
}

func (worker *Worker) Execute(w http.ResponseWriter, r *http.Request) {

	// http://localhost:8080/xx/yy/zz/hogehoge.png:e?hoge=fuga&k=v#f
	// http://localhost:8080/hogehoge.png
	// http://localhost:8080/path/to/hogehoge.png:w50:w100
	// http://localhost:8080/path/to/hogehoge.png:c100,200,10,50
	// http://localhost:8080/http://example.com/hogehoge.png
	// http://localhost:8080/example.com/hogehoge.png:c100,200,10,50

	if !worker.Query.Parse(r.URL.String()) {
		errors.New("TODO: return 404")
	}

	response, err := http.Get(worker.Query.SourceUrl)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	io.Copy(w, response.Body)
}
