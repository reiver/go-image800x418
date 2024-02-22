package image800x418

import (
	"image"
)

func Bounds() image.Rectangle {
	return image.Rectangle{
		Min: image.Point{0,0},
		Max: image.Point{Width,Height},
	}
}
