package model

import (
	"fmt"
	"github.com/goccy/go-json"
)

type EntityID string

func (e *EntityID) UnmarshalJSON(data []byte) error {

	var v interface{}
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	switch v.(type) {
	case string:
		*e = EntityID(v.(string))
	case []interface{}:
		switch v.([]interface{})[0].(type) {
		case string:
			*e = EntityID(v.([]interface{})[0].(string))
		default:
			return fmt.Errorf("invalid type for EntityID: %T", v)
		}
	default:
		return fmt.Errorf("invalid type for EntityID: %T", v)
	}

	return nil
}
