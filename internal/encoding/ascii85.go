package encoding

import (
	"encoding/ascii85"
	"errors"

	"github.com/vmihailenco/taskq/v3/internal"
)

var ASCII85EncDec ASCII85EncoderDecoder

// ASCII85EncoderDecoder is a struct that implements the Encoder interface for ASCII85 encoding
type ASCII85EncoderDecoder struct{}

// MaxEncodedLen returns the maximum length of an ASCII85-encoded byte slice with the given length
func (e ASCII85EncoderDecoder) MaxEncodedLen(n int) int {
	return ascii85.MaxEncodedLen(n)
}

// EncodeToString returns the ASCII85-encoded form of the input byte slice
func (e ASCII85EncoderDecoder) EncodeToString(src []byte) (string, error) {
	dst := make([]byte, e.MaxEncodedLen(len(src)))
	n := ascii85.Encode(dst, src)
	dst = dst[:n]
	return internal.BytesToString(dst), nil
}

// DecodeString returns the ASCII85-decoded form of the input byte slice
func (e ASCII85EncoderDecoder) DecodeString(src string) ([]byte, error) {
	dst := make([]byte, len(src))
	ndst, nsrc, err := ascii85.Decode(dst, internal.StringToBytes(src), true)
	if err != nil {
		return nil, err
	}
	if nsrc != len(src) {
		return nil, errors.New("ascii85: src is not fully decoded")
	}
	return dst[:ndst], nil
}
