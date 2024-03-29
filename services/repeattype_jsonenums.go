// Code generated by jsonenums -type=RepeatType; DO NOT EDIT.

package services

import (
	"fmt"
	"github.com/goccy/go-json"
)

var (
	_RepeatTypeNameToValue = map[string]RepeatType{
		"RepeatTypepause":  RepeatTypepause,
		"RepeatTyperepeat": RepeatTyperepeat,
		"RepeatTypesingle": RepeatTypesingle,
	}

	_RepeatTypeValueToName = map[RepeatType]string{
		RepeatTypepause:  "RepeatTypepause",
		RepeatTyperepeat: "RepeatTyperepeat",
		RepeatTypesingle: "RepeatTypesingle",
	}
)

func init() {
	var v RepeatType
	if _, ok := interface{}(v).(fmt.Stringer); ok {
		_RepeatTypeNameToValue = map[string]RepeatType{
			interface{}(RepeatTypepause).(fmt.Stringer).String():  RepeatTypepause,
			interface{}(RepeatTyperepeat).(fmt.Stringer).String(): RepeatTyperepeat,
			interface{}(RepeatTypesingle).(fmt.Stringer).String(): RepeatTypesingle,
		}
	}
}

// MarshalJSON is generated so RepeatType satisfies json.Marshaler.
func (r RepeatType) MarshalJSON() ([]byte, error) {
	if s, ok := interface{}(r).(fmt.Stringer); ok {
		return json.Marshal(s.String())
	}
	s, ok := _RepeatTypeValueToName[r]
	if !ok {
		return nil, fmt.Errorf("invalid RepeatType: %d", r)
	}
	return json.Marshal(s)
}

// UnmarshalJSON is generated so RepeatType satisfies json.Unmarshaler.
func (r *RepeatType) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("RepeatType should be a string, got %s", data)
	}
	v, ok := _RepeatTypeNameToValue[s]
	if !ok {
		return fmt.Errorf("invalid RepeatType %q", s)
	}
	*r = v
	return nil
}
