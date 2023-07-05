package list

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSimpleQueue_Enqueue(t *testing.T) {
	var q Queue = MakeSimpleQueue()
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	assert.Equal(t, 3, q.Size())
}

func TestSimpleQueue_Size(t *testing.T) {
	var q Queue = MakeSimpleQueue()
	q.Enqueue(1)
	q.Enqueue(2)

	assert.Equal(t, 2, q.Size())
}

func TestSimpleQueue_IsEmpty(t *testing.T) {
	var q Queue = MakeSimpleQueue()

	assert.True(t, q.IsEmpty())
}

func TestSimpleQueue_IsNotEmpty(t *testing.T) {
	var q Queue = MakeSimpleQueue()
	q.Enqueue(1)

	assert.False(t, q.IsEmpty())
}

func TestSimpleQueue_Dequeue(t *testing.T) {
	var q Queue = MakeSimpleQueue()
	q.Enqueue(1)

	val, ok := q.Dequeue()
	assert.True(t, ok)
	assert.Equal(t, 1, val)
	assert.Equal(t, 0, q.Size())
	assert.True(t, q.IsEmpty())
}

func TestSimpleQueue_Peek(t *testing.T) {
	var q Queue = MakeSimpleQueue()
	q.Enqueue(1)

	val, ok := q.Peek()
	assert.True(t, ok)
	assert.Equal(t, 1, val)
	assert.Equal(t, 1, q.Size())
	assert.False(t, q.IsEmpty())
}
