package main

import "os"

type FileEngineAdapter struct {
}

func NewFileEngineAdapter() *FileEngineAdapter {
	return &FileEngineAdapter{}
}

func (adapter *FileEngineAdapter) CacheKey(query *Query) string {
	config := LoadConfig()
	path := config.Cache.File.Path
	// TODO: directory traversal
	return path + string(os.PathSeparator) + query.SourceUrl + string(os.PathSeparator) + query.StringQueryParams()
}
