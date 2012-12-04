package encoder

import (
	"image"
	"io"
)

type Encoder struct{}

type EncoderInterface interface {
	EncodeURI(n int) (slug string)
	DecodeURI(uri string) (n int)
	DecodeImage(stream io.Reader) (m image.Image, err error)
	Asciify(m image.Image, res io.Writer) (err error)
}
