package manager

import "proudlygeek/goscii/encoder"

type MongoArtManager struct {
	Encoder encoder.EncoderInterface
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
