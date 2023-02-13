package encoding

type EncoderDecoder interface {
	MaxEncodedLen(n int) int
	EncodeToString(src []byte) (string, error)
	DecodeString(src string) ([]byte, error)
}
