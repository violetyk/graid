package main

import (
	"image"
	"image/jpeg"
	"io"

	"github.com/disintegration/gift"
)

type Processor struct {
	gift *gift.GIFT
}

func NewProcessor() *Processor {
	return &Processor{gift: gift.New()}
}

func (processor *Processor) Execute(src *Image, dst io.Writer, query *Query) {
	processor.gift.Empty()

	if query.Count() > 0 {
		// resize
		if query.Has("w") || query.Has("h") {
			width := query.GetInt("w")
			height := query.GetInt("h")

			if width > 0 || height > 0 {
				processor.gift.Add(gift.Resize(width, height, gift.LanczosResampling))
			}
		}

		// crop
		if query.Has("c") {
			c := query.GetIntArray("c")
			if len(c) == 4 {
				processor.gift.Add(gift.Crop(image.Rect(c[0], c[1], c[2], c[3])))
			}
		}

		// Draw
		if len(processor.gift.Filters) > 0 {
			rgba := image.NewRGBA(processor.gift.Bounds(src.Object.Bounds()))
			processor.gift.Draw(rgba, src.Object)
			jpeg.Encode(dst, rgba, nil)
			return
		}
	}

	// default
	jpeg.Encode(dst, src.Object, nil)
}
