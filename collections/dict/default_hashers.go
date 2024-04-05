package dict

import "hash/fnv"

// File: default_hashers.go
// Common hash function for go types

func stringHasher(s string) int {
	h := fnv.New32a()
	_, err := h.Write([]byte(s))
	if err != nil {
		return 0
	}
	return int(h.Sum32()) // TODO make the hasher interface be a uint32
}

func intHasher(i int) int {
	return i
}
