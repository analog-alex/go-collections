package dict

import (
	"fmt"
	"reflect"
)

const defaultHashTableSize = 128

type hashTableEntry[K any, T any] struct {
	Entry[K, T]
	next *hashTableEntry[K, T]
}

type HashMap[K any, T any] struct {
	table  [defaultHashTableSize]*hashTableEntry[K, T]
	hasher func(K) int
}

func MakeHashMap[K any, T any](h func(K) int) *HashMap[K, T] {
	return &HashMap[K, T]{hasher: h}
}

// Put adds a new entry to the map.
//
// If an entry with the key already exists the value is updated with the one provided
func (s *HashMap[K, T]) Put(key K, val T) {
	hash := s.hasher(key) % defaultHashTableSize
	node := s.table[hash]
	if node == nil {
		s.table[hash] = &hashTableEntry[K, T]{Entry: Entry[K, T]{Key: key, Val: val}}
		return
	}

	for {
		if reflect.DeepEqual(node.Key, key) {
			node.Val = val
			return
		}
		if node.next == nil {
			node.next = &hashTableEntry[K, T]{Entry: Entry[K, T]{Key: key, Val: val}}
			return
		}
		node = node.next
	}
}

// Remove removes the entry identified by the key, returning true if the entry was found and removed
// and false if the entry was not found
func (s *HashMap[K, T]) Remove(key K) bool {
	hash := s.hasher(key) % defaultHashTableSize
	node := s.table[hash]
	if node == nil {
		return false
	}

	var prev *hashTableEntry[K, T]
	for {
		if reflect.DeepEqual(node.Key, key) {
			if prev == nil {
				s.table[hash] = node.next
			} else {
				prev.next = node.next
			}
			return true
		}
		if node.next == nil {
			return false
		}
		prev = node
		node = node.next
	}
}

// ContainsKey returns true if the map contains an entry with the provided key and false if otherwise
func (s *HashMap[K, T]) ContainsKey(key K) bool {
	hash := s.hasher(key) % defaultHashTableSize
	node := s.table[hash]
	if node == nil {
		return false
	}

	for {
		if reflect.DeepEqual(node.Key, key) {
			return true
		}
		if node.next == nil {
			return false
		}
		node = node.next
	}
}

// Get returns the value associated with the provided key and true if the key was found and false if otherwise
func (s *HashMap[K, T]) Get(key K) (T, bool) {
	var zero T
	hash := s.hasher(key) % defaultHashTableSize
	node := s.table[hash]
	if node == nil {
		return zero, false
	}

	for {
		if reflect.DeepEqual(node.Key, key) {
			return node.Val, true
		}
		if node.next == nil {
			return zero, false
		}
		node = node.next
	}
}

// Size returns the number of entries in the map
func (s *HashMap[K, T]) Size() int {
	size := 0
	for _, node := range s.table {
		if node != nil {
			size++
			for node.next != nil {
				size++
				node = node.next
			}
		}
	}
	return size
}

// IsEmpty returns true if the map is empty and false if otherwise
func (s *HashMap[K, T]) IsEmpty() bool {
	return s.Size() == 0
}

// IsNotEmpty returns true if the map is not empty and false if otherwise
func (s *HashMap[K, T]) IsNotEmpty() bool {
	return !s.IsEmpty()
}

// Formatted returns a string representation of the map
func (s *HashMap[K, T]) Formatted() string {
	str := "{"
	for _, node := range s.table {
		if node != nil {
			str += fmt.Sprintf("%T: %T, ", node.Key, node.Val)
			for node.next != nil {
				str += ", " + fmt.Sprintf("%T: %T, ", node.Key, node.Val)
				node = node.next
			}
		}
	}
	// remove last comma and space str if it exists
	if len(str) > 1 {
		str = str[:len(str)-2]
	}
	str += "}"
	return str
}

// Clear removes all entries from the map
func (s *HashMap[K, T]) Clear() {
	s.table = [defaultHashTableSize]*hashTableEntry[K, T]{}
}

// Entries returns a slice of all entries in the map
func (s *HashMap[K, T]) Entries() []Entry[K, T] {
	entries := make([]Entry[K, T], 0, s.Size())
	for _, node := range s.table {
		if node != nil {
			entries = append(entries, Entry[K, T]{Key: node.Key, Val: node.Val})
			for node.next != nil {
				entries = append(entries, Entry[K, T]{Key: node.Key, Val: node.Val})
				node = node.next
			}
		}
	}
	return entries
}

// Keys returns a slice of all keys in the map
func (s *HashMap[K, T]) Keys() []K {
	keys := make([]K, 0, s.Size())
	for _, node := range s.table {
		if node != nil {
			keys = append(keys, node.Key)
			for node.next != nil {
				keys = append(keys, node.Key)
				node = node.next
			}
		}
	}
	return keys
}

// Values returns a slice of all values in the map
func (s *HashMap[K, T]) Values() []T {
	values := make([]T, 0, s.Size())
	for _, node := range s.table {
		if node != nil {
			values = append(values, node.Val)
			for node.next != nil {
				values = append(values, node.Val)
				node = node.next
			}
		}
	}
	return values
}
