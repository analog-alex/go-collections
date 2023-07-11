package dict

type redBlackTreeNode struct {
	key    int
	val    string
	red    bool
	left   *redBlackTreeNode
	right  *redBlackTreeNode
	parent *redBlackTreeNode
}

type RedBlackTreeMap struct {
	root *redBlackTreeNode
}

func MakeRedBlackTreeMap() *RedBlackTreeMap {
	return &RedBlackTreeMap{}
}

func (s *RedBlackTreeMap) Put(key int, val string) {
	if s.root == nil {
		s.root = &redBlackTreeNode{key: key, val: val, red: false}
		return
	}

	node := s.root
	prev := s.root
	for {
		if key < node.key {
			if node.left == nil {
				node.left = &redBlackTreeNode{key: key, val: val, red: true, parent: prev}
				break
			}
			prev = node
			node = node.left
		} else if key > node.key {
			if node.right == nil {
				node.right = &redBlackTreeNode{key: key, val: val, red: true, parent: prev}
				break
			}
			prev = node
			node = node.right
		} else {
			break
		}
	}
	balanceFromNode(node)
}

func balanceFromNode(node *redBlackTreeNode) {
	if node.parent == nil {
		node.red = false
		return
	}
}

func rotateRightOnRedBlackNode(node *redBlackTreeNode) {
	// TODO
}

func rotateLeftOnRedBlackNode(node *redBlackTreeNode) {
	// TODO
}
