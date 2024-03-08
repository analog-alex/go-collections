package dict

import "fmt"

const defaultHashTableSize = 128

type hashTableEntry[T any] struct {
	Entry[T]
	next *hashTableEntry[T]
}

type HashMap[T any] struct {
	table [defaultHashTableSize]*hashTableEntry[T]
}

func MakeHashMap[T any]() *HashMap[T] {
	return &HashMap[T]{}
}

// Put adds a new entry to the map.
//
// If an entry with the key already exists the value is updated with the one provided
func (s *HashMap[T]) Put(key int, val T) {
	hash := key % defaultHashTableSize
	if s.table[hash] == nil {
		s.table[hash] = &hashTableEntry[T]{Entry: Entry[T]{Key: key, Val: val}}
		return
	}

	node := s.table[hash]
	for {
		if node.Key == key {
			node.Val = val
			return
		}
		if node.next == nil {
			node.next = &hashTableEntry[T]{Entry: Entry[T]{Key: key, Val: val}}
			return
		}
		node = node.next
	}
}

// Remove removes the entry identified by the key, returning true if the entry was found and removed
// and false if the entry was not found
func (s *HashMap[T]) Remove(key int) bool {
	hash := key % defaultHashTableSize
	if s.table[hash] == nil {
		return false
	}

	node := s.table[hash]
	var prev *hashTableEntry[T]
	for {
		if node.Key == key {
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
func (s *HashMap[T]) ContainsKey(key int) bool {
	hash := key % defaultHashTableSize
	if s.table[hash] == nil {
		return false
	}

	node := s.table[hash]
	for {
		if node.Key == key {
			return true
		}
		if node.next == nil {
			return false
		}
		node = node.next
	}
}

// Get returns the value associated with the provided key and true if the key was found and false if otherwise
func (s *HashMap[T]) Get(key int) (T, bool) {
	var zero T
	hash := key % defaultHashTableSize
	if s.table[hash] == nil {
		return zero, false
	}

	node := s.table[hash]
	for {
		if node.Key == key {
			return node.Val.(T), true
		}
		if node.next == nil {
			return zero, false
		}
		node = node.next
	}
}

// Size returns the number of entries in the map
func (s *HashMap[T]) Size() int {
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
func (s *HashMap[T]) IsEmpty() bool {
	return s.Size() == 0
}

// IsNotEmpty returns true if the map is not empty and false if otherwise
func (s *HashMap[T]) IsNotEmpty() bool {
	return !s.IsEmpty()
}

// Formatted returns a string representation of the map
func (s *HashMap[T]) Formatted() string {
	str := "{"
	for _, node := range s.table {
		if node != nil {
			str += fmt.Sprintf("%d: %s, ", node.Key, node.Val)
			for node.next != nil {
				str += ", " + fmt.Sprintf("%d: %s, ", node.Key, node.Val)
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
func (s *HashMap[T]) Clear() {
	s.table = [defaultHashTableSize]*hashTableEntry[T]{}
}

// Entries returns a slice of all entries in the map
func (s *HashMap[T]) Entries() []Entry[T] {
	entries := make([]Entry[T], 0, s.Size())
	for _, node := range s.table {
		if node != nil {
			entries = append(entries, Entry[T]{Key: node.Key, Val: node.Val})
			for node.next != nil {
				entries = append(entries, Entry[T]{Key: node.Key, Val: node.Val})
				node = node.next
			}
		}
	}
	return entries
}

// Keys returns a slice of all keys in the map
func (s *HashMap[T]) Keys() []int {
	keys := make([]int, 0, s.Size())
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
func (s *HashMap[T]) Values() []T {
	values := make([]T, 0, s.Size())
	for _, node := range s.table {
		if node != nil {
			values = append(values, node.Val.(T))
			for node.next != nil {
				values = append(values, node.Val.(T))
				node = node.next
			}
		}
	}
	return values
}
