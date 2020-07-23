package main

import (
	"Algorithm/Base"
	"Algorithm/BinaryTree"
	"fmt"
)

func main() {
	left := Base.BinaryTreeNode{Left: nil, Val: 0, Right: nil}
	right := Base.BinaryTreeNode{Left: nil, Val: 2, Right: nil}
	root := Base.BinaryTreeNode{Left: &left, Val: 1, Right: &right}
	left_left := left.AddLeftNode(&Base.BinaryTreeNode{Left: nil, Val: 10, Right: nil})
	left.AddRightNode(&Base.BinaryTreeNode{Left: nil, Val: 6, Right: nil})
	left_left.AddRightNode(&Base.BinaryTreeNode{Left: nil, Val: 5, Right: nil})
	BinaryTree.PreorderTraversal(&root)
	valList := BinaryTree.PreorderTraversal2(&root)
	fmt.Println(valList)
}
