package leaf_similar

import "testing"

func TestLeafSimilar(t *testing.T) {
	cases := []struct {
		tree1  []any
		tree2  []any
		output bool
	}{
		{[]any{3, 5, 1, 6, 2, 9, 8, nil, nil, 7, 4}, []any{3, 5, 1, 6, 7, 4, 2, nil, nil, nil, nil, nil, nil, 9, 8}, true},
		{[]any{1, 2, 3}, []any{1, 3, 2}, false},
	}
	for _, c := range cases {
		root1 := buildTreeLevelOrder(c.tree1)
		root2 := buildTreeLevelOrder(c.tree2)
		res := leafSimilar(root1, root2)
		if res != c.output {
			t.Errorf("maxDepth(%v,%v)=%v, expected: %v", c.tree1, c.tree2, res, c.output)
		}
	}
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
