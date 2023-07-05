package set

import (
	"fmt"
	"io.analogalex.collections/collections/dict"
)

// BinaryTreeSet is a set implementation using a binary tree map -- roughly speaking it implements a binary tree solution.
//
// It stores its elements in natural order i.e. ints are stores in ascending order, strings are stored in alphabetical order.
// It is not thread safe and should not be used for concurrent access
// (see concurrent package for thread safe implementations).
//
// It's performance characteristics are:
//
// - Add: O(log n)
//
// - Remove: O(log n)
//
// - Contains: O(log n)
type BinaryTreeSet struct {
	innerMap *dict.BinaryTreeMap
}

// MakeBinaryTreeSet creates a new BinaryTreeSet
func MakeBinaryTreeSet() *BinaryTreeSet {
	return &BinaryTreeSet{innerMap: dict.MakeBinaryTreeMap()}
}

// Add adds a new element to the set
// This operation is idempotent, so if the element already exists in the set, it is equivalent to a no-op
func (s *BinaryTreeSet) Add(val int) {
	s.innerMap.Put(val, "")
}

// Remove removes an element from the set
func (s *BinaryTreeSet) Remove(val int) bool {
	return s.innerMap.Remove(val)
}

// Contains checks if the set contains an element
func (s *BinaryTreeSet) Contains(val int) bool {
	return s.innerMap.ContainsKey(val)
}

// Size returns the size of the set
func (s *BinaryTreeSet) Size() int {
	return s.innerMap.Size()
}

// Clear removes all elements from the set
func (s *BinaryTreeSet) Clear() {
	s.innerMap.Clear()
}

// IsEmpty checks if the set is empty
func (s *BinaryTreeSet) IsEmpty() bool {
	return s.innerMap.IsEmpty()
}

// IsNotEmpty checks if the set is not empty
func (s *BinaryTreeSet) IsNotEmpty() bool {
	return s.innerMap.IsNotEmpty()
}

// Formatted returns a string representation of the set
func (s *BinaryTreeSet) Formatted() string {
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

// ----------------
// OrderedSet methods

// First returns the first element of the set
func (s *BinaryTreeSet) First() int {
	key, _ := s.innerMap.First()
	return key
}

// Last returns the last element of the set
func (s *BinaryTreeSet) Last() int {
	key, _ := s.innerMap.Last()
	return key
}

// RemoveFirst removes the first element of the set
func (s *BinaryTreeSet) RemoveFirst() bool {
	return s.innerMap.RemoveFirst()
}

// RemoveLast removes the last element of the set
func (s *BinaryTreeSet) RemoveLast() bool {
	return s.innerMap.RemoveLast()
}

// ToSortedSlice returns an array of the elements of the set in order
func (s *BinaryTreeSet) ToSortedSlice() []int {
	return s.innerMap.Keys()
}
