package list

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLinkedList_Add(t *testing.T) {
	var l List[int] = MakeLinkedList[int]()
	l.Add(1)
	l.Add(2)
	l.Add(3)

	assert.Equal(t, 3, l.Size())
}

func TestLinkedList_Head(t *testing.T) {
	var l List[int] = MakeLinkedList[int]()
	l.Add(1)

	assert.Equal(t, 1, l.Size())
}

func TestLinkedList_Size(t *testing.T) {
	var l List[int] = MakeLinkedList[int]()
	l.Add(1)
	l.Add(2)
	l.Add(3)

	assert.Equal(t, 3, l.Size())
}

func TestLinkedList_Get(t *testing.T) {
	var l List[int] = MakeLinkedList[int]()
	l.Add(1)

	val, ok := l.Get(0)
	assert.True(t, ok)
	assert.Equal(t, 1, val)

	val, ok = l.Get(1)
	assert.False(t, ok)
	assert.Equal(t, 0, val)
}

func TestLinkedList_Remove(t *testing.T) {
	var l List[int] = MakeLinkedList[int]()
	l.Add(1)
	l.Add(2)
	l.Add(3)

	assert.True(t, l.Remove(2))
	assert.False(t, l.Remove(5))
	assert.Equal(t, 2, l.Size())
}

func TestLinkedList_RemoveHead(t *testing.T) {
	var l List[int] = MakeLinkedList[int]()
	l.Add(1)

	assert.True(t, l.Remove(1))
	assert.Equal(t, 0, l.Size())
	assert.True(t, l.IsEmpty())
}

func TestLinkedList_TailInvariantConservedAfterRemoval(t *testing.T) {
	var l List[int] = MakeLinkedList[int]()
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

func TestLinkedList_Contains(t *testing.T) {
	var l List[int] = MakeLinkedList[int]()
	l.Add(1)
	l.Add(2)
	l.Add(3)

	assert.True(t, l.Contains(2))
	assert.False(t, l.Contains(5))
	assert.Equal(t, 3, l.Size())
}

func TestLinkedList_IsEmpty(t *testing.T) {
	var l List[int] = MakeLinkedList[int]()
	assert.True(t, l.IsEmpty())

	l.Add(1)
	assert.False(t, l.IsEmpty())
}

func TestLinkedList_IsNotEmpty(t *testing.T) {
	var l List[string] = MakeLinkedList[string]()
	assert.False(t, l.IsNotEmpty())

	l.Add("pi")
	assert.True(t, l.IsNotEmpty())
}
