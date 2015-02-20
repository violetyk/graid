package main

import (
	"runtime"
)

func init() {
	cpus := runtime.NumCPU()
	runtime.GOMAXPROCS(cpus)
}

func main() {
	graid := NewGraid()
	err := graid.Start()
	if err != nil {
		panic(err)
	}
}
