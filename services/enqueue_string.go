// Code generated by "stringer -type=Enqueue -trimprefix=Enqueue"; DO NOT EDIT.

package services

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Enqueueadd-0]
	_ = x[Enqueuenext-1]
	_ = x[Enqueueplay-2]
	_ = x[Enqueuereplace-3]
}

const _Enqueue_name = "addnextplayreplace"

var _Enqueue_index = [...]uint8{0, 3, 7, 11, 18}

func (i Enqueue) String() string {
	if i < 0 || i >= Enqueue(len(_Enqueue_index)-1) {
		return "Enqueue(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Enqueue_name[_Enqueue_index[i]:_Enqueue_index[i+1]]
}
