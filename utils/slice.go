package utils

type IntType interface {
	int | int8 | int16 | int32 | int64
}

type PtrIntType interface {
	*int | *int8 | *int16 | *int32 | *int64
}

type UIntType interface {
	uint | uint8 | uint16 | uint32 | uint64
}

type PtrUIntType interface {
	*uint | *uint8 | *uint16 | *uint32 | *uint64
}

type SimpleType interface {
	IntType | PtrIntType | UIntType | PtrUIntType | string | *string
}

func InSlice[T SimpleType](s []T, i T) bool {
	if len(s) == 0 {
		return false
	}

	for _, e := range s {
		if e == i {
			return true
		}
	}
	return false
}
