package set

import "utils-generics/collections"

type Set[K any] interface {
	collections.Collection
	Add(val K)
	Remove(val K) bool
	Contains(val K) bool
}

type OrderedSet[K any] interface {
	Set[K]
	First() K
	Last() K
	RemoveFirst() bool
	RemoveLast() bool
	ToSortedSlice() []K
}
