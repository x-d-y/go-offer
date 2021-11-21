package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	//root := buildNode()
	fmt.Println(levelOrder(new(TreeNode)))
}

func buildNode() *TreeNode {
	node3 := new(TreeNode)
	node3.Val = 3

	node9 := new(TreeNode)
	node9.Val = 9

	node20 := new(TreeNode)
	node20.Val = 20

	node15 := new(TreeNode)
	node15.Val = 15

	node7 := new(TreeNode)
	node7.Val = 7

	node3.Left, node3.Right = node9, node20
	node20.Left, node20.Right = node15, node7

	return node3
}

var res [][]int

func levelOrder(root *TreeNode) [][]int {
	res = make([][]int, 0)
	level := 0
	r(root, level)
	
	return res
}

func r(root *TreeNode, level int) {
	if len(res) <= level {
		res = append(res, make([]int, 0))
	}
	res[level] = append(res[level], root.Val)
	if root.Left != nil {
		r(root.Left, level+1)
	}
	if root.Right != nil {
		r(root.Right, level+1)
	}
}
