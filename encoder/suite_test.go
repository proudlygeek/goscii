package encoder_test

import (
	"github.com/proudlygeek/goscii/encoder"
	"image"
	_ "image/jpeg"
	_ "image/png"
	. "launchpad.net/gocheck"
	"log"
	"os"
	"testing"
)

// Hook up gocheck into the gotest runner.
func Test(t *testing.T) { TestingT(t) }

type EncoderSuite struct {
	Encoder *encoder.Encoder
}

// Mocks
type ReaderMock struct {
	Data []byte
}

type WriterMock struct {
	Data []byte
}

func (this *ReaderMock) Read(p []byte) (n int, err error) {
	for _, b := range this.Data {
		p = append(p, b)
	}

	return len(p), err
}

func (this *WriterMock) Write(p []byte) (n int, err error) {
	for _, b := range p {
		this.Data = append(this.Data, b)
	}

	return len(this.Data), err
}

// Init gocheck

var _ = Suite(&EncoderSuite{})

// Setup
func (s *EncoderSuite) SetUpTest(c *C) {
	s.Encoder = &encoder.Encoder{}
}

// Load test image utils
func (s *EncoderSuite) LoadImage(filename string) (m image.Image, err error) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	m, _, err = image.Decode(file)
	if err != nil {
		log.Fatal(err)
	}
	return
}
