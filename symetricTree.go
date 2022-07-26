package main

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	var isMirror func(left, right *TreeNode) bool
	isMirror = func(left, right *TreeNode) bool {
		if left == nil && right == nil {
			return true
		}
		if left == nil || right == nil {
			return false
		}
		return left.Val == right.Val && isMirror(left.Left, right.Right) && isMirror(left.Right, right.Left)
	}

	return isMirror(root.Left, root.Right)
}
