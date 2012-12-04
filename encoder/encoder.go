package encoder

import (
    "fmt"
    "io"
    "strings"
    "math"
    "image"
    _ "image/jpeg"
    _ "image/png"
)

const alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
const base = len(alphabet)

//
// Decodes an image.
//
func (this *Encoder) DecodeImage(stream io.Reader) (m image.Image, err error) {
    m, _, err = image.Decode(stream)
    return
}

/** 
 * Converts an image to a stream of ASCII Chars.
 * 
 * The image stream is scanned in blocks of 4x3 pixels and the
 * algorithm performs as follows:
 * 
 * 1. Read pixel's R, G, B and Alpha channels and generate a <span>
 *    tag;
 * 2. Calculate the pixel's luminance and approximates it with an 
 *    ASCII symbol.
 *
 * The result is direcly written in res io.Writer.
 *
 */
func (this *Encoder) Asciify(m image.Image, res io.Writer) (err error) {

    bounds := m.Bounds()

    horizontalLimit := bounds.Max.X
    verticalLimit := bounds.Max.Y

    fmt.Fprintln(res, "<div class='content'><div class='result'>")
    fmt.Fprintln(res, "<br>")

    // Init ascii result buffer
    var ascii = make([][]string, verticalLimit / 5 + 1)

    for i := 0; i < len(ascii); i++ {
        ascii[i] = make([]string, horizontalLimit / 4 + 1)
    }

    i := 0
    for y := bounds.Min.Y; y < bounds.Max.Y; y += 5 {
        j := 0
        for x := bounds.Min.X; x < bounds.Max.X; x += 4 {
            r, g, b, a := m.At(x, y).RGBA()
            r ,g, b, a = r >> 8, g >> 8, b >> 8, a >> 8

            // Luminance formula (see http://en.wikipedia.org/wiki/YUV#Formulas)
            luminance :=  0.299 * float64(r) + 0.587 * float64(g) + 0.114 * float64(b)
            sym := "~"
            scale := int(luminance) >> 4

            switch scale {
                case 1, 2, 3, 4: sym = "^"
                case 5, 6, 7, 8: sym = "G"
                case 10, 11, 12, 13, 14, 15, 16: sym = "@"
            }

            ascii[i][j] = fmt.Sprintf("<span style='color: rgba(%d, %d, %d, %d);'>%s</span>", r, g, b, a, sym)
            j++
        }
        i++
    }

    for i := 0; i < len(ascii); i+=1 {
        for j := 0; j < len(ascii[i]); j+=1 {
            fmt.Fprint(res, ascii[i][j])
        }
        fmt.Fprint(res, "<br>")
    }
    fmt.Fprintln(res, "</div></div>")

    return
}

/**
 * Encodes a base 10 integer to a base 62 base.
 *
 * The alphabet is composed of the following 62 characters:
 * 
 * /[a-zA-Z0-9]/
 *
 */
func (this *Encoder) EncodeURI(n int) (slug string) {

    for n > 0 {
        remainder := n % base
        n = n / base
        slug += fmt.Sprintf("%c", alphabet[remainder])
    }

    return
}

/**
 * Decodes a base 62 string to a base 10 integer.
 */
func  (this *Encoder) DecodeURI(uri string) (n int) {

    for i, c := range uri {
        n += int(math.Pow(float64(base), float64(i))) * strings.IndexRune(alphabet, c)
    }

    return
}