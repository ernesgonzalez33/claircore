// Code generated by "stringer -type=Severity"; DO NOT EDIT.

package claircore

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Unknown-0]
	_ = x[Negligible-1]
	_ = x[Low-2]
	_ = x[Medium-3]
	_ = x[High-4]
	_ = x[Critical-5]
	_ = x[Defcon1-6]
}

const _Severity_name = "UnknownNegligibleLowMediumHighCriticalDefcon1"

var _Severity_index = [...]uint8{0, 7, 17, 20, 26, 30, 38, 45}

func (i Severity) String() string {
	if i >= Severity(len(_Severity_index)-1) {
		return "Severity(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Severity_name[_Severity_index[i]:_Severity_index[i+1]]
}