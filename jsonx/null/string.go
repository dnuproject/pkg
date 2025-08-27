package null

import (
	"bytes"
	"encoding/json"

	"github.com/dnuproject/pkg/jsonx"
)

var (
	_ json.Marshaler   = (*StringData)(nil)
	_ json.Unmarshaler = (*StringData)(nil)
)

type String = *StringData

type StringData struct {
	Value       string
	Initialized bool
}

func (s StringData) MarshalJSON() ([]byte, error) {
	if !s.Initialized {
		return bytes.Clone(jsonx.NullBytes), nil
	}

	return []byte(s.Value), nil
}

func (s *StringData) UnmarshalJSON(v []byte) error {
	if jsonx.IsNull(v) {
		return nil
	}

	s.Value = string(v)
	s.Initialized = true
	return nil
}
