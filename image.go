package main

import (
	"bytes"
	"errors"
	"image"
	"image/jpeg"
	"reflect"
)

const (
	JPEG = 1
)

type Image struct {
	Format int
	Object image.Image
	Width  int
	Height int
}

func NewImage(data []byte) (self *Image, err error) {

	var i image.Image
	var w, h int
	var f int

	// detect format
	if reflect.DeepEqual(data[0:2], []byte{0xff, 0xd8}) {
		f = JPEG
	} else {
		err = errors.New("unsupported format")
	}

	r := bytes.NewReader(data)

	switch f {
	case JPEG:
		i, err = jpeg.Decode(r)
		if err == nil {
			r.Seek(0, 0)
			c, err := jpeg.DecodeConfig(r)
			if err == nil {
				w = c.Width
				h = c.Height
			}
		}
	}

	return &Image{
		Format: f,
		Object: i,
		Width:  w,
		Height: h,
	}, err
}
