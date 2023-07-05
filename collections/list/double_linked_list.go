package list

import "fmt"

type biDirectionalEntry struct {
	val  int
	next *biDirectionalEntry
	prev *biDirectionalEntry
}

type DoubleLinkedList struct {
	head *biDirectionalEntry
	tail *biDirectionalEntry
}

func MakeDoubleLinkedList() *DoubleLinkedList {
	return &DoubleLinkedList{nil, nil}
}

// Add adds a new entry to the end of the list
func (l *DoubleLinkedList) Add(val int) {
	newEntry := &biDirectionalEntry{val, nil, nil}
	if l.head == nil {
		l.head = newEntry
		l.tail = newEntry
		return
	}

	if l.head.next == nil {
		l.head.next = newEntry
		newEntry.prev = l.head
		l.tail = newEntry
		return
	}

	// tail should never be nil at this point
	l.tail.next = newEntry
	newEntry.prev = l.tail
	l.tail = l.tail.next
}

func (l *DoubleLinkedList) Get(index int) (int, bool) {
	if l.head == nil {
		return 0, false
	}

	current := l.head
	for i := 0; i < index; i++ {
		if current.next == nil {
			return 0, false
		}
		current = current.next
	}

	return current.val, true
}

func (l *DoubleLinkedList) Remove(val int) bool {
	if l.head == nil {
		return false
	}

	if l.head.val == val {
		l.head = l.head.next
		if l.head != nil {
			l.head.prev = nil
		} else {
			l.tail = nil
		}
		return true
	}

	current := l.head.next
	for current != nil {
		if current.val == val {
			current.prev.next = current.next
			if current.next != nil {
				current.next.prev = current.prev
			} else {
				l.tail = current.prev
			}
			return true
		}
		current = current.next
	}

	return false
}

// Contains returns true if the list contains the given value
func (l *DoubleLinkedList) Contains(val int) bool {
	if l.head == nil {
		return false
	}

	current := l.head
	for current != nil {
		if current.val == val {
			return true
		}
		current = current.next
	}

	return false
}

// Size returns the number of entries in the list
func (l *DoubleLinkedList) Size() int {
	if l.head == nil {
		return 0
	}

	current := l.head
	size := 1
	for current.next != nil {
		size++
		current = current.next
	}

	return size
}

// Clear removes all entries from the list
func (l *DoubleLinkedList) Clear() {
	l.head = nil
	l.tail = nil
}

// IsEmpty returns true if the list is empty
func (l *DoubleLinkedList) IsEmpty() bool {
	return l.head == nil
}

// IsNotEmpty returns true if the list is not empty
func (l *DoubleLinkedList) IsNotEmpty() bool {
	return l.head != nil
}

// Formatted returns a string representation of the list
func (l *DoubleLinkedList) Formatted() string {
	if l.head == nil {
		return "[]"
	}

	current := l.head
	formatted := "["
	for current != nil {
		formatted += fmt.Sprintf("%d", current.val)
		if current.next != nil {
			formatted += ", "
		}
		current = current.next
	}
	formatted += "]"

	return formatted
}
