package handler

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
	Params           map[string]string
}

func NewQuery() *Query {
	return &Query{}
}

func (query *Query) Clear() {
	query.Raw = ""
	query.SourceUrl = ""
	query.IsExternalSource = false
	query.Params = make(map[string]string)
}

var regexp_protocol *regexputil.RegexpUtil = regexputil.MustCompile(`^/(?P<protocol>http|https):/`)
var regexp_params *regexputil.RegexpUtil = regexputil.MustCompile(`(?P<operator>^[a-z]+)(?P<value>[0-9,]+$)`)

func (query *Query) Parse(urlString string) bool {

	query.Clear()
	query.Raw = urlString

	u, err := url.Parse(query.Raw)
	if err != nil {
		return false
	}

	s := strings.Split(u.Path, ":")

	match_protocol := regexp_protocol.FindStringSubmatchMap(u.Path)
	if len(match_protocol) > 0 {
		query.IsExternalSource = true
		query.SourceUrl = match_protocol["protocol"] + ":/" + s[1]
	} else {
		query.IsExternalSource = false
		query.SourceUrl = config.Load().Origin.Url + s[0]
	}

	var match_param map[string]string = map[string]string{}
	for _, v := range s {
		match_param = regexp_params.FindStringSubmatchMap(v)
		if len(match_param) > 0 {
			query.Params[match_param["operator"]] = match_param["value"]
		}
	}

	return true
}
