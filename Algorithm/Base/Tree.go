package Base

type BinaryTreeNode struct {
	Left  *BinaryTreeNode
	Val   int
	Right *BinaryTreeNode
}

func (node *BinaryTreeNode) AddLeftNode(newNode *BinaryTreeNode) *BinaryTreeNode {
	node.Left = &BinaryTreeNode{newNode.Left, newNode.Val, newNode.Right}
	return node.Left
}

func (node *BinaryTreeNode) AddRightNode(newNode *BinaryTreeNode) *BinaryTreeNode {
	node.Right = &BinaryTreeNode{newNode.Left, newNode.Val, newNode.Right}
	return node.Right
}
