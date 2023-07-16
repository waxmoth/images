package utils

import (
	"bytes"
	"errors"
	"golang.org/x/image/draw"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
)

// CropImage crop the image from position x, y and with special width, height
func CropImage(img image.Image, x, y, width, height int) image.Image {
	type SubImage interface {
		image.Image
		SubImage(image.Rectangle) image.Image
	}
	rgbImg := img.(SubImage)
	subImg := rgbImg.SubImage(image.Rect(x, y, width, height)).(SubImage)
	return subImg
}

// DecodeImage decode image from reader and get the image format
func DecodeImage(body io.Reader) (image.Image, string, error) {
	img, format, err := image.Decode(body)
	return img, format, err
}

// ResizeImage resize image to special size and keep the ratio
func ResizeImage(img image.Image, width, height int) image.Image {
	ratio := float32(img.Bounds().Max.X) / float32(img.Bounds().Max.Y)
	if width > height {
		width = int(float32(height) * ratio)
	} else {
		height = int(float32(width) / ratio)
	}
	dst := image.NewRGBA(image.Rect(0, 0, width, height))
	if width > img.Bounds().Max.X || height > img.Bounds().Max.Y {
		draw.CatmullRom.Scale(dst, dst.Rect, img, img.Bounds(), draw.Over, nil)
	} else {
		draw.ApproxBiLinear.Scale(dst, dst.Rect, img, img.Bounds(), draw.Over, nil)
	}
	return dst
}

// ImageToBuffer convert image to buffer stream, return error if the format not be supported
func ImageToBuffer(img image.Image, format string) (bytes.Buffer, error) {
	buffer := new(bytes.Buffer)
	err := error(nil)
	switch format {
	case `jpeg`:
		err = jpeg.Encode(buffer, img, nil)
	case `png`:
		err = png.Encode(buffer, img)
	case `gif`:
		err = gif.Encode(buffer, img, nil)
	default:
		err = errors.New(`Unsupported format: ` + format)
	}

	return *buffer, err
}
