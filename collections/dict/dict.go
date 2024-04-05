package dict

import "utils-generics/collections"

// Entry is a key-value pair
// K and T are generic types that parametrize the entry
// For different map types the user must provide either a Hash function or
// a compare function to allow the underlying maps to provide functionality
// for all kinds of possible types
// For out-of-the-box types, functions will be provided by the lib
type Entry[K any, T any] struct {
	Key K
	Val T
}

type Map[K any, T any] interface {
	collections.Collection
	Get(key K) (T, bool)
	Put(key K, val T)
	Remove(key K) bool
	ContainsKey(key K) bool

	Entries() []Entry[K, T]
	Keys() []K
	Values() []T
}

type OrderedMap[K any, T any] interface {
	Map[K, T]
	First() K
	Last() K
	RemoveFirst() bool
	RemoveLast() bool
}
