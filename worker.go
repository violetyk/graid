package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"runtime"

	"github.com/violetyk/graid/config"
)

type Worker struct {
	Id int
}

func NewWorker(id int) *Worker {
	return &Worker{
		Id: id,
	}
}

func (worker *Worker) Execute(w http.ResponseWriter, r *http.Request) {
	// http://localhost:8080/xx/yy/zz/hogehoge.png:e?hoge=fuga&k=v#f
	// http://localhost:8080/hogehoge.png
	// http://localhost:8080/path/to/hogehoge.png:w50:w100
	// http://localhost:8080/path/to/hogehoge.png:c100,200,10,50
	// http://localhost:8080/example.com/hogehoge.png:e
	// http://localhost:8080/example.com/hogehoge.png:e:c100,200,10,50

	log.Println(runtime.NumGoroutine())
	u, err := url.Parse(r.URL.String())
	if err != nil {
		panic(err)
	}

	fmt.Printf("%d : %v\n", worker.Id, u.Path)

	// s := strings.Split(u.Path, ":")
	// fmt.Printf("%v\n", s)

	// include e? s[0] is external
	// include ? external

	// fmt.Printf("%v\n", u.Fragment)
	// fmt.Printf("%v\n", u.RawQuery)
	// m, _ := url.ParseQuery(u.RawQuery)
	// fmt.Printf("%v\n", m)
	// fmt.Printf("%v\n", m["k"][0])

	config := config.Load()
	url := config.Origin.Url + u.Path
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	io.Copy(w, response.Body)
	// fmt.Fprint(w, "Hello world")
}
