// Code generated by "stringer -type=Repeat -trimprefix=Repeat"; DO NOT EDIT.

package services

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Repeatall-0]
	_ = x[Repeatoff-1]
	_ = x[Repeatone-2]
}

const _Repeat_name = "alloffone"

var _Repeat_index = [...]uint8{0, 3, 6, 9}

func (i Repeat) String() string {
	if i < 0 || i >= Repeat(len(_Repeat_index)-1) {
		return "Repeat(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Repeat_name[_Repeat_index[i]:_Repeat_index[i+1]]
}