package main

import "fmt"

func main() {
	n := 0
	fmt.Scan(&n)

	var head *Node
	s := Solution{}

	for i := 0; i < n; i++ {
		var t int
		fmt.Scan(&t)

		head = s.Insert(head, t)
	}

	s.Display(head)
}

type Node struct {
	data int
	next *Node
}

type Solution struct {
}

func (s *Solution) Insert(head *Node, data int) *Node {
	n := &Node{data: data}

	if head == nil {
		return n
	} else if head.next == nil {
		head.next = n
	} else {
		s.Insert(head.next, data)
	}

	return head
}

func (s *Solution) Display(head *Node) {
	start := head
	for start != nil {
		fmt.Print(start.data, " ")
		start = start.next
	}
	fmt.Println("")
}
