package list

import "io.analogalex.collections/collections"

type List interface {
	collections.Collection
	Add(val int)
	Remove(val int) bool
	Get(index int) (int, bool)
	Contains(val int) bool
}

type Queue interface {
	collections.Collection
	Enqueue(val int)
	Dequeue() (int, bool)
	Peek() (int, bool)
}

type Stack interface {
	collections.Collection
	Push(val int)
	Pop() (int, bool)
	Peek() (int, bool)
}
