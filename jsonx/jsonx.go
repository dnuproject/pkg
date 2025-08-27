package jsonx

import "bytes"

var (
	NullBytes = []byte(`null`)
	ZeroBytes = []byte(`0`)
)

func IsNull(value []byte) bool {
	return bytes.Equal(value, NullBytes)
}

func IsZero(value []byte) bool {
	return bytes.Equal(value, ZeroBytes)
}
