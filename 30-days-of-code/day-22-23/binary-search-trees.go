package main

import "fmt"

func main() {
	N := 0

	fmt.Scan(&N)

	nums := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&nums[i])
	}

	tree := New(nil)
	for _, v := range nums {
		tree.insert(Int(v))
	}

	// 9 20 50 35 44 9 15 62 11 13 // 4
	// 7 3 5 2 1 4 6 7 // 3
	fmt.Println(tree.getMaxHeight())
}

type Item interface {
	Less(than Item) bool
}
type Int int

func (i Int) Less(than Item) bool {
	return i < than.(Int)
}

type Bst struct {
	data  Item
	left  *Bst
	right *Bst
}

func New(data Item) *Bst {
	b := &Bst{}
	b.data = data
	return b
}

func (b *Bst) insert(data Item) {

	if b.data == nil {
		b.data = data
		b.left = New(nil)
		b.right = New(nil)
		return
	}

	if data.Less(b.data) {
		if b.left != nil {
			b.left.insert(data)
		} else {
			b.left = New(data)
		}
	} else {
		if b.right != nil {
			b.right.insert(data)
		} else {
			b.right = New(data)
		}
	}

}

func (b *Bst) getMaxHeight() int {

	if b.data == nil {
		return -1
	}

	leftHeight := 1
	rightHeight := 1

	leftHeight += b.left.getMaxHeight()
	rightHeight += b.right.getMaxHeight()

	if leftHeight > rightHeight {
		return leftHeight
	}

	return rightHeight
}
