package list

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDoubleLinkedList_Add(t *testing.T) {
	var l List[int] = MakeDoubleLinkedList[int]()
	l.Add(1)
	l.Add(2)
	l.Add(3)

	assert.Equal(t, 3, l.Size())
}

func TestDoubleLinkedList_Head(t *testing.T) {
	var l List[int] = MakeDoubleLinkedList[int]()
	l.Add(1)

	assert.Equal(t, 1, l.Size())
}

func TestDoubleLinkedList_Size(t *testing.T) {
	var l List[int] = MakeDoubleLinkedList[int]()
	l.Add(1)
	l.Add(2)
	l.Add(3)

	assert.Equal(t, 3, l.Size())
}

func TestDoubleLinkedList_Get(t *testing.T) {
	var l List[int] = MakeDoubleLinkedList[int]()
	l.Add(1)

	val, ok := l.Get(0)
	assert.True(t, ok)
	assert.Equal(t, 1, val)

	val, ok = l.Get(1)
	assert.False(t, ok)
	assert.Equal(t, 0, val)
}

func TestDoubleLinkedList_Remove(t *testing.T) {
	var l List[int] = MakeDoubleLinkedList[int]()
	l.Add(1)
	l.Add(2)
	l.Add(3)

	assert.True(t, l.Remove(2))
	assert.False(t, l.Remove(5))
	assert.Equal(t, 2, l.Size())
}

func TestDoubleLinkedList_RemoveHead(t *testing.T) {
	var l List[int] = MakeDoubleLinkedList[int]()
	l.Add(1)

	assert.True(t, l.Remove(1))
	assert.Equal(t, 0, l.Size())
	assert.True(t, l.IsEmpty())
}

func TestDoubleLinkedList_TailInvariantConservedAfterRemoval(t *testing.T) {
	var l List[int] = MakeDoubleLinkedList[int]()
	l.Add(1)
	l.Add(2)
	l.Add(3)

	assert.Equal(t, 3, l.Size())

	l.Remove(3)

	assert.Equal(t, 2, l.Size())

	l.Add(4)
	l.Add(5)

	assert.Equal(t, 4, l.Size())

	val, ok := l.Get(3)
	assert.True(t, ok)
	assert.Equal(t, 5, val)
}

func TestDoubleLinkedList_Contains(t *testing.T) {
	var l List[int] = MakeDoubleLinkedList[int]()
	l.Add(1)
	l.Add(2)
	l.Add(3)

	assert.True(t, l.Contains(2))
	assert.False(t, l.Contains(5))
	assert.Equal(t, 3, l.Size())
}

func TestDoubleLinkedList_IsEmpty(t *testing.T) {
	var l List[int] = MakeDoubleLinkedList[int]()
	assert.True(t, l.IsEmpty())

	l.Add(1)
	assert.False(t, l.IsEmpty())
}

func TestDoubleLinkedList_IsNotEmpty(t *testing.T) {
	var l List[int] = MakeDoubleLinkedList[int]()
	assert.False(t, l.IsNotEmpty())

	l.Add(1)
	assert.True(t, l.IsNotEmpty())
}
