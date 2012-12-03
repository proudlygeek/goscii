//
// Unit tests for package proudlygeek/goscii/encoder 
//
package encoder_test

import (
	. "launchpad.net/gocheck"
	"fmt"
	"strings"
)

// Simple hello world test
func (s *EncoderSuite) TestHello(c *C) {
	c.Check(42, Equals, 42)
}

// Test Passes 
// if TestEncodeURI returns a base-62
// string.
func (s *EncoderSuite) TestEncodeUri(c *C) {
	res := s.Encoder.EncodeURI(123456789)
	c.Assert(res, Equals, "HUawi")
}

// Test Passes 
// if TestDecodeURI returns a base-10
// integer.
func (s *EncoderSuite) TestDecodeUri(c *C) {
	res := s.Encoder.DecodeURI("HUawi")
	c.Assert(res, Equals, 123456789)
}

// Test Passes
// if Asciify returns an HTML string containing "~".
func (s *EncoderSuite) TestBlackPixelAsciify(c *C) {
	input, _ := s.LoadImage("tests/black_1x1.jpg")
	output := &WriterMock{}
	err := s.Encoder.Asciify(input, output)
	c.Assert(err, Equals, nil)
	sym := fmt.Sprintf("%s", output.Data)

	c.Assert(strings.Contains(sym, "~"), Equals, true)
}

// Test Passes
// if Asciify returns an HTML string containing "@".
func (s *EncoderSuite) TestWhitePixelAsciify(c *C) {
	input, _ := s.LoadImage("tests/white_1x1.jpg")
	output := &WriterMock{}
	err := s.Encoder.Asciify(input, output)
	c.Assert(err, Equals, nil)
	sym := fmt.Sprintf("%s", output.Data)

	c.Assert(strings.Contains(sym, "@"), Equals, true)
}
