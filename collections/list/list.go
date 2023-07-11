package list

import "io.analogalex.collections/collections"

type List[T any] interface {
	collections.Collection
	Add(val T)
	Remove(val T) bool
	Get(index int) (T, bool)
	Contains(val T) bool
	// TODO Sort(func (T, T) bool)
}

type Queue[T any] interface {
	collections.Collection
	Enqueue(val T)
	Dequeue() (T, bool)
	Peek() (T, bool)
}

type Stack[T any] interface {
	collections.Collection
	Push(val T)
	Pop() (T, bool)
	Peek() (T, bool)
}
