package BPlus_tree

type Key []byte
type Value []byte

type PagePointer struct {
	pageNo int
}

type BPlusTree struct {
	root *BPlusTreeNode
}

type BPlusTreeNode struct {
	keys     []Key
	children []PagePointer
}

var (
	root *BPlusTreeNode
)

func Less(l Key, r Key) bool {
	return true
}

func (n *BPlusTreeNode) find(key Key) PagePointer {
	be, en := 0, len(n.keys)
	res := -1
	for be <= en {
		mid := (be + en) / 2
		if Less(n.keys[be], key) {
			res = mid
			en = mid - 1
		} else {
			be = mid + 1
		}
	}
	return n.children[res+1]
}

func (bt *BPlusTree) Get(key Key) Value {
	if root == nil {
		return nil
	}
	cur := bt.root
	for cur != nil {
		nxt := cur.find(key)
	}
	n := cur.find(key)
}
