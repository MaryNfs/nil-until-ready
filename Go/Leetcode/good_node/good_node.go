package good_node

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func goodNodes(root *TreeNode) int {
	return dfs(root, root.Val)
}

func dfs(root *TreeNode, maxVal int) int {
	var res int
	if root == nil {
		return 0
	}
	if root.Val >= maxVal {
		res = 1
	} else {
		res = 0
	}
	maxVal = max(root.Val, maxVal)
	res += dfs(root.Right, maxVal)
	res += dfs(root.Left, maxVal)

	return res
}
