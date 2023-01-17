package binarytree

type BinaryTreeNode struct {
	data  int
	left  *BinaryTreeNode
	right *BinaryTreeNode

	// below data members used only for some of the problems
	next   *BinaryTreeNode
	parent *BinaryTreeNode
	count  int
}

// NewEduBinaryTreeNode initializes and returns the BinaryTreeNode
// type object
func NewBinaryTreeNode(data int) *BinaryTreeNode {
	t := new(BinaryTreeNode)
	t.data = data
	t.left = nil
	t.right = nil
	t.next = nil
	t.parent = nil
	t.count = 0
	return t
}
