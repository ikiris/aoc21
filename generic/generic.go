package generic

type SignedInteger interface {
	int | int8 | int16 | int32 | int64
}

type UnSignedInteger interface {
	uint | uint8 | uint16 | uint32 | uint64
}

type Number interface {
	SignedInteger | UnSignedInteger
}

func AddMapVals[T Number, T2 comparable](m map[T2]T) T {
	var result T
	for _, n := range m {
		result += n
	}
	return result
}
