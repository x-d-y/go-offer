package main

import "fmt"

func main() {
	root := buildNode()
	levelOrder(root)
}

var res [][]int

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	res = make([][]int, 0)
	cQueue := make([]*TreeNode, 0)
	cQueue = append(cQueue, root)
	r(0, cQueue)
	fmt.Println(res)
	return res
}

func r(level int, cQueue []*TreeNode) {
	if len(cQueue) == 0 {
		return
	}
	resSub := make([]int, 0)
	nQueue := make([]*TreeNode, 0)
	for _, v := range cQueue {
		if level%2 == 0 {
			resSub = append(resSub, v.Val)
		} else {
			resSub = pushOnleft(resSub, v.Val)
		}
		if v.Left != nil {
			nQueue = append(nQueue, v.Left)
		}
		if v.Right != nil {
			nQueue = append(nQueue, v.Right)
		}
	}
	res = append(res, resSub)
	r(level+1, nQueue)
}

func pushOnleft(a []int, b int) []int {
	c := make([]int, 0)
	c = append(c, b)
	c = append(c, a...)
	return c
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
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
