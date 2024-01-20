package set

import "io.analogalex.collections/collections/dict"

type FlatSet struct {
	innerMap *dict.FlatMap
}

// MakeFlatSet creates a new FlatSet.
func MakeFlatSet() *FlatSet {
	return &FlatSet{innerMap: dict.MakeFlatMap()}
}

// Add adds a new element to the set.
// This operation is idempotent, so if the element already exists in the set, it is equivalent to a no-op.
func (s *FlatSet) Add(val int) {
	s.innerMap.Put(val, "")
}

// Remove removes an element from the set.
// If the element exists it is removed and the function returns true, otherwise it returns false.
func (s *FlatSet) Remove(val int) bool {
	return s.innerMap.Remove(val)
}

// Contains returns true if the element exists in the set, otherwise it returns false.
func (s *FlatSet) Contains(val int) bool {
	_, ok := s.innerMap.Get(val)
	return ok
}

// Size returns the number of elements in the set.
func (s *FlatSet) Size() int {
	return s.innerMap.Size()
}

func (s *FlatSet) IsEmpty() bool {
	return s.innerMap.IsEmpty()
}

func (s *FlatSet) IsNotEmpty() bool {
	return s.innerMap.IsNotEmpty()
}

func (s *FlatSet) Clear() {
	s.innerMap.Clear()
}

func (s *FlatSet) Formatted() string {
	//TODO implement me
	panic("implement me")
}
