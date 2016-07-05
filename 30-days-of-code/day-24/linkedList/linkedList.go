package linkedList

type Int int
type Item interface {
	IsEqual(than Item) bool
}

func (i Int) IsEqual(than Item) bool {
	return i == than.(Int)
}

type Node struct {
	Data Item
	Next *Node
}

func NewNode(data Item) *Node {
	return &Node{Data: data}
}

func (n *Node) Insert(data Item) {
	node := NewNode(data)

	if n.Data == nil {
		n.Data = data
		n.Next = NewNode(nil)
		return
	} else if n.Next == nil {
		n.Next = node
	} else {
		n.Next.Insert(data)
	}
}

func (n *Node) Traverse() []int {
	elems := []int{}
	ll := n
	for ll.Next != nil {
		elems = append(elems, int(ll.Data.(Int)))
		ll = ll.Next
	}

	return elems
}

func (n *Node) RemoveDuplicates() {

	current := n
	for current.Next.Data != nil {
		if current.Next.Data.IsEqual(current.Data) {
			current.Next = current.Next.Next
		} else {
			current = current.Next
		}
	}

}
