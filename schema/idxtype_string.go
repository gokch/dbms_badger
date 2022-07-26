// Code generated by "stringer -type=IdxType"; DO NOT EDIT.

package schema

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[IdxTypePrimary-0]
	_ = x[IdxTypeUnique-1]
	_ = x[IdxTypeRange-2]
	_ = x[IdxTypeBitset-3]
}

const _IdxType_name = "IdxTypePrimaryIdxTypeUniqueIdxTypeRangeIdxTypeBitset"

var _IdxType_index = [...]uint8{0, 14, 27, 39, 52}

func (i IdxType) String() string {
	if i >= IdxType(len(_IdxType_index)-1) {
		return "IdxType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _IdxType_name[_IdxType_index[i]:_IdxType_index[i+1]]
}
