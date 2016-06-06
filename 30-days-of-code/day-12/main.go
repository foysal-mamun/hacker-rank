package main

import "fmt"

func main() {
	firstName, lastName := "", ""
	id, numScores := 0, 0
	scores := []int{}

	fmt.Scan(&firstName)
	fmt.Scan(&lastName)
	fmt.Scan(&id)
	fmt.Scan(&numScores)

	for i := 0; i < numScores; i++ {
		tmpScore := 0
		fmt.Scan(&tmpScore)
		scores = append(scores, tmpScore)
	}

	s := NewStudent(firstName, lastName, id, scores)
	s.PrintPerson()
	fmt.Println("Grade:", s.Calculate())

}

type Person struct {
	firstName string
	lastName  string
	id        int
}

func NewPerson(firstName string, lastName string, id int) Person {
	p := Person{}
	p.firstName = firstName
	p.lastName = lastName
	p.id = id
	return p
}

func (p Person) PrintPerson() {
	fmt.Println("Name:", p.lastName+",", p.firstName)
	fmt.Println("ID:", p.id)
}

type Student struct {
	Person
	TestScores []int
}

func NewStudent(firstName string, lastName string, id int, scores []int) Student {
	s := Student{}
	s.Person = NewPerson(firstName, lastName, id)
	s.TestScores = scores

	return s
}

func (s *Student) Calculate() string {
	avg := 0
	for _, v := range s.TestScores {
		avg += v
	}

	return s.grading(avg / len(s.TestScores))
}

func (s *Student) grading(a int) string {
	switch {
	case 90 <= a && a <= 100:
		return "O"
	case 80 <= a && a <= 90:
		return "E"
	case 70 <= a && a <= 80:
		return "A"
	case 55 <= a && a <= 70:
		return "P"
	case 40 <= a && a <= 50:
		return "D"
	case a <= 40:
		return "T"
	}

	return ""
}
