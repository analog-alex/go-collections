package dict

import (
	"fmt"
)

type binaryTreeNode struct {
	key   int
	val   string
	left  *binaryTreeNode
	right *binaryTreeNode
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
type BinaryTreeMap struct {
	root *binaryTreeNode
}

// MakeBinaryTreeMap creates a new BinaryTreeMap
func MakeBinaryTreeMap() *BinaryTreeMap {
	return &BinaryTreeMap{}
}

// Put adds a new entry to the map. If the key already exists, it is overwritten.
//
// Time complexity: O(log n)
func (s *BinaryTreeMap) Put(key int, val string) {
	if s.root == nil {
		s.root = &binaryTreeNode{key: key, val: val}
		return
	}

	// TODO balance tree on inserts
	node := s.root
	for {
		if key < node.key {
			if node.left == nil {
				node.left = &binaryTreeNode{key: key, val: val}
				return
			}
			node = node.left
		} else if key > node.key {
			if node.right == nil {
				node.right = &binaryTreeNode{key: key, val: val}
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
func (s *BinaryTreeMap) Remove(key int) bool {
	var parent *binaryTreeNode
	node := s.root

	for node != nil {
		if key < node.key {
			parent = node
			node = node.left
		} else if key > node.key {
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

func removeNode(node *binaryTreeNode) *binaryTreeNode {
	if node.left == nil {
		return node.right
	} else if node.right == nil {
		return node.left
	}

	minNode := findMinNode(node.right)
	node.val = minNode.val
	node.right = removeNode(minNode)

	return node
}

func findMinNode(node *binaryTreeNode) *binaryTreeNode {
	min := node
	for min.left != nil {
		min = min.left
	}
	return min
}

// Get returns the entry from the map identified by the key along with a boolean value indicating if the key exists.
// If the key does not exist, the second return value is false.
//
// Time complexity: O(log n)
func (s *BinaryTreeMap) Get(key int) (string, bool) {
	node := s.root
	for node != nil {
		if key < node.key {
			node = node.left
		} else if key > node.key {
			node = node.right
		} else {
			return node.val, true
		}
	}
	return "", false
}

// ContainsKey checks if a given keu exists in the map
func (s *BinaryTreeMap) ContainsKey(key int) bool {
	node := s.root
	for node != nil {
		if key < node.key {
			node = node.left
		} else if key > node.key {
			node = node.right
		} else {
			return true
		}
	}
	return false
}

// Size returns the size of the set
func (s *BinaryTreeMap) Size() int {
	return size(s.root)
}

func size(node *binaryTreeNode) int {
	if node == nil {
		return 0
	}
	return 1 + size(node.left) + size(node.right)
}

// IsEmpty checks if the set is empty
func (s *BinaryTreeMap) IsEmpty() bool {
	return s.root == nil
}

// IsNotEmpty checks if the set is not empty
func (s *BinaryTreeMap) IsNotEmpty() bool {
	return s.root != nil
}

// Clear removes all elements from the set
func (s *BinaryTreeMap) Clear() {
	s.root = nil
}

// Formatted returns a string representation of the set
func (s *BinaryTreeMap) Formatted() string {
	return formatted(s.root)
}

func formatted(node *binaryTreeNode) string {
	if node == nil {
		return ""
	}

	return fmt.Sprintf("(%s) - %d - (%s)", formatted(node.left), node.key, formatted(node.right))
}

func (s *BinaryTreeMap) Entries() []Entry {
	var entries []Entry
	transverse(s.root, func(node *binaryTreeNode) {
		entries = append(entries, Entry{Key: node.key, Val: node.val})
	})
	return entries
}

func (s *BinaryTreeMap) Keys() []int {
	var keys []int
	transverse(s.root, func(node *binaryTreeNode) {
		keys = append(keys, node.key)
	})
	return keys
}

func (s *BinaryTreeMap) Values() []string {
	var values []string
	transverse(s.root, func(node *binaryTreeNode) {
		values = append(values, node.val)
	})
	return values
}

func transverse(node *binaryTreeNode, visit func(node *binaryTreeNode)) {
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
func (s *BinaryTreeMap) First() (int, string) {
	if s.root == nil {
		return 0, ""
	}

	node := s.root
	for node.left != nil {
		node = node.left
	}
	return node.key, node.val
}

// Last returns the last entry of the map
func (s *BinaryTreeMap) Last() (int, string) {
	if s.root == nil {
		return 0, ""
	}

	node := s.root
	for node.right != nil {
		node = node.right
	}
	return node.key, node.val
}

// RemoveFirst removes the first element of the set
func (s *BinaryTreeMap) RemoveFirst() bool {
	var parent *binaryTreeNode
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
func (s *BinaryTreeMap) RemoveLast() bool {
	var parent *binaryTreeNode
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
