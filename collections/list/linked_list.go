package list

import (
	"fmt"
	"reflect"
)

type entry[T any] struct {
	val  T
	next *entry[T]
}

// LinkedList is a single linked list
type LinkedList[T any] struct {
	head *entry[T]
	tail *entry[T] // we keep a pointer to the tail just to make adding elements faster
}

// MakeLinkedList returns a pointer to a new LinkedList
func MakeLinkedList[T any]() *LinkedList[T] {
	return &LinkedList[T]{nil, nil}
}

// Add adds a new entry to the end of the list
func (l *LinkedList[T]) Add(val T) {
	newEntry := &entry[T]{val, nil}
	if l.head == nil {
		l.head = newEntry
		l.tail = newEntry
		return
	}

	if l.head.next == nil {
		l.head.next = newEntry
		l.tail = newEntry
		return
	}

	// tail should never be nil at this point
	l.tail.next = newEntry
	l.tail = l.tail.next
}

// Get returns the value at the given index and true if the index is valid, otherwise 0 and false
func (l *LinkedList[T]) Get(index int) (T, bool) {
	if l.head == nil {
		var zero T
		return zero, false
	}

	current := l.head
	for i := 0; i < index; i++ {
		if current.next == nil {
			var zero T
			return zero, false
		}

		current = current.next
	}

	return current.val, true
}

// Remove removes the first entry with the given value, return true if an entry was removed or false is not found
func (l *LinkedList[T]) Remove(val T) bool {
	if l.head == nil {
		return false
	}

	if reflect.DeepEqual(l.head.val, val) {
		l.head = l.head.next
		if l.head == nil { // is this if necessary? me thinks not, but good to be safe
			l.tail = nil
		}
		return true
	}

	prev := l.head
	current := l.head.next
	for current != nil {
		if reflect.DeepEqual(current.val, val) {
			prev.next = current.next
			if current == l.tail {
				l.tail = prev
			}
			return true
		}
		prev = current
		current = current.next
	}

	return false
}

// Contains returns true if the list contains the given value
func (l *LinkedList[T]) Contains(val T) bool {
	if l.head == nil {
		return false
	}

	current := l.head
	for current != nil {
		if reflect.DeepEqual(current.val, val) {
			return true
		}
		current = current.next
	}

	return false
}

// Size returns the number of entries in the list
func (l *LinkedList[T]) Size() int {
	if l.head == nil {
		return 0
	}

	count := 0
	current := l.head
	for current != nil {
		count++
		current = current.next
	}
	return count
}

// Clear removes all entries from the list
func (l *LinkedList[T]) Clear() {
	l.head = nil
	l.tail = nil
}

// IsEmpty returns true if the list is empty
func (l *LinkedList[T]) IsEmpty() bool {
	return l.head == nil
}

// IsNotEmpty returns true if the list is not empty
func (l *LinkedList[T]) IsNotEmpty() bool {
	return l.head != nil
}

func (l *LinkedList[T]) Formatted() string {
	if l.head == nil {
		return "[]"
	}

	current := l.head
	s := "["
	for current != nil {
		s += fmt.Sprintf("%v", current.val)
		if current.next != nil {
			s += ", "
		}
		current = current.next
	}
	s += "]"
	return s
}
