package collections

type Collection interface {
	IsEmpty() bool
	IsNotEmpty() bool
	Size() int
	Clear()
	Formatted() string
}
