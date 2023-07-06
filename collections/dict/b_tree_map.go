package dict

type bTreeNode struct {
	key      int
	val      string
	children []*bTreeNode
}

type BTreeMap struct {
	root     *bTreeNode
	leafSize int
}

func MakeBTreeMap(leafSize int) *BTreeMap {
	return &BTreeMap{leafSize: leafSize}
}

func (s *BTreeMap) Put(key int, val string) {

}
