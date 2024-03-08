package dict

import (
	"fmt"
)

type binaryTreeNode[T any] struct {
	Entry[T]
	left  *binaryTreeNode[T]
	right *binaryTreeNode[T]
}

// BinaryTreeMap is a map implementation using a binary tree -- roughly speaking it implements a binary tree solution.
//
// It stores its keys in natural order i.e. ints are stores in ascending order, strings are stored in alphabetical order.
// It is not thread safe and should not be used for concurrent access
// (see concurrent package for thread safe implementations).
//
// It's performance characteristics are:
//
// - Put: O(log n)
//
// - Get: O(log n)
//
// - Remove: O(log n)
type BinaryTreeMap[T any] struct {
	root *binaryTreeNode[T]
}

// MakeBinaryTreeMap creates a new BinaryTreeMap
func MakeBinaryTreeMap[T any]() *BinaryTreeMap[T] {
	return &BinaryTreeMap[T]{}
}

// Put adds a new entry to the map. If the key already exists, it is overwritten.
//
// Time complexity: O(log n)
func (s *BinaryTreeMap[T]) Put(key int, val T) {
	if s.root == nil {
		s.root = &binaryTreeNode[T]{Entry: Entry[T]{Key: key, Val: val}}
		return
	}

	// TODO balance tree on inserts
	node := s.root
	for {
		if key < node.Key {
			if node.left == nil {
				node.left = &binaryTreeNode[T]{Entry: Entry[T]{Key: key, Val: val}}
				return
			}
			node = node.left
		} else if key > node.Key {
			if node.right == nil {
				node.right = &binaryTreeNode[T]{Entry: Entry[T]{Key: key, Val: val}}
				return
			}
			node = node.right
		} else {
			return
		}
	}
}

// Remove removes an entry from map
//
// Time complexity: O(log n)
func (s *BinaryTreeMap[T]) Remove(key int) bool {
	var parent *binaryTreeNode[T]
	node := s.root

	for node != nil {
		if key < node.Key {
			parent = node
			node = node.left
		} else if key > node.Key {
			parent = node
			node = node.right
		} else {
			if parent == nil {
				s.root = removeNode(node)
			} else if parent.left == node {
				parent.left = removeNode(node)
			} else {
				parent.right = removeNode(node)
			}
			return true
		}
	}

	return false
}

func removeNode[T any](node *binaryTreeNode[T]) *binaryTreeNode[T] {
	if node.left == nil {
		return node.right
	} else if node.right == nil {
		return node.left
	}

	minNode := findMinNode(node.right)
	node.Val = minNode.Val
	node.right = removeNode(minNode)

	return node
}

func findMinNode[T any](node *binaryTreeNode[T]) *binaryTreeNode[T] {
	minNode := node
	for minNode.left != nil {
		minNode = minNode.left
	}
	return minNode
}

// Get returns the entry from the map identified by the key along with a boolean value indicating if the key exists.
// If the key does not exist, the second return value is false.
//
// Time complexity: O(log n)
func (s *BinaryTreeMap[T]) Get(key int) (T, bool) {
	node := s.root
	for node != nil {
		if key < node.Key {
			node = node.left
		} else if key > node.Key {
			node = node.right
		} else {
			// how should we return the value
			return node.Val.(T), true
		}
	}

	var zero T
	return zero, false
}

// ContainsKey checks if a given keu exists in the map
func (s *BinaryTreeMap[T]) ContainsKey(key int) bool {
	node := s.root
	for node != nil {
		if key < node.Key {
			node = node.left
		} else if key > node.Key {
			node = node.right
		} else {
			return true
		}
	}
	return false
}

// Size returns the size of the set
func (s *BinaryTreeMap[T]) Size() int {
	return size(s.root)
}

func size[T any](node *binaryTreeNode[T]) int {
	if node == nil {
		return 0
	}
	return 1 + size(node.left) + size(node.right)
}

// IsEmpty checks if the set is empty
func (s *BinaryTreeMap[T]) IsEmpty() bool {
	return s.root == nil
}

// IsNotEmpty checks if the set is not empty
func (s *BinaryTreeMap[T]) IsNotEmpty() bool {
	return s.root != nil
}

// Clear removes all elements from the set
func (s *BinaryTreeMap[T]) Clear() {
	s.root = nil
}

// Formatted returns a string representation of the set
func (s *BinaryTreeMap[T]) Formatted() string {
	return formatted(s.root)
}

func formatted[T any](node *binaryTreeNode[T]) string {
	if node == nil {
		return ""
	}

	return fmt.Sprintf("(%s) - %d - (%s)", formatted(node.left), node.Key, formatted(node.right))
}

func (s *BinaryTreeMap[T]) Entries() []Entry[T] {
	var entries []Entry[T]
	transverse(s.root, func(node *binaryTreeNode[T]) {
		entries = append(entries, Entry[T]{Key: node.Key, Val: node.Val})
	})
	return entries
}

func (s *BinaryTreeMap[T]) Keys() []int {
	var keys []int
	transverse(s.root, func(node *binaryTreeNode[T]) {
		keys = append(keys, node.Key)
	})
	return keys
}

func (s *BinaryTreeMap[T]) Values() []T {
	var values []T
	transverse(s.root, func(node *binaryTreeNode[T]) {
		values = append(values, node.Val.(T))
	})
	return values
}

func transverse[T any](node *binaryTreeNode[T], visit func(node *binaryTreeNode[T])) {
	if node == nil {
		return
	}

	transverse(node.left, visit)
	visit(node)
	transverse(node.right, visit)
}

// ---------------
// OrderedMap methods

// First returns the first entry of the map
func (s *BinaryTreeMap[T]) First() (int, T) {
	var zero T
	if s.root == nil {
		return 0, zero
	}

	node := s.root
	for node.left != nil {
		node = node.left
	}
	return node.Key, node.Val.(T)
}

// Last returns the last entry of the map
func (s *BinaryTreeMap[T]) Last() (int, T) {
	var zero T
	if s.root == nil {
		return 0, zero
	}

	node := s.root
	for node.right != nil {
		node = node.right
	}
	return node.Key, node.Val.(T)
}

// RemoveFirst removes the first element of the set
func (s *BinaryTreeMap[T]) RemoveFirst() bool {
	var parent *binaryTreeNode[T]
	node := s.root

	for node.left != nil {
		parent = node
		node = node.left
	}

	if parent == nil {
		s.root = node.right
	} else {
		parent.left = node.right
	}

	return true
}

// RemoveLast removes the last element of the set
func (s *BinaryTreeMap[T]) RemoveLast() bool {
	var parent *binaryTreeNode[T]
	node := s.root

	for node.right != nil {
		parent = node
		node = node.right
	}

	if parent == nil {
		s.root = node.left
	} else {
		parent.right = node.left
	}

	return true
}

// ----------------
// Specialized methods

// RightRotation performs a right rotation on the tree root node
func (s *BinaryTreeMap[T]) RightRotation() {
	rightRotationOnNode(s.root)
}

// LeftRotation performs a left rotation on the tree root node
func (s *BinaryTreeMap[T]) LeftRotation() {
	leftRotationOnNode(s.root)
}

func rightRotationOnNode[T any](node *binaryTreeNode[T]) {
	if node == nil {
		return
	}

	leftTree := node.left
	node.left = leftTree.right
	leftTree.right = node
}

func leftRotationOnNode[T any](node *binaryTreeNode[T]) {
	if node == nil {
		return
	}

	rightTree := node.right
	node.right = rightTree.left
	rightTree.left = node
}
