// Code generated by "stringer -type Type"; DO NOT EDIT.

package licensecheck

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[AGPL-0]
	_ = x[Apache-1]
	_ = x[BSD-2]
	_ = x[CC-3]
	_ = x[GPL-4]
	_ = x[JSON-5]
	_ = x[MIT-6]
	_ = x[Unlicense-7]
	_ = x[Zlib-8]
	_ = x[Other-9]
}

const _Type_name = "AGPLApacheBSDCCGPLJSONMITUnlicenseZlibOther"

var _Type_index = [...]uint8{0, 4, 10, 13, 15, 18, 22, 25, 34, 38, 43}

func (i Type) String() string {
	if i < 0 || i >= Type(len(_Type_index)-1) {
		return "Type(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Type_name[_Type_index[i]:_Type_index[i+1]]
}
