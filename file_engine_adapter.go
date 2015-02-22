package main

import (
	"os"
	"strings"
)

var pathSeparator string = string(os.PathSeparator)
var pathReplacer *strings.Replacer = strings.NewReplacer(
	"http:/", "",
	"https:/", "",
	"/", pathSeparator,
	"..", "",
)

type FileEngineAdapter struct {
}

func NewFileEngineAdapter() *FileEngineAdapter {
	return &FileEngineAdapter{}
}

func (adapter *FileEngineAdapter) CacheKey(query *Query) string {
	return LoadConfig().Cache.File.Path + pathSeparator + pathReplacer.Replace(query.SourceUrl) + pathSeparator + query.StringQueryParams()
}
