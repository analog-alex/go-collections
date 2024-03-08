package dict

type flatMapEntry struct {
	key int
	val string
}

type FlatMap struct {
	array []flatMapEntry
}

func MakeFlatMap() *FlatMap {
	return &FlatMap{}
}

/*
* Binary search for the key in the array.
* We keep the array sorted by key, so we can use binary search to find the key.
 */
func (s *FlatMap) binarySearch(key int) (flatMapEntry, bool) {
	l, h := 0, len(s.array)-1
	for l <= h {
		m := (l + h) / 2
		if s.array[m].key == key {
			return s.array[m], true
		}

		if s.array[m].key < key {
			l = m + 1
		} else {
			h = m - 1
		}
	}

	return flatMapEntry{}, false
}

// Put inserts a new key-value pair into the map. If the key already exists, the value is updated.
func (s *FlatMap) Put(key int, val string) {
	s.array = append(s.array, flatMapEntry{key: key, val: val})

	for i := len(s.array) - 1; i > 0; i-- {
		// conserve the well ordering of the array elements
		if s.array[i].key < s.array[i-1].key {
			s.array[i], s.array[i-1] = s.array[i-1], s.array[i]
		} else {
			break
		}
	}
}

func (s *FlatMap) Remove(key int) bool {
	for i, entry := range s.array {
		if entry.key == key {
			s.array = append(s.array[:i], s.array[i+1:]...)
			return true
		}
	}

	return false
}

func (s *FlatMap) Get(key int) (string, bool) {
	l, ok := s.binarySearch(key)
	if !ok {
		return "", false
	}
	return l.val, true
}

func (s *FlatMap) ContainsKey(key int) bool {
	_, ok := s.binarySearch(key)
	return ok
}

func (s *FlatMap) Size() int {
	return len(s.array)
}

func (s *FlatMap) IsEmpty() bool {
	return len(s.array) == 0
}

func (s *FlatMap) IsNotEmpty() bool {
	return !s.IsEmpty()
}

func (s *FlatMap) Clear() {
	s.array = []flatMapEntry{}
}

func (s *FlatMap) Formatted() string {
	//TODO implement me
	panic("implement me")
}

func (s *FlatMap) Entries() []Entry {
	var entries []Entry
	for _, entry := range s.array {
		entries = append(entries, Entry{Key: entry.key, Val: entry.val})
	}
	return entries
}

func (s *FlatMap) Keys() []int {
	var keys []int
	for _, entry := range s.array {
		keys = append(keys, entry.key)
	}

	return keys
}

func (s *FlatMap) Values() []string {
	var values []string
	for _, entry := range s.array {
		values = append(values, entry.val)
	}

	return values
}
