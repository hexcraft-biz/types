package types

import (
	"database/sql/driver"
	"encoding/hex"
	"errors"
)

type RfidUid string

func (r *RfidUid) Scan(value any) error {
	switch v := value.(type) {
	case []byte:
		*r = RfidUid(hex.EncodeToString(v))
		return nil
	case string:
		*r = RfidUid(v)
		return nil
	default:
		return errors.New("invalid type for RfidUid")
	}
}

func (r RfidUid) Value() (driver.Value, error) {
	bytes, err := hex.DecodeString(string(r))
	if err != nil {
		return nil, err
	}
	return bytes, nil
}
