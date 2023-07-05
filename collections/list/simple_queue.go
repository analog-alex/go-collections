package list

// Make this based on our implementation of a linked list later

type SimpleQueue struct {
	elements []int
}

func MakeSimpleQueue() *SimpleQueue {
	return &SimpleQueue{elements: make([]int, 0)}
}

func (q *SimpleQueue) Enqueue(val int) {
	q.elements = append(q.elements, val)
}

func (q *SimpleQueue) Dequeue() (int, bool) {
	if q.IsEmpty() {
		return 0, false
	}

	val := q.elements[0]
	q.elements = q.elements[1:]
	return val, true
}

func (q *SimpleQueue) Peek() (int, bool) {
	if q.IsEmpty() {
		return 0, false
	}

	return q.elements[0], true
}

func (q *SimpleQueue) Size() int {
	return len(q.elements)
}

func (q *SimpleQueue) Clear() {
	q.elements = make([]int, 0)
}

func (q *SimpleQueue) IsEmpty() bool {
	return len(q.elements) == 0
}

func (q *SimpleQueue) IsNotEmpty() bool {
	return len(q.elements) > 0
}

func (q *SimpleQueue) Formatted() string {
	// TODO print out the list in fifo order
	return ""
}
