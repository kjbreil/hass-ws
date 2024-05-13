// Code generated by jsonenums -type=Qos; DO NOT EDIT.

package services

import (
	"encoding/json"
	"fmt"
)

var (
	_QosNameToValue = map[string]Qos{
		"Qos0": Qos0,
		"Qos1": Qos1,
		"Qos2": Qos2,
	}

	_QosValueToName = map[Qos]string{
		Qos0: "Qos0",
		Qos1: "Qos1",
		Qos2: "Qos2",
	}
)

func init() {
	var v Qos
	if _, ok := interface{}(v).(fmt.Stringer); ok {
		_QosNameToValue = map[string]Qos{
			interface{}(Qos0).(fmt.Stringer).String(): Qos0,
			interface{}(Qos1).(fmt.Stringer).String(): Qos1,
			interface{}(Qos2).(fmt.Stringer).String(): Qos2,
		}
	}
}

// MarshalJSON is generated so Qos satisfies json.Marshaler.
func (r Qos) MarshalJSON() ([]byte, error) {
	if s, ok := interface{}(r).(fmt.Stringer); ok {
		return json.Marshal(s.String())
	}
	s, ok := _QosValueToName[r]
	if !ok {
		return nil, fmt.Errorf("invalid Qos: %d", r)
	}
	return json.Marshal(s)
}

// UnmarshalJSON is generated so Qos satisfies json.Unmarshaler.
func (r *Qos) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("Qos should be a string, got %s", data)
	}
	v, ok := _QosNameToValue[s]
	if !ok {
		return fmt.Errorf("invalid Qos %q", s)
	}
	*r = v
	return nil
}
