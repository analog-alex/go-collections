package dict

import "utils-generics/collections"

type Entry[T any] struct {
	Key int
	Val T
}

type Map[T any] interface {
	collections.Collection
	Get(key int) (T, bool)
	Put(key int, val T)
	Remove(key int) bool
	ContainsKey(key int) bool

	Entries() []Entry[T]
	Keys() []int
	Values() []T
}

type OrderedMap[T any] interface {
	Map[T]
	First() int
	Last() int
	RemoveFirst() bool
	RemoveLast() bool
}
