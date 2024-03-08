package set

import (
	"fmt"
	"utils-generics/collections/dict"
)

// HashSet is a set implementation using a hash map -- roughly speaking it implements a hash table solution.
//
// It makes no guarantees on ordering of its elements.
// It is not thread safe and should not be used for concurrent access
// (see concurrent package for thread safe implementations).
//
// It's performance characteristics are:
//
// - Add: O(1)
//
// - Remove: O(1)
//
// - Contains: O(1)
type HashSet struct {
	innerMap *dict.HashMap[bool]
}

// MakeHashSet creates a new HashSet.
func MakeHashSet() *HashSet {
	return &HashSet{innerMap: dict.MakeHashMap[bool]()}
}

// Add adds a new element to the set.
// This operation is idempotent, so if the element already exists in the set, it is equivalent to a no-op.
func (s *HashSet) Add(val int) {
	s.innerMap.Put(val, true)
}

// Remove removes an element from the set.
// If the element exists it is removed and the function returns true, otherwise it returns false.
func (s *HashSet) Remove(val int) bool {
	return s.innerMap.Remove(val)
}

// Contains returns true if the element exists in the set, otherwise it returns false.
func (s *HashSet) Contains(val int) bool {
	return s.innerMap.ContainsKey(val)
}

// Size returns the number of elements in the set.
func (s *HashSet) Size() int {
	return s.innerMap.Size()
}

// Clear removes all elements from the set.
func (s *HashSet) Clear() {
	s.innerMap.Clear()
}

// IsEmpty returns true if the set is empty, otherwise it returns false.
func (s *HashSet) IsEmpty() bool {
	return s.Size() == 0
}

// IsNotEmpty returns true if the set is not empty, otherwise it returns false.
func (s *HashSet) IsNotEmpty() bool {
	return !s.IsEmpty()
}

// Formatted returns a string representation of the set.
func (s *HashSet) Formatted() string {
	entries := s.innerMap.Entries()
	str := "{"

	for _, entry := range entries {
		str += fmt.Sprintf("%d, ", entry.Key)
	}

	// remove last comma and space str if it exists
	if len(str) > 1 {
		str = str[:len(str)-2]
	}

	str += "}"
	return str
}
