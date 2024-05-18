package dict

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"utils-generics/collections/types"
)

func TestFlatMap_Put(t *testing.T) {
	var m Map[string, string] = MakeFlatMap[string, string](types.StringComparator)
	m.Put("1", "one")
	m.Put("2", "two")
	m.Put("3", "three")

	assert.Equal(t, 3, m.Size())
}

func TestFlatMap_Put_ConservesOrder(t *testing.T) {
	var m Map[string, string] = MakeFlatMap[string, string](types.StringComparator)
	m.Put("1", "one")
	m.Put("3", "three")
	m.Put("2", "two")

	assert.Equal(t, 3, m.Size())

	entries := m.Entries()

	assert.Equal(t, "1", entries[0].Key)
	assert.Equal(t, "2", entries[1].Key)
	assert.Equal(t, "3", entries[2].Key)
}

func TestFlatMap_Size(t *testing.T) {
	var m Map[string, string] = MakeFlatMap[string, string](types.StringComparator)
	m.Put("1", "one")

	assert.Equal(t, 1, m.Size())
}

func TestFlatMap_IsEmpty(t *testing.T) {
	var m Map[string, string] = MakeFlatMap[string, string](types.StringComparator)

	assert.True(t, m.IsEmpty())
}

func TestFlatMap_IsNotEmpty(t *testing.T) {
	var m Map[string, string] = MakeFlatMap[string, string](types.StringComparator)
	m.Put("1", "one")

	assert.False(t, m.IsEmpty())
}

func TestFatMap_ContainsKey(t *testing.T) {
	var m Map[string, string] = MakeFlatMap[string, string](types.StringComparator)
	m.Put("1", "one")
	m.Put("2", "two")

	assert.True(t, m.ContainsKey("1"))
	assert.False(t, m.ContainsKey("3"))
}

func TestFlatMap_Remove(t *testing.T) {
	var m Map[string, string] = MakeFlatMap[string, string](types.StringComparator)
	m.Put("1", "one")
	m.Put("2", "two")

	assert.Equal(t, 2, m.Size())

	ok := m.Remove("2")
	assert.True(t, ok)
	assert.Equal(t, 1, m.Size())
}

func TestFlatMap_RemoveNonExistent(t *testing.T) {
	var m Map[string, string] = MakeFlatMap[string, string](types.StringComparator)
	m.Put("1", "one")
	m.Put("2", "two")

	assert.Equal(t, 2, m.Size())

	ok := m.Remove("3")
	assert.False(t, ok)
	assert.Equal(t, 2, m.Size())
}

func TestFlatMap_Get(t *testing.T) {
	var m Map[string, string] = MakeFlatMap[string, string](types.StringComparator)
	m.Put("1", "one")
	m.Put("2", "two")

	val, ok := m.Get("1")

	assert.True(t, ok)
	assert.Equal(t, "one", val)
}

func TestFlatMap_GetNonExistent(t *testing.T) {
	var m Map[string, string] = MakeFlatMap[string, string](types.StringComparator)
	m.Put("1", "one")
	m.Put("2", "two")

	val, ok := m.Get("3")

	assert.False(t, ok)
	assert.Equal(t, "", val)
}
