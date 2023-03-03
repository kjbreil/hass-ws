// Code generated by jsonenums -type=Mode; DO NOT EDIT.

package services

import (
	"encoding/json"
	"fmt"
)

var (
	_ModeNameToValue = map[string]Mode{
		"Modedark":  Modedark,
		"Modelight": Modelight,
	}

	_ModeValueToName = map[Mode]string{
		Modedark:  "Modedark",
		Modelight: "Modelight",
	}
)

func init() {
	var v Mode
	if _, ok := interface{}(v).(fmt.Stringer); ok {
		_ModeNameToValue = map[string]Mode{
			interface{}(Modedark).(fmt.Stringer).String():  Modedark,
			interface{}(Modelight).(fmt.Stringer).String(): Modelight,
		}
	}
}

// MarshalJSON is generated so Mode satisfies json.Marshaler.
func (r Mode) MarshalJSON() ([]byte, error) {
	if s, ok := interface{}(r).(fmt.Stringer); ok {
		return json.Marshal(s.String())
	}
	s, ok := _ModeValueToName[r]
	if !ok {
		return nil, fmt.Errorf("invalid Mode: %d", r)
	}
	return json.Marshal(s)
}

// UnmarshalJSON is generated so Mode satisfies json.Unmarshaler.
func (r *Mode) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("Mode should be a string, got %s", data)
	}
	v, ok := _ModeNameToValue[s]
	if !ok {
		return fmt.Errorf("invalid Mode %q", s)
	}
	*r = v
	return nil
}