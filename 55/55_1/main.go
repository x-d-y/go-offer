package main

func main() {

}

var max int

func maxDepth(root *TreeNode) int {
	max = 0
	r(root,1)
	return max-1
}

func r(root *TreeNode, level int) {
	if level > max{
		max = level
	}
	if root == nil {
		return
	}
	r(root.Left, level++)
	r(root.Right, level++)
}
