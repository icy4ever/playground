package treap

import (
	"testing"
)

func TestInsert(t *testing.T) {
	cases := []struct {
		Vals []int
	}{
		{[]int{2, 3, 1, 4, 5}},
	}
	for _, c := range cases {
		root := NewTreap(c.Vals[0])
		for i := 1; i < len(c.Vals); i++ {
			root = Insert(root, c.Vals[i])
		}
		if !checker(root) {
			t.Error("test failed")
		}
	}
}

func checker(root *Treap) bool {
	if root == nil {
		return true
	}
	if (root.Left == nil || root.Val > root.Left.Val && root.Pri > root.Left.Pri) &&
		(root.Right == nil || root.Val < root.Right.Val && root.Pri > root.Right.Pri) {
		return checker(root.Left) && checker(root.Right)
	}
	return false
}
