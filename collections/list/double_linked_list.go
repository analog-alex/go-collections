package list

import (
	"fmt"
	"reflect"
)

type biDirectionalEntry[T any] struct {
	val  T
	next *biDirectionalEntry[T]
	prev *biDirectionalEntry[T]
}

type DoubleLinkedList[T any] struct {
	head *biDirectionalEntry[T]
	tail *biDirectionalEntry[T]
}

func MakeDoubleLinkedList[T any]() *DoubleLinkedList[T] {
	return &DoubleLinkedList[T]{nil, nil}
}

// Add adds a new entry to the end of the list
func (l *DoubleLinkedList[T]) Add(val T) {
	newEntry := &biDirectionalEntry[T]{val, nil, nil}
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

func (l *DoubleLinkedList[T]) Get(index int) (T, bool) {
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

func (l *DoubleLinkedList[T]) Remove(val T) bool {
	if l.head == nil {
		return false
	}

	if reflect.DeepEqual(l.head.val, val) {
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
		if reflect.DeepEqual(current.val, val) {
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
func (l *DoubleLinkedList[T]) Contains(val T) bool {
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
func (l *DoubleLinkedList[T]) Size() int {
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
func (l *DoubleLinkedList[T]) Clear() {
	l.head = nil
	l.tail = nil
}

// IsEmpty returns true if the list is empty
func (l *DoubleLinkedList[T]) IsEmpty() bool {
	return l.head == nil
}

// IsNotEmpty returns true if the list is not empty
func (l *DoubleLinkedList[T]) IsNotEmpty() bool {
	return l.head != nil
}

// Formatted returns a string representation of the list
func (l *DoubleLinkedList[T]) Formatted() string {
	if l.head == nil {
		return "[]"
	}

	current := l.head
	formatted := "["
	for current != nil {
		formatted += fmt.Sprintf("%v", current.val)
		if current.next != nil {
			formatted += ", "
		}
		current = current.next
	}
	formatted += "]"

	return formatted
}
