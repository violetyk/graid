package main

import (
	"net/url"
	"strings"

	"github.com/violetyk/graid/config"
	"github.com/violetyk/graid/util"
)

type Query struct {
	Raw              string
	SourceUrl        string
	IsExternalSource bool
	ResizeWidth      int
	ResizeHeight     int
	CropX            int
	CropY            int
	CropWidth        int
	CropHeight       int
}

func NewQuery() *Query {
	return &Query{}
}

func (query *Query) Clear() {
	query.Raw = ""
	query.SourceUrl = ""
	query.IsExternalSource = false
	query.ResizeWidth = 0
	query.ResizeHeight = 0
	query.CropX = 0
	query.CropY = 0
	query.CropWidth = 0
	query.CropHeight = 0
}

func (query *Query) Parse(urlString string) bool {
	query.Clear()

	query.Raw = urlString

	u, err := url.Parse(query.Raw)
	if err != nil {
		return false
	}

	var match map[string]string = map[string]string{}
	// var pattern RegexpUtil

	pattern, err := regexputil.Compile(`^/(?P<protocol>http|https)://`)
	if err != nil {
		return false
	}

	match = pattern.FindStringSubmatchMap(u.Path)

	s := strings.Split(u.Path, ":")
	// for k, v := range s {
	// switch v {
	// case m, _ := regexp.MatchString(`^\d+$`, v):
	// query.ResizeHeight = strconv.Atoi(v[1:])

	// }
	// }

	if len(match) > 0 {
		query.IsExternalSource = true
		query.SourceUrl = match["protocol"] + s[1]
	} else {
		query.IsExternalSource = false
		query.SourceUrl = config.Load().Origin.Url + s[0]
	}

	return true
}

// func main() {

// var s string = "http://localhost:8080/http://hoge.com/path/to/hogehoge.png:h50:w100"
// // var s string = "http://localhost:8080/path/to/hogehoge.png:w50:w100"
// // var s string = "http://localhost:8080/path/to/hogehoge.png"

// // sum := md5.Sum(([]byte)s)

// query := NewQuery()
// ok := query.Parse(s)
// if ok {
// pp.Print(query)
// }

// // pp.Print(len(s))
// }
