package bst

type Int int

type Item interface {
	Less(than Item) bool
}

func (i Int) Less(than Item) bool {
	return i < than.(Int)
}

type Bst struct {
	node  Item
	left  *Bst
	right *Bst
}

func NewBst(node Item) *Bst {
	return &Bst{node: node}
}

func (b *Bst) Insert(node Item) {

	if b.node == nil {
		b.node = node
		b.left = NewBst(nil)
		b.right = NewBst(nil)
		return
	}

	if node.Less(b.node) {
		if b.left != nil {
			b.left.Insert(node)
		} else {
			b.left = NewBst(node)
		}
	} else {
		if b.right != nil {
			b.right.Insert(node)
		} else {
			b.right = NewBst(node)
		}
	}

}

func (b *Bst) GetMaxHeight() int {
	if b.node == nil {
		return -1
	}

	leftDepth := 1
	rightDepth := 1

	leftDepth += b.left.GetMaxHeight()
	rightDepth += b.right.GetMaxHeight()

	if leftDepth > rightDepth {
		return leftDepth
	}

	return rightDepth
}

func (b *Bst) levelOrder() []int {

	values := []*Bst{}
	values = append(values, b)
	output := []int{}

	for len(values) > 0 {
		v := values[0]
		values = values[1:]

		if v.node != nil {
			output = append(output, int(v.node.(Int)))
		}

		if v.left != nil {
			values = append(values, v.left)
		}
		if v.right != nil {
			values = append(values, v.right)
		}

	}

	return output

}
