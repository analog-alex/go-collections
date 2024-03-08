package dict

type FlatMap[T any] struct {
	array []Entry[T]
}

func MakeFlatMap[T any]() *FlatMap[T] {
	return &FlatMap[T]{}
}

/*
* Binary search for the key in the array.
* We keep the array sorted by key, so we can use binary search to find the key.
 */
func (s *FlatMap[T]) binarySearch(key int) (Entry[T], bool) {
	l, h := 0, len(s.array)-1
	for l <= h {
		m := (l + h) / 2
		if s.array[m].Key == key {
			return s.array[m], true
		}

		if s.array[m].Key < key {
			l = m + 1
		} else {
			h = m - 1
		}
	}

	return Entry[T]{}, false
}

// Put inserts a new key-value pair into the map. If the key already exists, the value is updated.
func (s *FlatMap[T]) Put(key int, val T) {
	s.array = append(s.array, Entry[T]{Key: key, Val: val})

	for i := len(s.array) - 1; i > 0; i-- {
		// conserve the well ordering of the array elements
		if s.array[i].Key < s.array[i-1].Key {
			s.array[i], s.array[i-1] = s.array[i-1], s.array[i]
		} else {
			break
		}
	}
}

func (s *FlatMap[T]) Remove(key int) bool {
	for i, entry := range s.array {
		if entry.Key == key {
			s.array = append(s.array[:i], s.array[i+1:]...)
			return true
		}
	}

	return false
}

func (s *FlatMap[T]) Get(key int) (T, bool) {
	l, ok := s.binarySearch(key)
	if !ok {
		var zero T
		return zero, false
	}
	return l.Val.(T), true
}

func (s *FlatMap[T]) ContainsKey(key int) bool {
	_, ok := s.binarySearch(key)
	return ok
}

func (s *FlatMap[T]) Size() int {
	return len(s.array)
}

func (s *FlatMap[T]) IsEmpty() bool {
	return len(s.array) == 0
}

func (s *FlatMap[T]) IsNotEmpty() bool {
	return !s.IsEmpty()
}

func (s *FlatMap[T]) Clear() {
	s.array = []Entry[T]{}
}

func (s *FlatMap[T]) Formatted() string {
	//TODO implement me
	panic("implement me")
}

func (s *FlatMap[T]) Entries() []Entry[T] {
	var entries []Entry[T]
	for _, entry := range s.array {
		entries = append(entries, Entry[T]{Key: entry.Key, Val: entry.Val})
	}
	return entries
}

func (s *FlatMap[T]) Keys() []int {
	var keys []int
	for _, entry := range s.array {
		keys = append(keys, entry.Key)
	}

	return keys
}

func (s *FlatMap[T]) Values() []T {
	var values []T
	for _, entry := range s.array {
		values = append(values, entry.Val.(T))
	}

	return values
}
