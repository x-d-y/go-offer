package main

func main() {

}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
if root == nil {
	return nil
}

if root == p || root == q {
	return root
}

left := lowestCommonAncestor(root.Left, p, q)
right := lowestCommonAncestor(root.Right, p, q)

if left == nil {
	return right
}

if right == nil {
	return left
}

if left != nil && right != nil {
	return root
}
return nil
}
