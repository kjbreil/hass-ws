// Code generated by "stringer -type=Format -trimprefix=Format"; DO NOT EDIT.

package services

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Formathls-0]
}

const _Format_name = "hls"

var _Format_index = [...]uint8{0, 3}

func (i Format) String() string {
	if i < 0 || i >= Format(len(_Format_index)-1) {
		return "Format(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Format_name[_Format_index[i]:_Format_index[i+1]]
}
