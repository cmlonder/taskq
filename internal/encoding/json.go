package encoding

import (
	"encoding/json"
)

var JSONEncDec JSONEncoderDecoder

// JSONEncoderDecoder is a struct that implements the Encoder interface for JSON encoding
type JSONEncoderDecoder struct{}

// MaxEncodedLen returns the maximum length of a JSON-encoded byte slice with the given length
func (e JSONEncoderDecoder) MaxEncodedLen(n int) int {
	return n
}

// EncodeToString returns the JSON-encoded form of the input byte slice
func (e JSONEncoderDecoder) EncodeToString(src []byte) (string, error) {
	var decoded interface{}
	if err := json.Unmarshal(src, &decoded); err != nil {
		return "", err
	}
	encoded, err := json.Marshal(decoded)
	if err != nil {
		return "", err
	}
	return string(encoded), nil
}

// DecodeString returns the JSON-decoded form of the input string
func (e JSONEncoderDecoder) DecodeString(src string) ([]byte, error) {
	var decoded interface{}
	if err := json.Unmarshal([]byte(src), &decoded); err != nil {
		return nil, err
	}
	encoded, err := json.Marshal(decoded)
	if err != nil {
		return nil, err
	}
	return encoded, nil
}
