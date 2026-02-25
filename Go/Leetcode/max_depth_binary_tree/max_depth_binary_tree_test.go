package max_depth_binary_tree

import "testing"

func TestMaxDepth(t *testing.T) {

	cases := []struct {
		tree   []any
		output int
	}{
		{nil, 0},
		{[]any{1}, 1},
		{[]any{3, 9, 20, nil, nil, 15, 7}, 3},
	}
	for _, c := range cases {
		root := buildTreeLevelOrder(c.tree)
		res := maxDepth(root)
		if res != c.output {
			t.Errorf("maxDepth(%v)=%v, expected: %v", c.tree, res, c.output)
		}
	}
}

// AI
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
