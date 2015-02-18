package handler

import (
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"net/url"

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

	var cache_keys []string
	if worker.useCache {
		cache_keys = []string{worker.stringQuerySourceUrl(), worker.stringQueryParams()}
		if worker.Cache.Exists(cache_keys) {
			log.Println("cache exists")
		}
	}

	response, err := http.Get(worker.Query.SourceUrl)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	if worker.useCache {
		worker.Cache.Write(response.Body, cache_keys)
	}

	io.Copy(w, response.Body)
}

func (worker *Worker) stringQuerySourceUrl() string {
	return url.QueryEscape(worker.Query.SourceUrl)
}

func (worker *Worker) stringQueryParams() string {
	length := len(worker.Query.Params)
	if length == 0 {
		return "default"
	}

	j, err := json.Marshal(worker.Query.Params)
	if err != nil {
		return "default"
	}

	h := sha1.New()
	io.WriteString(h, string(j))
	return hex.EncodeToString(h.Sum(nil))
}
