package set

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHashTableSet_Add(t *testing.T) {
	var s Set = MakeHashSet()
	s.Add(1)
	s.Add(2)
	s.Add(3)

	assert.Equal(t, 3, s.Size())
}

func TestHashTableSet_Size(t *testing.T) {
	var s Set = MakeHashSet()
	s.Add(1)
	s.Add(2)

	assert.Equal(t, 2, s.Size())
}

func TestHashTableSet_IsEmpty(t *testing.T) {
	var s Set = MakeHashSet()

	assert.True(t, s.IsEmpty())
}

func TestHashTableSet_IsNotEmpty(t *testing.T) {
	var s Set = MakeHashSet()
	s.Add(1)

	assert.False(t, s.IsEmpty())
}

func TestHashTableSet_Contains(t *testing.T) {
	var s Set = MakeHashSet()
	s.Add(1)

	assert.True(t, s.Contains(1))
}

func TestHashTableSet_DoesNotContain(t *testing.T) {
	var s Set = MakeHashSet()
	s.Add(1)

	assert.False(t, s.Contains(2))
}

func TestHashTableSet_Remove(t *testing.T) {
	var s Set = MakeHashSet()
	s.Add(1)
	s.Add(2)

	assert.Equal(t, 2, s.Size())

	ok := s.Remove(2)
	assert.True(t, ok)
	assert.Equal(t, 1, s.Size())
}

func TestHashTableSet_RemoveNonExistent(t *testing.T) {
	var s Set = MakeHashSet()
	s.Add(1)
	s.Add(2)

	assert.Equal(t, 2, s.Size())

	ok := s.Remove(3)
	assert.False(t, ok)
	assert.Equal(t, 2, s.Size())
}
