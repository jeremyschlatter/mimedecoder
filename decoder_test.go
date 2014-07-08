package mimedecoder

import (
	"fmt"
	"io/ioutil"
	"strings"
	"testing"
)

func Example() {
	decoder, _ := NewDecoder(strings.NewReader("=22I think, therefore I am.=22 =\n-Ren=E9 Descartes"), "windows-1252", "quoted-printable")
	b, _ := ioutil.ReadAll(decoder)
	fmt.Println(string(b))
	// Output: "I think, therefore I am." -René Descartes
}

func TestDecoder(t *testing.T) {
	decoderTests := []struct {
		s, charset, contentTransferEncoding, want string
	}{
		{"I=92m happy", "latin1", "quoted-printable", "I’m happy"},
		{"dGhpcyBpcyBiYXNlNjQtZW5jb2RlZA==", "utf-8", "base64", "this is base64-encoded"},
		{"Plain utf-8. ✓", "utf-8", "", "Plain utf-8. ✓"},
	}
	for _, tt := range decoderTests {
		decoder, err := NewDecoder(strings.NewReader(tt.s), tt.charset, tt.contentTransferEncoding)
		if err != nil {
			t.Errorf("NewDecoder(strings.NewReader(%s), %s, %s) returned an unexpected error: %v\n", tt.s, tt.charset, tt.contentTransferEncoding, err)
			continue
		}
		b, err := ioutil.ReadAll(decoder)
		if err != nil {
			t.Errorf("unexpected error returned from reader produced by NewDecoder(strings.NewReader(%s), %s, %s): %v\n", tt.s, tt.charset, tt.contentTransferEncoding, err)
			continue
		}
		if string(b) != tt.want {
			t.Errorf("Decoding \"%s\" with charset \"%s\" and Content-Transfer-Encoding \"%s\".\nGot:\n%s\nWant:\n%s\n", tt.s, tt.charset, tt.contentTransferEncoding, string(b), tt.want)
		}
	}
}
