package set

import "utils-generics/collections/dict"

type FlatSet[K any] struct {
	innerMap *dict.FlatMap[K, bool]
}

// MakeFlatSet creates a new FlatSet.
func MakeFlatSet[K any](c func(a, b K) int) *FlatSet[K] {
	return &FlatSet[K]{innerMap: dict.MakeFlatMap[K, bool](c)}
}

// Add adds a new element to the set.
// This operation is idempotent, so if the element already exists in the set, it is equivalent to a no-op.
func (s *FlatSet[K]) Add(val K) {
	s.innerMap.Put(val, true)
}

// Remove removes an element from the set.
// If the element exists it is removed and the function returns true, otherwise it returns false.
func (s *FlatSet[K]) Remove(val K) bool {
	return s.innerMap.Remove(val)
}

// Contains returns true if the element exists in the set, otherwise it returns false.
func (s *FlatSet[K]) Contains(val K) bool {
	_, ok := s.innerMap.Get(val)
	return ok
}

// Size returns the number of elements in the set.
func (s *FlatSet[K]) Size() int {
	return s.innerMap.Size()
}

func (s *FlatSet[K]) IsEmpty() bool {
	return s.innerMap.IsEmpty()
}

func (s *FlatSet[K]) IsNotEmpty() bool {
	return s.innerMap.IsNotEmpty()
}

func (s *FlatSet[K]) Clear() {
	s.innerMap.Clear()
}

func (s *FlatSet[K]) Formatted() string {
	//TODO implement me
	panic("implement me")
}
