package main

import "fmt"

func main() {
	title, author := "", ""
	price := 0

	fmt.Scan(&title)
	fmt.Scan(&author)
	fmt.Scan(&price)

	var ba BookAst
	ba = NewBook(title, author, price)
	ba.Dispaly()

}

type BookAst interface {
	Dispaly()
}

type Book struct {
	title, author string
	price         int
}

func NewBook(title, author string, price int) Book {
	b := Book{}
	b.title = title
	b.author = author
	b.price = price

	return b
}

func (b Book) Dispaly() {
	fmt.Println("Title:", b.title)
	fmt.Println("Author:", b.author)
	fmt.Println("Price:", b.price)
}
