package set

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBinaryTreeSet_Add(t *testing.T) {
	var s Set = MakeBinaryTreeSet()
	s.Add(1)
	s.Add(2)
	s.Add(3)

	assert.Equal(t, 3, s.Size())
}

func TestBinaryTreeSet_Size(t *testing.T) {
	var s Set = MakeBinaryTreeSet()
	s.Add(1)
	s.Add(2)

	assert.Equal(t, 2, s.Size())
}

func TestBinaryTreeSet_IsEmpty(t *testing.T) {
	var s Set = MakeBinaryTreeSet()

	assert.True(t, s.IsEmpty())
}

func TestBinaryTreeSet_IsNotEmpty(t *testing.T) {
	var s Set = MakeBinaryTreeSet()
	s.Add(1)

	assert.False(t, s.IsEmpty())
}

func TestBinaryTreeSet_Contains(t *testing.T) {
	var s Set = MakeBinaryTreeSet()
	s.Add(1)

	assert.True(t, s.Contains(1))
}

func TestBinaryTreeSet_DoesNotContain(t *testing.T) {
	var s Set = MakeBinaryTreeSet()
	s.Add(1)

	assert.False(t, s.Contains(2))
}

func TestBinaryTreeSet_Remove(t *testing.T) {
	var s Set = MakeBinaryTreeSet()
	s.Add(1)
	s.Add(2)

	assert.Equal(t, 2, s.Size())

	ok := s.Remove(2)
	assert.True(t, ok)
	assert.Equal(t, 1, s.Size())
}

func TestBinaryTreeSet_RemoveNonExistent(t *testing.T) {
	var s Set = MakeBinaryTreeSet()
	s.Add(1)
	s.Add(2)

	assert.Equal(t, 2, s.Size())

	ok := s.Remove(3)
	assert.False(t, ok)
	assert.Equal(t, 2, s.Size())
}

// ordered set operations

func TestBinaryTreeSet_First(t *testing.T) {
	var s OrderedSet = MakeBinaryTreeSet()
	s.Add(1)
	s.Add(2)

	assert.Equal(t, 1, s.First())
}

func TestBinaryTreeSet_Last(t *testing.T) {
	var s OrderedSet = MakeBinaryTreeSet()
	s.Add(1)
	s.Add(2)

	assert.Equal(t, 2, s.Last())
}

func TestBinaryTreeSet_RemoveFirst(t *testing.T) {
	var s OrderedSet = MakeBinaryTreeSet()
	s.Add(1)
	s.Add(2)

	assert.Equal(t, 2, s.Size())

	ok := s.RemoveFirst()
	assert.True(t, ok)
	assert.Equal(t, 1, s.Size())
	assert.Equal(t, 2, s.First())
}

func TestBinaryTreeSet_RemoveLast(t *testing.T) {
	var s OrderedSet = MakeBinaryTreeSet()
	s.Add(1)
	s.Add(2)

	assert.Equal(t, 2, s.Size())

	ok := s.RemoveLast()
	assert.True(t, ok)
	assert.Equal(t, 1, s.Size())
	assert.Equal(t, 1, s.Last())
}
