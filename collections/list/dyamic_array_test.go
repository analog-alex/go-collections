package list

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDynamicArray_Add(t *testing.T) {
	var l List[int] = MakeDynamicArray[int]()
	l.Add(1)
	l.Add(2)
	l.Add(3)

	assert.Equal(t, 3, l.Size())
}

func TestDynamicArray_Add2(t *testing.T) {
	var l List[int] = MakeDynamicArray[int]()
	l.Add(1)
	l.Add(2)
	l.Add(3)
	l.Add(4)
	l.Add(5)
	l.Add(6)
	l.Add(7)
	l.Add(8)
	l.Add(9)
	l.Add(10)
	l.Add(11)

	assert.Equal(t, 11, l.Size())
}

func TestDynamicArray_Add3(t *testing.T) {
	var l List[int] = MakeDynamicArrayWithCapacity[int](5)
	l.Add(1)
	l.Add(2)
	l.Add(3)
	l.Add(4)
	l.Add(5)
	l.Add(6)
	l.Add(7)
	l.Add(8)

	assert.Equal(t, 8, l.Size())
}

func TestDynamicArray_CreateWithValues(t *testing.T) {
	var l List[int] = MakeDynamicArrayWithValues[int](1, 2, 3, 4, 5)
	assert.Equal(t, 5, l.Size())

	val, ok := l.Get(2)
	assert.True(t, ok)
	assert.Equal(t, 3, val)
}

func TestDynamicArray_Clear(t *testing.T) {
	var l List[int] = MakeDynamicArray[int]()
	l.Add(1)
	l.Add(2)
	l.Add(3)

	l.Clear()
	assert.Equal(t, 0, l.Size())
}

func TestDynamicArray_Size(t *testing.T) {
	var l List[int] = MakeDynamicArray[int]()
	l.Add(1)
	l.Add(2)
	l.Add(3)

	assert.Equal(t, 3, l.Size())
}

func TestDynamicArray_Get(t *testing.T) {
	var l List[int] = MakeDynamicArray[int]()
	l.Add(1)

	val, ok := l.Get(0)
	assert.True(t, ok)
	assert.Equal(t, 1, val)

	val, ok = l.Get(1)
	assert.False(t, ok)
	assert.Equal(t, 0, val)
}

func TestDynamicArray_Remove(t *testing.T) {
	var l List[int] = MakeDynamicArray[int]()
	l.Add(1)
	l.Add(2)
	l.Add(3)

	assert.True(t, l.Remove(2))
	assert.False(t, l.Remove(5))
	assert.Equal(t, 2, l.Size())
}

func TestDynamicArray_Contains(t *testing.T) {
	var l List[int] = MakeDynamicArray[int]()
	l.Add(1)
	l.Add(2)
	l.Add(3)

	assert.True(t, l.Contains(2))
	assert.False(t, l.Contains(5))

}

func TestDynamicArray_IsEmpty(t *testing.T) {
	var l List[int] = MakeDynamicArray[int]()
	assert.True(t, l.IsEmpty())

	l.Add(1)
	assert.False(t, l.IsEmpty())
}

func TestDynamicArray_IsNotEmpty(t *testing.T) {
	var l List[int] = MakeDynamicArray[int]()
	l.Add(1)
	assert.True(t, l.IsNotEmpty())

	l.Remove(1)
	assert.False(t, l.IsNotEmpty())
}
