package list

// Make this based on our implementation of a linked list later

type SimpleStack struct {
	elements []int
}

func MakeSimpleStack() *SimpleStack {
	return &SimpleStack{elements: make([]int, 0)}
}

func (s *SimpleStack) Push(val int) {
	s.elements = append(s.elements, val)
}

func (s *SimpleStack) Pop() (int, bool) {
	if s.IsEmpty() {
		return 0, false
	}

	val := s.elements[len(s.elements)-1]
	s.elements = s.elements[:len(s.elements)-1]
	return val, true
}

func (s *SimpleStack) Peek() (int, bool) {
	if s.IsEmpty() {
		return 0, false
	}

	return s.elements[len(s.elements)-1], true
}

func (s *SimpleStack) Size() int {
	return len(s.elements)
}

func (s *SimpleStack) Clear() {
	s.elements = make([]int, 0)
}

func (s *SimpleStack) IsEmpty() bool {
	return len(s.elements) == 0
}

func (s *SimpleStack) IsNotEmpty() bool {
	return len(s.elements) > 0
}

func (s *SimpleStack) Formatted() string {
	// TODO print out the list in fifo order
	return ""
}
