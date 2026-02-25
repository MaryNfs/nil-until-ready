package max_depth_binary_tree

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	count := 0
	if root == nil {
		return 0
	}
	count = 1 + max(maxDepth(root.Right), maxDepth(root.Left))

	return count
}
