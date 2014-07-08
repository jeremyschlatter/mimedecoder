// Package mimedecoder decodes byte streams with various charsets and Content-Transfer-Encodings.
package mimedecoder

import (
	"encoding/base64"
	"errors"
	"io"

	charset_ "code.google.com/p/go.net/html/charset"
	"github.com/bom-d-van/qpencoding"
)

// NewDecoder returns an io.Reader that converts the contents of r to UTF-8.
func NewDecoder(r io.Reader, charset, contentTransferEncoding string) (io.Reader, error) {
	switch contentTransferEncoding {
	case "", "7bit", "8bit", "binary":
	case "base64":
		r = base64.NewDecoder(base64.StdEncoding, r)
	case "quoted-printable":
		r = qpencoding.NewReader(r)
	default:
		return nil, errors.New("decoder: unknown Content-Transfer-Encoding:" + contentTransferEncoding)
	}
	return charset_.NewReader(r, "text/plain; charset="+charset)
}
