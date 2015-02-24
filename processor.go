package main

import "github.com/k0kubun/pp"

type Processor struct {
}

func NewProcessor() *Processor {
	return &Processor{}
}

func (processor *Processor) Execute(image *Image, query *Query) {

	if len(query.Params) == 0 {
		return
	}

	pp.Println(query.Params, image.Width, image.Height)
}
