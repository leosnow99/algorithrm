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
