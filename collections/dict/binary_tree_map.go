package dict

import (
	"fmt"
)

type binaryTreeNode[K any, T any] struct {
	Entry[K, T]
	left  *binaryTreeNode[K, T]
	right *binaryTreeNode[K, T]
}

// BinaryTreeMap is a map implementation using a binary tree -- roughly speaking it implements a binary tree solution.
//
// It stores its keys in natural order i.e. int's are stored in ascending order, strings are stored in alphabetical order.
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
type BinaryTreeMap[K any, T any] struct {
	root       *binaryTreeNode[K, T]
	comparator func(a, b K) int
}

// MakeBinaryTreeMap creates a new BinaryTreeMap
func MakeBinaryTreeMap[K any, T any](c func(a, b K) int) *BinaryTreeMap[K, T] {
	return &BinaryTreeMap[K, T]{comparator: c}
}

// Put adds a new entry to the map. If the key already exists, it is overwritten.
//
// Time complexity: O(log n)
func (s *BinaryTreeMap[K, T]) Put(key K, val T) {
	if s.root == nil {
		s.root = &binaryTreeNode[K, T]{Entry: Entry[K, T]{Key: key, Val: val}}
		return
	}

	// TODO balance tree on inserts
	node := s.root
	for {
		c := s.comparator(key, node.Key)
		if c < 0 {
			if node.left == nil {
				node.left = &binaryTreeNode[K, T]{Entry: Entry[K, T]{Key: key, Val: val}}
				return
			}
			node = node.left
		} else if c > 0 {
			if node.right == nil {
				node.right = &binaryTreeNode[K, T]{Entry: Entry[K, T]{Key: key, Val: val}}
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
func (s *BinaryTreeMap[K, T]) Remove(key K) bool {
	var parent *binaryTreeNode[K, T]
	node := s.root

	for node != nil {
		c := s.comparator(key, node.Key)
		if c < 0 {
			parent = node
			node = node.left
		} else if c > 0 {
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

func removeNode[K any, T any](node *binaryTreeNode[K, T]) *binaryTreeNode[K, T] {
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

func findMinNode[K any, T any](node *binaryTreeNode[K, T]) *binaryTreeNode[K, T] {
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
func (s *BinaryTreeMap[K, T]) Get(key K) (T, bool) {
	node := s.root
	for node != nil {
		c := s.comparator(key, node.Key)
		if c < 0 {
			node = node.left
		} else if c > 0 {
			node = node.right
		} else {
			// how should we return the value
			return node.Val, true
		}
	}

	var zero T
	return zero, false
}

// ContainsKey checks if a given keu exists in the map
func (s *BinaryTreeMap[K, T]) ContainsKey(key K) bool {
	node := s.root
	for node != nil {
		c := s.comparator(key, node.Key)
		if c < 0 {
			node = node.left
		} else if c > 0 {
			node = node.right
		} else {
			return true
		}
	}
	return false
}

// Size returns the size of the set
func (s *BinaryTreeMap[K, T]) Size() int {
	return size(s.root)
}

func size[K any, T any](node *binaryTreeNode[K, T]) int {
	if node == nil {
		return 0
	}
	return 1 + size(node.left) + size(node.right)
}

// IsEmpty checks if the set is empty
func (s *BinaryTreeMap[K, T]) IsEmpty() bool {
	return s.root == nil
}

// IsNotEmpty checks if the set is not empty
func (s *BinaryTreeMap[K, T]) IsNotEmpty() bool {
	return s.root != nil
}

// Clear removes all elements from the set
func (s *BinaryTreeMap[K, T]) Clear() {
	s.root = nil
}

// Formatted returns a string representation of the set
func (s *BinaryTreeMap[K, T]) Formatted() string {
	return formatted(s.root)
}

func formatted[K any, T any](node *binaryTreeNode[K, T]) string {
	if node == nil {
		return ""
	}

	return fmt.Sprintf("(%s) - %T - (%s)", formatted(node.left), node.Key, formatted(node.right))
}

func (s *BinaryTreeMap[K, T]) Entries() []Entry[K, T] {
	var entries []Entry[K, T]
	transverse(s.root, func(node *binaryTreeNode[K, T]) {
		entries = append(entries, Entry[K, T]{Key: node.Key, Val: node.Val})
	})
	return entries
}

func (s *BinaryTreeMap[K, T]) Keys() []K {
	var keys []K
	transverse(s.root, func(node *binaryTreeNode[K, T]) {
		keys = append(keys, node.Key)
	})
	return keys
}

func (s *BinaryTreeMap[K, T]) Values() []T {
	var values []T
	transverse(s.root, func(node *binaryTreeNode[K, T]) {
		values = append(values, node.Val)
	})
	return values
}

func transverse[K any, T any](node *binaryTreeNode[K, T], visit func(node *binaryTreeNode[K, T])) {
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
func (s *BinaryTreeMap[K, T]) First() (K, T) {
	var zkey K
	var zero T
	if s.root == nil {
		return zkey, zero
	}

	node := s.root
	for node.left != nil {
		node = node.left
	}
	return node.Key, node.Val
}

// Last returns the last entry of the map
func (s *BinaryTreeMap[K, T]) Last() (K, T) {
	var zkey K
	var zero T
	if s.root == nil {
		return zkey, zero
	}

	node := s.root
	for node.right != nil {
		node = node.right
	}
	return node.Key, node.Val
}

// RemoveFirst removes the first element of the set
func (s *BinaryTreeMap[K, T]) RemoveFirst() bool {
	var parent *binaryTreeNode[K, T]
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
func (s *BinaryTreeMap[K, T]) RemoveLast() bool {
	var parent *binaryTreeNode[K, T]
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
func (s *BinaryTreeMap[K, T]) RightRotation() {
	rightRotationOnNode(s.root)
}

// LeftRotation performs a left rotation on the tree root node
func (s *BinaryTreeMap[K, T]) LeftRotation() {
	leftRotationOnNode(s.root)
}

func rightRotationOnNode[K any, T any](node *binaryTreeNode[K, T]) {
	if node == nil {
		return
	}

	leftTree := node.left
	node.left = leftTree.right
	leftTree.right = node
}

func leftRotationOnNode[K any, T any](node *binaryTreeNode[K, T]) {
	if node == nil {
		return
	}

	rightTree := node.right
	node.right = rightTree.left
	rightTree.left = node
}
