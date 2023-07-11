package list

// Make this based on our implementation of a linked list later

type SimpleQueue[T any] struct {
	elements []T
}

func MakeSimpleQueue[T any]() *SimpleQueue[T] {
	return &SimpleQueue[T]{elements: make([]T, 0)}
}

func (q *SimpleQueue[T]) Enqueue(val T) {
	q.elements = append(q.elements, val)
}

func (q *SimpleQueue[T]) Dequeue() (T, bool) {
	if q.IsEmpty() {
		var zero T
		return zero, false
	}

	val := q.elements[0]
	q.elements = q.elements[1:]
	return val, true
}

func (q *SimpleQueue[T]) Peek() (T, bool) {
	if q.IsEmpty() {
		var zero T
		return zero, false
	}

	return q.elements[0], true
}

func (q *SimpleQueue[T]) Size() int {
	return len(q.elements)
}

func (q *SimpleQueue[T]) Clear() {
	q.elements = make([]T, 0)
}

func (q *SimpleQueue[T]) IsEmpty() bool {
	return len(q.elements) == 0
}

func (q *SimpleQueue[T]) IsNotEmpty() bool {
	return len(q.elements) > 0
}

func (q *SimpleQueue[T]) Formatted() string {
	// TODO print out the list in fifo order
	return ""
}
