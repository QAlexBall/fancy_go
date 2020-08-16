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

func InorderTraversal(root *Base.BinaryTreeNode) []int {
	result := make([]int, 0)
	if root == nil {
		return result
	}
	stack := make([]*Base.BinaryTreeNode, 0)
	for len(stack) > 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		val := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result = append(result, val.Val)
		root = val.Right
	}
	return result
}

func PostorderTraversal(root *Base.BinaryTreeNode) []int {
	if root == nil {
		return nil
	}
	result := make([]int, 0)
	stack := make([]*Base.BinaryTreeNode, 0)
	var lastVisit *Base.BinaryTreeNode
	for root != nil || len(stack) != 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		node := stack[len(stack)-1]
		if node.Right == nil || node.Right == lastVisit {
			stack = stack[:len(stack)-1]
			result = append(result, node.Val)
			lastVisit = node
		} else {
			root = node.Right
		}
	}
	return result
}

// DFS

func DfsPreOrderTraversal(root *Base.BinaryTreeNode) []int {
	result := make([]int, 0)
	dfs(root, &result)
	return result
}

func dfs(root *Base.BinaryTreeNode, result *[]int) {
	if root == nil {
		return
	}
	*result = append(*result, root.Val)
	dfs(root.Left, result)
	dfs(root.Right, result)
}

func DfsPreOrderTraversal2(root *Base.BinaryTreeNode) []int {
	result := divideAndConquer(root)
	return result
}

func divideAndConquer(root *Base.BinaryTreeNode) []int {
	result := make([]int, 0)
	if root == nil {
		return result
	}
	left := divideAndConquer(root.Left)
	right := divideAndConquer(root.Right)

	result = append(result, root.Val)
	result = append(result, left...)
	result = append(result, right...)
	return result
}

// BFS
func LevelOrder(root *Base.BinaryTreeNode) [][]int {
	result := make([][]int, 0)
	if root == nil {
		return result
	}
	queue := make([]*Base.BinaryTreeNode, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		list := make([]int, 0)
		l := len(queue)
		for i := 0; i < l; i++ {
			level := queue[0]
			queue = queue[1:]
			list = append(list, level.Val)
			if level.Left != nil {
				queue = append(queue, level.Left)
			}
			if level.Right != nil {
				queue = append(queue, level.Right)
			}
		}
		result = append(result, list)
	}
	return result
}
