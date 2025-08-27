package null

import (
	"bytes"
	"encoding/json"
	"strconv"

	"github.com/dnuproject/pkg/jsonx"
)

var (
	_ json.Marshaler   = (*UintData)(nil)
	_ json.Unmarshaler = (*UintData)(nil)
)

type Uint = *UintData

type UintData struct {
	Value       uint64
	Initialized bool
}

func (u UintData) MarshalJSON() ([]byte, error) {
	if !u.Initialized {
		return bytes.Clone(jsonx.NullBytes), nil
	}

	return []byte(strconv.FormatUint(u.Value, 10)), nil
}

func (u *UintData) UnmarshalJSON(v []byte) error {
	if jsonx.IsNull(v) {
		return nil
	}

	value, err := strconv.ParseUint(string(v), 10, 64)
	if err != nil {
		return err
	}

	u.Value = value
	u.Initialized = true
	return nil
}
