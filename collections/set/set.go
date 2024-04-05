package set

import "utils-generics/collections"

type Set interface {
	collections.Collection
	Add(val int)
	Remove(val int) bool
	Contains(val int) bool
}

type OrderedSet interface {
	Set
	First() int
	Last() int
	RemoveFirst() bool
	RemoveLast() bool
	ToSortedSlice() []int
}
