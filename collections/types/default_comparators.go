package types

func StringComparator(a, b string) int {
	if a == b {
		return 0
	}
	if a < b {
		return -1
	}
	return 1
}

func IntComparator(a, b int) int {
	return a - b
}

func Float32Comparator(a, b float32) int {
	if a == b {
		return 0
	}
	if a < b {
		return -1
	}
	return 1
}

func Float64Comparator(a, b float64) int {
	if a == b {
		return 0
	}
	if a < b {
		return -1
	}
	return 1
}

func BoolComparator(a, b bool) int {
	if a == b {
		return 0
	}
	if a {
		return 1
	}
	return -1
}

func RuneComparator(a, b rune) int {
	return int(a) - int(b)
}

func ByteComparator(a, b byte) int {
	return int(a) - int(b)
}

func Complex64Comparator(a, b complex64) int {
	if real(a) == real(b) {
		return 0
	}
	if real(a) < real(b) {
		return -1
	}
	return 1
}

func Complex128Comparator(a, b complex128) int {
	if real(a) == real(b) {
		return 0
	}
	if real(a) < real(b) {
		return -1
	}
	return 1
}
