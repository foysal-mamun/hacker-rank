package main

import "fmt"

func main() {
	N := 0
	fmt.Scan(&N)

	var node *Node

	for i := 0; i < N; i++ {
		value := ""
		fmt.Scan(&value)
		node = node.Insert(value)
	}

	node.Display()
}

type Node struct {
	next  *Node
	value interface{}
}

func (n *Node) Insert(value interface{}) *Node {

	newNode := &Node{value: value}

	if n == nil {
		return newNode
	} else if n.next == nil {
		n.next = newNode
	} else {

		n.next.Insert(value)
	}

	return n
}

func (n *Node) Display() {
	for n != nil {
		fmt.Print(n.value, " ")
		n = n.next
	}
	fmt.Println("")
}
