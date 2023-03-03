// Code generated by jsonenums -type=Level; DO NOT EDIT.

package services

import (
	"encoding/json"
	"fmt"
)

var (
	_LevelNameToValue = map[string]Level{
		"Levelcritical": Levelcritical,
		"Leveldebug":    Leveldebug,
		"Levelerror":    Levelerror,
		"Levelfatal":    Levelfatal,
		"Levelinfo":     Levelinfo,
		"Levelwarning":  Levelwarning,
	}

	_LevelValueToName = map[Level]string{
		Levelcritical: "Levelcritical",
		Leveldebug:    "Leveldebug",
		Levelerror:    "Levelerror",
		Levelfatal:    "Levelfatal",
		Levelinfo:     "Levelinfo",
		Levelwarning:  "Levelwarning",
	}
)

func init() {
	var v Level
	if _, ok := interface{}(v).(fmt.Stringer); ok {
		_LevelNameToValue = map[string]Level{
			interface{}(Levelcritical).(fmt.Stringer).String(): Levelcritical,
			interface{}(Leveldebug).(fmt.Stringer).String():    Leveldebug,
			interface{}(Levelerror).(fmt.Stringer).String():    Levelerror,
			interface{}(Levelfatal).(fmt.Stringer).String():    Levelfatal,
			interface{}(Levelinfo).(fmt.Stringer).String():     Levelinfo,
			interface{}(Levelwarning).(fmt.Stringer).String():  Levelwarning,
		}
	}
}

// MarshalJSON is generated so Level satisfies json.Marshaler.
func (r Level) MarshalJSON() ([]byte, error) {
	if s, ok := interface{}(r).(fmt.Stringer); ok {
		return json.Marshal(s.String())
	}
	s, ok := _LevelValueToName[r]
	if !ok {
		return nil, fmt.Errorf("invalid Level: %d", r)
	}
	return json.Marshal(s)
}

// UnmarshalJSON is generated so Level satisfies json.Unmarshaler.
func (r *Level) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("Level should be a string, got %s", data)
	}
	v, ok := _LevelNameToValue[s]
	if !ok {
		return fmt.Errorf("invalid Level %q", s)
	}
	*r = v
	return nil
}
