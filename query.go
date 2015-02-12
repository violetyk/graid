package main

import (
	"net/url"
	"regexp"

	"github.com/k0kubun/pp"
)

type Query struct {
	Raw         string
	SourceUrl   string
	IsExternal  bool
	ResizeOrder resizeOrder
	CropOrder   cropOrder
}

type resizeOrder struct {
	width  int
	height int
}

type cropOrder struct {
	width  int
	height int
	x      int
	y      int
}

func NewQuery() *Query {
	return &Query{}
}

func (query *Query) Clear() {
}

func (query *Query) Parse(urlString string) bool {
	query.Clear()

	query.Raw = urlString

	u, err := url.Parse(query.Raw)
	if err != nil {
		return false
	}

	// pattern := regexp.MustCompile("^/(http|https)://")
	// match := pattern.MatchString(u.Path)
	// pattern := regexp.MustCompile(`^/(?P<protocol>http|https)://`)
	match := regexp.MustCompile(`(?P<protocol>wef)`).FindStringSubmatch(u.Path)
	match := regexp.MustCompile(`(?P<protocol>wef)`).FindStringSubmatch(u.Path)
	pp.Print(len(match))

	// s := strings.Split(u.Path, ":")

	// if match {
	// query.IsExternal = true
	// query.SourceUrl = (s[0] + ":" + s[1])[1:]
	// } else {
	// query.IsExternal = false
	// query.SourceUrl = config.Load().Origin.Url + s[0]
	// }

	// return true
}

func main() {

	// var s string = "http://localhost:8080/http://hoge.com/path/to/hogehoge.png:w50:w100"
	var s string = "http://localhost:8080/path/to/hogehoge.png:w50:w100"

	// sum := md5.Sum(([]byte)s)

	// query := NewQuery(s)
	query := NewQuery()
	ok := query.Parse(s)
	pp.Print(ok)
	if !ok {
		pp.Print(query)
	}

	// pp.Print(len(s))
}
