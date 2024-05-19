package list

import (
	"fmt"
	"reflect"
)

const (
	defaultCapacity = 10
	scalingFactor   = 2
)

type DynamicArray[T any] struct {
	data []T
}

func MakeDynamicArray[T any]() *DynamicArray[T] {
	return &DynamicArray[T]{make([]T, 0, defaultCapacity)} // use defaults
}

func MakeDynamicArrayWithCapacity[T any](capacity int) *DynamicArray[T] {
	return &DynamicArray[T]{make([]T, 0, capacity)}
}

func MakeDynamicArrayWithValues[T any](values ...T) *DynamicArray[T] {
	return &DynamicArray[T]{values}
}

func (d *DynamicArray[T]) Add(val T) {
	if len(d.data) == cap(d.data) {
		newData := make([]T, len(d.data), cap(d.data)*scalingFactor)
		copy(newData, d.data)
		d.data = newData
	}

	d.data = append(d.data, val)
}

func (d *DynamicArray[T]) Remove(val T) bool {
	for i, v := range d.data {
		if reflect.DeepEqual(v, val) {
			d.data = append(d.data[:i], d.data[i+1:]...)
			return true
		}
	}
	return false
}

func (d *DynamicArray[T]) Get(index int) (T, bool) {
	if index < 0 || index >= len(d.data) {
		var zero T
		return zero, false
	}
	return d.data[index], true
}

func (d *DynamicArray[T]) Contains(val T) bool {
	for _, v := range d.data {
		if reflect.DeepEqual(v, val) {
			return true
		}
	}
	return false
}

func (d *DynamicArray[T]) Size() int {
	return len(d.data)
}

func (d *DynamicArray[T]) IsEmpty() bool {
	return len(d.data) == 0
}

func (d *DynamicArray[T]) IsNotEmpty() bool {
	return len(d.data) > 0
}

func (d *DynamicArray[T]) Clear() {
	d.data = make([]T, 0)
}

func (d *DynamicArray[T]) Formatted() string {
	var s = "["
	for i, v := range d.data {
		if i == len(d.data)-1 {
			s += fmt.Sprintf("%T", v)
		} else {
			s += fmt.Sprintf("%T", v) + ", "
		}
	}
	s += "]"
	return s
}
