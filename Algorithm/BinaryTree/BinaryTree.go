package BinaryTree

import (
	"Algorithm/Base"
	"fmt"
)

func PreorderTraversal(root *Base.BinaryTreeNode) {
	if root == nil {
		return
	}
	fmt.Println(root.Val)
	PreorderTraversal(root.Left)
	PreorderTraversal(root.Right)
}

func PreorderTraversal2(root *Base.BinaryTreeNode) []int {
	if root == nil {
		return nil
	}

	result := make([]int, 0)
	stack := make([]*Base.BinaryTreeNode, 0)

	for root != nil || len(stack) != 0 {
		for root != nil {
			result = append(result, root.Val)
			stack = append(stack, root)
			root = root.Left
		}
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		root = node.Right
	}
	return result
}
