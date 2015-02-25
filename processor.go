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
	quality := 100

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

		// grayscale
		if query.Has("grayscale") {
			processor.gift.Add(gift.Grayscale())
		}

		// sepia
		if query.Has("sepia") {
			sepia := query.GetInt("sepia")
			if sepia <= 100 {
				processor.gift.Add(gift.Sepia(float32(sepia)))
			}
		}

		// contrast
		if query.Has("contrast") {
			contrast := query.GetInt("contrast")
			processor.gift.Add(gift.Contrast(float32(contrast)))
		}

		// brightness
		if query.Has("brightness") {
			brightness := query.GetInt("brightness")
			processor.gift.Add(gift.Brightness(float32(brightness)))
		}

		// saturation
		if query.Has("saturation") {
			saturation := query.GetInt("saturation")
			processor.gift.Add(gift.Saturation(float32(saturation)))
		}

		// colorize
		if query.Has("colorize") {
			colorize := query.GetIntArray("colorize")
			if len(colorize) == 3 {
				processor.gift.Add(gift.Colorize(float32(colorize[0]), float32(colorize[1]), float32(colorize[2])))
			}
		}

		// colorbalance
		if query.Has("colorbalance") {
			colorbalance := query.GetIntArray("colorbalance")
			if len(colorbalance) == 3 {
				processor.gift.Add(gift.ColorBalance(float32(colorbalance[0]), float32(colorbalance[1]), float32(colorbalance[2])))
			}
		}

		// quality
		if query.Has("q") {
			q := query.GetInt("q")
			if q > 0 && q < 100 {
				quality = q
			}
		}

		// Draw
		if len(processor.gift.Filters) > 0 {
			rgba := image.NewRGBA(processor.gift.Bounds(src.Object.Bounds()))
			processor.gift.Draw(rgba, src.Object)

			jpeg.Encode(dst, rgba, &jpeg.Options{Quality: quality})

			return
		}
	}

	// default
	jpeg.Encode(dst, src.Object, &jpeg.Options{Quality: quality})
}
