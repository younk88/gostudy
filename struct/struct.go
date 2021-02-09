package main

import "fmt"

type BookInfo struct {
	title string
	author string
	code int
}

func settle_book(b BookInfo)  {
	b.code += 5
}

func settle_book_p(b *BookInfo)  {
	b.code += 5
}

func print_book(b BookInfo)  {
	fmt.Printf("\n==== %s ==========\n", b.title)
	fmt.Printf("code:\t%d\n", b.code)
	fmt.Printf("author:\t%s\n", b.author)
}

func main()  {
	book := BookInfo{"What is the mean to life", "forgot", 11}
	b2 := BookInfo{code: 22, author: "alex", title: "Efficent"}

	fmt.Println(book)
	fmt.Println(b2)

	print_book(book)
	print_book(b2)

	settle_book(book)
	settle_book_p(&b2)
	print_book(book)
	print_book(b2)
}