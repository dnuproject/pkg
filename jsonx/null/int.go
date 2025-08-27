package null

import (
	"bytes"
	"encoding/json"
	"strconv"

	"github.com/dnuproject/pkg/jsonx"
)

var (
	_ json.Unmarshaler = (*IntData)(nil)
	_ json.Marshaler   = (*IntData)(nil)
)

type Int = *IntData

type IntData struct {
	Value       int64
	Initialized bool
}

func (i IntData) MarshalJSON() ([]byte, error) {
	if !i.Initialized {
		return bytes.Clone(jsonx.NullBytes), nil
	}

	return []byte(strconv.FormatInt(i.Value, 10)), nil
}

func (i *IntData) UnmarshalJSON(v []byte) error {
	if jsonx.IsNull(v) {
		return nil
	}

	value, err := strconv.ParseInt(string(v), 10, 64)
	if err != nil {
		return err
	}

	i.Value = value
	i.Initialized = true
	return nil
}
