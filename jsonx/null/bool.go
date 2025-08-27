package null

import (
	"bytes"
	"encoding/json"
	"strconv"

	"github.com/dnuproject/pkg/jsonx"
)

var (
	_ json.Unmarshaler = (*BoolData)(nil)
	_ json.Marshaler   = (*BoolData)(nil)
)

type Bool = *BoolData

type BoolData struct {
	Value       bool
	Initialized bool
}

var (
	NullBool = &BoolData{}

	True = &BoolData{
		Value:       true,
		Initialized: true,
	}

	False = &BoolData{
		Value:       false,
		Initialized: true,
	}
)

func (b BoolData) MarshalJSON() ([]byte, error) {
	if !b.Initialized {
		return bytes.Clone(jsonx.NullBytes), nil
	}

	return []byte(strconv.FormatBool(b.Value)), nil
}

func (b *BoolData) UnmarshalJSON(v []byte) error {
	if jsonx.IsNull(v) {
		return nil
	}

	value, err := strconv.ParseBool(string(v))
	if err != nil {
		return err
	}

	b.Value = value
	b.Initialized = true
	return nil
}
