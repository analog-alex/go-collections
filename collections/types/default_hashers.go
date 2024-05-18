package types

import "hash/fnv"

// File: default_hashers.go
// Common hash function for go types

func StringHash(s string) int {
	h := fnv.New32a()
	_, err := h.Write([]byte(s))
	if err != nil {
		return 0
	}
	return int(h.Sum32()) // TODO make the hasher interface be a uint32
}

func IntHash(i int) int {
	return i
}

func Float32Hash(f float32) int {
	return int(f)
}

func Float64Hash(f float64) int {
	return int(f)
}

func BoolHash(b bool) int {
	if b {
		return 1
	}
	return 0
}

func RuneHash(r rune) int {
	return int(r)
}

func ByteHash(b byte) int {
	return int(b)
}

func Complex64Hash(c complex64) int {
	return int(real(c)) + int(imag(c))
}

func Complex128Hash(c complex128) int {
	return int(real(c)) + int(imag(c))
}
