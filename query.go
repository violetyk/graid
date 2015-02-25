package main

import (
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"io"
	"net/url"
	"regexp"
	"strconv"
	"strings"
)

var SupportOperators map[string]string = map[string]string{
	"w":          "resize",
	"h":          "resize",
	"c":          "crop",
	"q":          "quality", // only jpeg
	"grayscale":  "grayscale",
	"sepia":      "sepia",
	"contrast":   "contrast",
	"brightness": "brightness",
	"saturation": "saturation",
}

type Query struct {
	Raw              string
	SourceUrl        string
	IsExternalSource bool
	params           map[string]string
}

func NewQuery() *Query {
	return &Query{}
}

func (query *Query) Clear() {
	query.Raw = ""
	query.SourceUrl = ""
	query.IsExternalSource = false
	query.params = make(map[string]string)
}

var regexp_protocol *RegexpUtil = &RegexpUtil{regexp.MustCompile(`^/(?P<protocol>http|https):/`)}
var regexp_params *RegexpUtil = &RegexpUtil{regexp.MustCompile(`(?P<operator>^[a-z]+)(?P<value>[0-9,-]+$)`)}

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
		query.SourceUrl = LoadConfig().Origin.Url + s[0]
	}

	match_param := map[string]string{}
	for _, v := range s {
		match_param = regexp_params.FindStringSubmatchMap(v)
		if len(match_param) > 0 {
			if _, ok := SupportOperators[match_param["operator"]]; ok {
				query.params[match_param["operator"]] = match_param["value"]
			}
		}
	}

	return true
}

func (query *Query) Stringify() string {
	if query.Count() == 0 {
		return "default"
	}

	j, err := json.Marshal(query.params)
	if err != nil {
		return "default"
	}

	h := sha1.New()
	io.WriteString(h, string(j))
	return hex.EncodeToString(h.Sum(nil))
}

func (query *Query) Count() int {
	return len(query.params)
}

func (query *Query) Has(operator string) bool {
	_, exists := query.params[operator]
	return exists
}

func (query *Query) GetInt(operator string) (value int) {
	if v, exists := query.params[operator]; exists {
		value, _ = strconv.Atoi(v)
	}
	return value
}

func (query *Query) GetIntArray(operator string) (values []int) {
	if v, exists := query.params[operator]; exists {
		for _, s := range strings.Split(v, ",") {
			i, _ := strconv.Atoi(s)
			values = append(values, i)
		}
	}
	return values
}
