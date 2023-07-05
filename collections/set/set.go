package set

import "io.analogalex.collections/collections"

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
