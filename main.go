package main

import (
	"runtime"

	. "github.com/violetyk/graid/server"
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
