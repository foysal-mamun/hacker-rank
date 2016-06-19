package main

import "fmt"

func main() {

	s := ""
	fmt.Scan(&s)

	stck := NewStack()
	quue := NewQueue()

	for _, v := range s {
		stck.Push(string(v))
		quue.Enqueue(string(v))
	}

	isPalindrom := true
	ls := len(s)
	for i := 0; i < ls/2; i++ {
		if stck.Pop() != quue.Dequeue() {
			isPalindrom = false
			break
		}
	}

	if isPalindrom {
		fmt.Printf("The word, %s, is a palindrom.\n", s)
	} else {
		fmt.Printf("The word, %s, is not a palindrom.\n", s)
	}
}

type Stack struct {
	value []interface{}
}

func NewStack() *Stack {
	return &Stack{}
}

func (s *Stack) Push(v interface{}) {
	s.value = append(s.value, v)
}

func (s *Stack) Pop() interface{} {
	l := len(s.value) - 1
	if l+1 == 0 {
		return nil
	}

	v := s.value[l]
	s.value = s.value[:l]
	return v
}

type Queue struct {
	value []interface{}
}

func NewQueue() *Queue {
	return &Queue{}
}

func (q *Queue) Enqueue(v interface{}) {
	q.value = append(q.value, v)
}

func (q *Queue) Dequeue() interface{} {
	if len(q.value) == 0 {
		return nil
	}

	value := q.value[0]

	q.value = q.value[1:]

	return value
}
