package main

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	http.HandleFunc("/", IndexHandler)
	http.ListenAndServe(":8080", nil)
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	// http://localhost:8080/xx/yy/zz/hogehoge.png:e?hoge=fuga&k=v#f

	// http://localhost:8080/hogehoge.png
	// http://localhost:8080/path/to/hogehoge.png:w50:w100
	// http://localhost:8080/path/to/hogehoge.png:c100,200,10,50
	// http://localhost:8080/example.com/hogehoge.png:e
	// http://localhost:8080/example.com/hogehoge.png:e:c100,200,10,50

	u, err := url.Parse(r.URL.String())
	if err != nil {
		panic(err)
	}

	fmt.Printf("%v\n", u.Path)

	s := strings.Split(u.Path, ":")
	fmt.Printf("%v\n", s)

	// include e? s[0] is external
	// include ? external

	// fmt.Printf("%v\n", u.Fragment)
	// fmt.Printf("%v\n", u.RawQuery)
	// m, _ := url.ParseQuery(u.RawQuery)
	// fmt.Printf("%v\n", m)
	// fmt.Printf("%v\n", m["k"][0])

	fmt.Fprint(w, "Hello world")
}
