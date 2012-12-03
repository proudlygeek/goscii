//
// An interface for implementing ascii art managers.
//
package manager

import (
	"io"
)

type ArtManager interface {
	Load(uri string) []byte
	Save(writer *io.Writer) string
}