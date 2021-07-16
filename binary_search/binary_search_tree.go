package binarysearch

// Node represents the components of a binary search tree
type Node struct {
	key   int
	Left  *Node
	Right *Node
}

// Insert add a node to the tree
// the key to add should not be already in the tree
func (n *Node) Insert(k int) {
	if n.key < k {
		// move right
		if n.Right == nil {
			n.Right = &Node{key: k}
		} else {
			n.Right.Insert(k)
		}
	} else if n.key > k {
		// move left
		if n.Left == nil {
			n.Left = &Node{key: k}
		} else {
			n.Left.Insert(k)
		}
	}
}

// Search will take in a key value and
// RETURN true if there is a node with that value
func (n *Node) Search(k int) bool {
	if n == nil {
		return false
	}

	if n.key < k {
		// move right
		return n.Right.Search(k)
	} else if n.key > k {
		// move left
		return n.Left.Search(k)
	}
	return true
}
