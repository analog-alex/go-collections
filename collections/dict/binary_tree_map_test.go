package dict

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBinaryTreeMap_Put(t *testing.T) {
	var m Map[string, string] = MakeBinaryTreeMap[string, string](stringComparator)
	m.Put("1", "one")
	m.Put("2", "two")
	m.Put("3", "three")

	assert.Equal(t, 3, m.Size())
}

func TestBinaryTreeMap_Size(t *testing.T) {
	var m Map[string, string] = MakeBinaryTreeMap[string, string](stringComparator)
	m.Put("1", "one")
	m.Put("2", "two")

	assert.Equal(t, 2, m.Size())
}

func TestBinaryTreeMap_IsEmpty(t *testing.T) {
	var m Map[string, string] = MakeBinaryTreeMap[string, string](stringComparator)

	assert.True(t, m.IsEmpty())
}

func TestBinaryTreeMap_IsNotEmpty(t *testing.T) {
	var m Map[string, string] = MakeBinaryTreeMap[string, string](stringComparator)
	m.Put("1", "one")

	assert.False(t, m.IsEmpty())
}

func TestBinaryTreeMap_ContainsKey(t *testing.T) {
	var m Map[string, string] = MakeBinaryTreeMap[string, string](stringComparator)
	m.Put("1", "one")
	m.Put("2", "two")

	assert.True(t, m.ContainsKey("1"))
}

func TestBinaryTreeMap_DoesNotContain(t *testing.T) {
	var m Map[string, string] = MakeBinaryTreeMap[string, string](stringComparator)
	m.Put("1", "one")

	assert.False(t, m.ContainsKey("2"))
}

func TestBinaryTreeMap_Remove(t *testing.T) {
	var m Map[string, string] = MakeBinaryTreeMap[string, string](stringComparator)
	m.Put("1", "one")
	m.Put("2", "two")

	assert.Equal(t, 2, m.Size())

	ok := m.Remove("2")
	assert.True(t, ok)
	assert.Equal(t, 1, m.Size())
}

func TestBinaryTreeMap_RemoveNonExistent(t *testing.T) {
	var m Map[string, string] = MakeBinaryTreeMap[string, string](stringComparator)
	m.Put("1", "one")
	m.Put("2", "two")

	assert.Equal(t, 2, m.Size())

	ok := m.Remove("3")
	assert.False(t, ok)
	assert.Equal(t, 2, m.Size())
}

func TestBinaryTreeMap_Get(t *testing.T) {
	var m Map[string, string] = MakeBinaryTreeMap[string, string](stringComparator)
	m.Put("1", "one")
	m.Put("2", "two")

	val, ok := m.Get("1")
	assert.True(t, ok)
	assert.Equal(t, "one", val)

	val, ok = m.Get("3")
	assert.False(t, ok)
	assert.Equal(t, "", val)
}
