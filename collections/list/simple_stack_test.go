package list

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSimpleStack_Enqueue(t *testing.T) {
	var s Stack[int] = MakeSimpleStack[int]()
	s.Push(1)
	s.Push(2)
	s.Push(3)

	assert.Equal(t, 3, s.Size())
}

func TestSimpleStack_Size(t *testing.T) {
	var s Stack[int] = MakeSimpleStack[int]()
	s.Push(1)
	s.Push(2)

	assert.Equal(t, 2, s.Size())
}

func TestSimpleStack_IsEmpty(t *testing.T) {
	var s Stack[int] = MakeSimpleStack[int]()

	assert.True(t, s.IsEmpty())
}

func TestSimpleStack_IsNotEmpty(t *testing.T) {
	var s Stack[int] = MakeSimpleStack[int]()
	s.Push(1)

	assert.False(t, s.IsEmpty())
}

func TestSimpleStack_Dequeue(t *testing.T) {
	var s Stack[int] = MakeSimpleStack[int]()
	s.Push(1)

	val, ok := s.Pop()
	assert.True(t, ok)
	assert.Equal(t, 1, val)
	assert.Equal(t, 0, s.Size())
	assert.True(t, s.IsEmpty())
}

func TestSimpleStack_Peek(t *testing.T) {
	var s Stack[int] = MakeSimpleStack[int]()
	s.Push(1)

	val, ok := s.Peek()
	assert.True(t, ok)
	assert.Equal(t, 1, val)
	assert.Equal(t, 1, s.Size())
	assert.False(t, s.IsEmpty())
}
