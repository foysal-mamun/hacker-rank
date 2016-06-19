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
	count int
}

func NewStack() *Stack {
	return &Stack{}
}

func (s *Stack) Push(v interface{}) {
	s.value = append(s.value[:s.count], v)
	s.count++
}

func (s *Stack) Pop() interface{} {
	if s.count == 0 {
		return nil
	}

	s.count--
	return s.value[s.count]
}

type Queue struct {
	value []interface{}
	front int
	rear  int
}

func NewQueue() *Queue {
	return &Queue{}
}

func (q *Queue) Enqueue(v interface{}) {
	q.value = append(q.value, v)
	q.rear++

}

func (q *Queue) Dequeue() interface{} {
	if q.rear == 0 {
		return nil
	}

	value := q.value[q.front]

	q.value = q.value[q.front+1:]
	q.rear--

	return value
}
