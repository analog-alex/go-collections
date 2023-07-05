package dict

import "io.analogalex.collections/collections"

type Entry struct {
	Key int
	Val string
}

type Map interface {
	collections.Collection
	Get(key int) (string, bool)
	Put(key int, val string)
	Remove(key int) bool
	ContainsKey(key int) bool

	Entries() []Entry
	Keys() []int
	Values() []string
}

type OrderedMap interface {
	Map
	First() int
	Last() int
	RemoveFirst() bool
	RemoveLast() bool
}
