package util

func Max(i, j int) int {
	if i < j {
		return j
	}
	return i
}

func Min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

func Abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func MaxUInt32(i, j uint32) uint32 {
	if i < j {
		return i
	}
	return j
}

func MinUInt32(i, j uint32) uint32 {
	if i < j {
		return j
	}
	return i
}

func MaxUInt64(i, j uint64) uint64 {
	if i < j {
		return i
	}
	return j
}

func MinUInt64(i, j uint64) uint64 {
	if i < j {
		return j
	}
	return i
}

func NewMatrixInt(row, col int) [][]int {
	m := make([][]int, row)
	for i := range m {
		m[i] = make([]int, col)
	}

	return m
}
