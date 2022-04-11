package treap

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// Treap is a data-structure which is a combine of tree and heap.
type Treap struct {
	Val         int
	Pri         int
	Left, Right *Treap
}

func (t *Treap) String() string {
	if t == nil {
		return " "
	}
	return fmt.Sprintf("t.val: %d, t.left: %v, t.right: %v", t.Val, t.Left, t.Right)
}

func NewTreap(val int) *Treap {
	return newTreap(val)
}

func newTreap(val int) *Treap {
	return &Treap{
		Val: val,
		Pri: rand.Intn(1000),
	}
}

// Find like bst find
func Find(root *Treap, val int) *Treap {
	if root == nil {
		return nil
	}
	if root.Val == val {
		return root
	}
	if val > root.Val {
		return Find(root.Right, val)
	}
	return Find(root.Left, val)
}

func leftRotation(root *Treap) *Treap {
	if root == nil {
		return nil
	}
	if root.Right == nil {
		return root
	}
	right := root.Right
	rl := root.Right.Left
	right.Left = root
	root.Right = rl
	return right
}

func rightRotation(root *Treap) *Treap {
	if root == nil {
		return nil
	}
	if root.Left == nil {
		return root
	}
	left := root.Left
	lr := root.Left.Right
	left.Right = root
	root.Left = lr
	return left
}

func Insert(root *Treap, val int) *Treap {
	if root == nil {
		return newTreap(val)
	}
	if val < root.Val {
		root.Left = Insert(root.Left, val)
	} else {
		root.Right = Insert(root.Right, val)
	}
	if root.Left != nil && root.Left.Pri > root.Pri {
		root = rightRotation(root)
	} else if root.Right != nil && root.Right.Pri > root.Pri {
		root = leftRotation(root)
	}
	return root
}
