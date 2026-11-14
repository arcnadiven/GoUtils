package utils

import (
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"os"
	"strings"

	"github.com/pkg/errors"
)

func DecodeImage(filePath string) (image.Image, error) {
	file, err := os.OpenFile(filePath, os.O_RDONLY, 0444)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	defer file.Close()

	img := image.Image(nil)
	decodeErrs := []string{}

	// 尝试所有的解码
	for _, fn := range []func(r io.Reader) (image.Image, error){jpeg.Decode, png.Decode, gif.Decode} {
		err := error(nil)
		img, err = fn(file)
		if err == nil {
			break
		}
		decodeErrs = append(decodeErrs, err.Error())
		if _, err := file.Seek(0, io.SeekStart); err != nil {
			return nil, errors.WithStack(err)
		}
	}
	if img == nil {
		return nil, errors.Errorf("image %s decode error: %s", filePath, strings.Join(decodeErrs, "; "))
	}
	return img, nil
}
