package main

import "fmt"

// 声明一种新的数据类型，是int的一个别名
type myInt int

// Book 定义一个结构体
type Book struct {
	name   string
	author string
}

func main() {

	var a myInt = 10
	fmt.Printf("a = %d, type of a is %T\n", a, a)

	var book1 Book
	book1.name = "book1"
	book1.author = "author1"

	fmt.Printf("%v\n", book1)

	changeBook(book1)

	fmt.Printf("%v\n", book1)

	changeBook2(&book1)
	fmt.Printf("%v\n", book1)
}

func changeBook(book Book) {
	//传递一个book的副本
	book.name = "new name"
	book.author = "new author"
}

func changeBook2(book *Book) {
	book.author = "new author(pointer)"
}
