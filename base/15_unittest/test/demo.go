package test

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func IntMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}
