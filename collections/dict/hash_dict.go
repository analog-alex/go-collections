package dict

import "fmt"

const defaultHashTableSize = 128

type hashTableEntry struct {
	key  int
	val  string
	next *hashTableEntry
}

type HashMap struct {
	table [defaultHashTableSize]*hashTableEntry
}

func MakeHashMap() *HashMap {
	return &HashMap{}
}

// Put adds a new entry to the map.
//
// If an entry with the key already exists the value is updated with the one provided
func (s *HashMap) Put(key int, val string) {
	hash := key % defaultHashTableSize
	if s.table[hash] == nil {
		s.table[hash] = &hashTableEntry{key: key, val: val}
		return
	}

	node := s.table[hash]
	for {
		if node.key == key {
			node.val = val
			return
		}
		if node.next == nil {
			node.next = &hashTableEntry{key: key, val: val}
			return
		}
		node = node.next
	}
}

// Remove removes the entry identified by the key, returning true if the entry was found and removed
// and false if the entry was not found
func (s *HashMap) Remove(key int) bool {
	hash := key % defaultHashTableSize
	if s.table[hash] == nil {
		return false
	}

	node := s.table[hash]
	var prev *hashTableEntry
	for {
		if node.key == key {
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
func (s *HashMap) ContainsKey(key int) bool {
	hash := key % defaultHashTableSize
	if s.table[hash] == nil {
		return false
	}

	node := s.table[hash]
	for {
		if node.key == key {
			return true
		}
		if node.next == nil {
			return false
		}
		node = node.next
	}
}

// Get returns the value associated with the provided key and true if the key was found and false if otherwise
func (s *HashMap) Get(key int) (string, bool) {
	hash := key % defaultHashTableSize
	if s.table[hash] == nil {
		return "", false
	}

	node := s.table[hash]
	for {
		if node.key == key {
			return node.val, true
		}
		if node.next == nil {
			return "", false
		}
		node = node.next
	}
}

// Size returns the number of entries in the map
func (s *HashMap) Size() int {
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
func (s *HashMap) IsEmpty() bool {
	return s.Size() == 0
}

// IsNotEmpty returns true if the map is not empty and false if otherwise
func (s *HashMap) IsNotEmpty() bool {
	return !s.IsEmpty()
}

// Formatted returns a string representation of the map
func (s *HashMap) Formatted() string {
	str := "{"
	for _, node := range s.table {
		if node != nil {
			str += fmt.Sprintf("%d: %s, ", node.key, node.val)
			for node.next != nil {
				str += ", " + fmt.Sprintf("%d: %s, ", node.key, node.val)
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
func (s *HashMap) Clear() {
	s.table = [defaultHashTableSize]*hashTableEntry{}
}

// Entries returns a slice of all entries in the map
func (s *HashMap) Entries() []Entry {
	entries := make([]Entry, 0, s.Size())
	for _, node := range s.table {
		if node != nil {
			entries = append(entries, Entry{Key: node.key, Val: node.val})
			for node.next != nil {
				entries = append(entries, Entry{Key: node.key, Val: node.val})
				node = node.next
			}
		}
	}
	return entries
}

// Keys returns a slice of all keys in the map
func (s *HashMap) Keys() []int {
	keys := make([]int, 0, s.Size())
	for _, node := range s.table {
		if node != nil {
			keys = append(keys, node.key)
			for node.next != nil {
				keys = append(keys, node.key)
				node = node.next
			}
		}
	}
	return keys
}

// Values returns a slice of all values in the map
func (s *HashMap) Values() []string {
	values := make([]string, 0, s.Size())
	for _, node := range s.table {
		if node != nil {
			values = append(values, node.val)
			for node.next != nil {
				values = append(values, node.val)
				node = node.next
			}
		}
	}
	return values
}
