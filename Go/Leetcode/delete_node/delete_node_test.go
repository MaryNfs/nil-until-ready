package delete_node

import (
	"slices"
	"testing"
)

func TestDeleteNode(t *testing.T) {
	cases := []struct {
		root []any
		key  int
		out  []any
	}{
		{[]any{5, 3, 6, 2, 4, nil, 7}, 3, []any{5, 4, 6, 2, nil, nil, 7}},
		{[]any{5, 3, 6, 2, 4, nil, 7}, 0, []any{5, 3, 6, 2, 4, nil, 7}},
		{[]any{}, 0, []any{}},
	}
	for _, c := range cases {
		root := buildTreeLevelOrder(c.root)
		res := deleteNode(root, c.key)
		got := treeToLevelOrder(res)

		if !slices.Equal(c.out, got) {
			t.Errorf("searchBST(%v,%v)=%v but expected %v", c.root, c.key, got, c.out)
		}

	}
}

func treeToLevelOrder(root *TreeNode) []any {
	if root == nil {
		return []any{}
	}

	var res []any
	q := []*TreeNode{root}

	for len(q) > 0 {
		node := q[0]
		q = q[1:]

		if node == nil {
			res = append(res, nil)
			continue
		}

		res = append(res, node.Val)

		// always push children (even if nil)
		q = append(q, node.Left)
		q = append(q, node.Right)
	}

	// trim trailing nils
	i := len(res) - 1
	for i >= 0 && res[i] == nil {
		i--
	}

	return res[:i+1]
}

func buildTreeLevelOrder(vals []any) *TreeNode {
	if len(vals) == 0 || vals[0] == nil {
		return nil
	}

	root := &TreeNode{Val: vals[0].(int)}
	q := []*TreeNode{root}
	i := 1

	for i < len(vals) && len(q) > 0 {
		cur := q[0]
		q = q[1:]

		// left
		if i < len(vals) && vals[i] != nil {
			cur.Left = &TreeNode{Val: vals[i].(int)}
			q = append(q, cur.Left)
		}
		i++

		// right
		if i < len(vals) && vals[i] != nil {
			cur.Right = &TreeNode{Val: vals[i].(int)}
			q = append(q, cur.Right)
		}
		i++
	}

	return root
}
