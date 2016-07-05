package bst

import "testing"

func TestGetMaxHeight(t *testing.T) {

	nums := []int{3, 5, 2, 1, 4, 6, 7}
	tree := NewBst(nil)
	for _, n := range nums {
		tree.Insert(Int(n))
	}

	maxHeight := tree.GetMaxHeight()
	if maxHeight != 3 {
		t.Errorf("%d incorrect answer, correct answer is 3", maxHeight)
	}

}

func TestLevelOrder(t *testing.T) {
	nums := []int{3, 5, 4, 7, 2, 1}
	tree := NewBst(nil)
	for _, n := range nums {
		tree.Insert(Int(n))
	}
	t.Errorf("%v", tree.levelOrder())
}
