package main

import "fmt"

func main() {
	root := buildTree([]int{9,6,8}, []int{6,9,8})
	fmt.Println(root)
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	root := new(TreeNode)
	middle := preorder[0]
	if len(preorder) == 1 {
		root.Val = middle
		return root
	}
	for i, v := range inorder {
		if v == middle {
			root.Val = v
			root.Left = buildTree(preorder[1:i+1], inorder[0:i])
			root.Right = buildTree(preorder[i+1:], inorder[i+1:])
			break
		}
	}
	return root
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
