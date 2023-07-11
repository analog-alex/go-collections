package list

// Make this based on our implementation of a linked list later

type SimpleStack[T any] struct {
	elements []T
}

func MakeSimpleStack[T any]() *SimpleStack[T] {
	return &SimpleStack[T]{elements: make([]T, 0)}
}

func (s *SimpleStack[T]) Push(val T) {
	s.elements = append(s.elements, val)
}

func (s *SimpleStack[T]) Pop() (T, bool) {
	if s.IsEmpty() {
		var zero T
		return zero, false
	}

	val := s.elements[len(s.elements)-1]
	s.elements = s.elements[:len(s.elements)-1]
	return val, true
}

func (s *SimpleStack[T]) Peek() (T, bool) {
	if s.IsEmpty() {
		var zero T
		return zero, false
	}

	return s.elements[len(s.elements)-1], true
}

func (s *SimpleStack[T]) Size() int {
	return len(s.elements)
}

func (s *SimpleStack[T]) Clear() {
	s.elements = make([]T, 0)
}

func (s *SimpleStack[T]) IsEmpty() bool {
	return len(s.elements) == 0
}

func (s *SimpleStack[T]) IsNotEmpty() bool {
	return len(s.elements) > 0
}

func (s *SimpleStack[T]) Formatted() string {
	// TODO print out the list in fifo order
	return ""
}
