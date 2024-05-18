package dict

type FlatMap[K any, T any] struct {
	array      []Entry[K, T]
	comparator func(a, b K) int
}

func MakeFlatMap[K any, T any](c func(a, b K) int) *FlatMap[K, T] {
	return &FlatMap[K, T]{comparator: c}
}

/*
* Binary search for the key in the array.
* We keep the array sorted by key, so we can use binary search to find the key.
 */
func (s *FlatMap[K, T]) binarySearch(key K) (Entry[K, T], bool) {
	l, h := 0, len(s.array)-1
	for l <= h {
		m := (l + h) / 2

		c := s.comparator(s.array[m].Key, key)
		if c == 0 {
			return s.array[m], true
		}
		if c < 0 {
			l = m + 1
		} else {
			h = m - 1
		}
	}

	return Entry[K, T]{}, false
}

// Put inserts a new key-value pair into the map. If the key already exists, the value is updated.
func (s *FlatMap[K, T]) Put(key K, val T) {
	s.array = append(s.array, Entry[K, T]{Key: key, Val: val})

	for i := len(s.array) - 1; i > 0; i-- {
		// conserve the well ordering of the array elements
		c := s.comparator(s.array[i].Key, s.array[i-1].Key)
		if c == -1 {
			s.array[i], s.array[i-1] = s.array[i-1], s.array[i]
		} else {
			break
		}
	}
}

func (s *FlatMap[K, T]) Remove(key K) bool {
	for i, entry := range s.array {
		if s.comparator(entry.Key, key) == 0 {
			s.array = append(s.array[:i], s.array[i+1:]...)
			return true
		}
	}

	return false
}

func (s *FlatMap[K, T]) Get(key K) (T, bool) {
	l, ok := s.binarySearch(key)
	if !ok {
		var zero T
		return zero, false
	}
	return l.Val, true
}

func (s *FlatMap[K, T]) ContainsKey(key K) bool {
	_, ok := s.binarySearch(key)
	return ok
}

func (s *FlatMap[K, T]) Size() int {
	return len(s.array)
}

func (s *FlatMap[K, T]) IsEmpty() bool {
	return len(s.array) == 0
}

func (s *FlatMap[K, T]) IsNotEmpty() bool {
	return !s.IsEmpty()
}

func (s *FlatMap[K, T]) Clear() {
	s.array = []Entry[K, T]{}
}

func (s *FlatMap[K, T]) Formatted() string {
	//TODO implement me
	panic("implement me")
}

func (s *FlatMap[K, T]) Entries() []Entry[K, T] {
	var entries []Entry[K, T]
	for _, entry := range s.array {
		entries = append(entries, Entry[K, T]{Key: entry.Key, Val: entry.Val})
	}
	return entries
}

func (s *FlatMap[K, T]) Keys() []K {
	var keys []K
	for _, entry := range s.array {
		keys = append(keys, entry.Key)
	}

	return keys
}

func (s *FlatMap[K, T]) Values() []T {
	var values []T
	for _, entry := range s.array {
		values = append(values, entry.Val)
	}

	return values
}
