package manager

import "github.com/proudlygeek/goscii/encoder"

type MongoArtManager struct {
	Encoder     encoder.EncoderInterface
	DatabaseURL string
}

type Art struct {
	Content []byte
}

type Doc struct {
	C int
}

type MongoWriter struct {
	Buffer []byte
}
