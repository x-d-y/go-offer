package main

import (
	"fmt"
	"strconv"
)

func main() {

	node := buildTree([]string{"1","0","-4","-3"})
	node2 := buildTree([]string{"1", "4"})
	fmt.Println(isSubStructure(node, node2))
}

func isSubStructure(A *TreeNode, B *TreeNode) bool {
	if B == nil {
		return false
	}
	return r(A, B, B)
}

func r(A *TreeNode, B *TreeNode, C *TreeNode) bool {
	if B == nil {
		return true
	}
	if A == nil {
		return false
	}
	if A.Val == B.Val {
		if r(A.Left, B.Left, C) && r(A.Right, B.Right, C) {
			return true
		}
	}
	if r(A.Left, C, C) {
		return true
	} else {
		return r(A.Right, C, C)
	}
}

func buildTree(nums []string) *TreeNode {
	nodes := make([]*TreeNode, len(nums))
	for i, v := range nums {
		var cNode *TreeNode
		if v != "null" {
			n, _ := strconv.Atoi(v)
			cNode = new(TreeNode)
			cNode.Val = n
		} else {
			nodes[i] = nil
		}
		nodes[i] = cNode
		if i > 2 {
			yu := (i - 1) % 2
			sh := (i - 1) / 2
			if yu == 0 {
				nodes[sh].Left = cNode
			} else {
				nodes[sh].Right = cNode
			}

		}
	}

	if len(nums) >= 2 && nodes[1] != nil {
		nodes[0].Left = nodes[1]
	}

	if len(nums) >= 3 && nodes[2] != nil {
		nodes[0].Right = nodes[2]
	}
	if len(nodes) >= 1 {
		return nodes[0]
	}
	return nil

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
