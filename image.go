package main

import (
	"bytes"
	"image"
	"image/jpeg"
)

type Image struct {
	Format string
	Object *image.Image
	Width  int
	Height int
}

func NewImage(data []byte, format string) (self *Image, err error) {

	var i image.Image
	var w, h int

	r := bytes.NewReader(data)

	switch format {
	case "jpg", "jpeg":
		i, err = jpeg.Decode(r)
		if err != nil {
			panic(err)
		}
		if err == nil {
			r.Seek(0, 0)
			c, err := jpeg.DecodeConfig(r)
			if err == nil {
				w = c.Width
				h = c.Height
			}
		}
		format = "jpeg"
	}

	return &Image{
		Format: format,
		Object: &i,
		Width:  w,
		Height: h,
	}, err
}
