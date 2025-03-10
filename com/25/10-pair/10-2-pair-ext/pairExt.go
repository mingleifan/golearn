package main

import "fmt"

// Book 实现了Reader接口和Writer接口
type Book struct {
}

type Reader interface {
	ReadBook()
}

type Writer interface {
	WriteBook()
}

func (this *Book) ReadBook() {
	fmt.Println("read book")
}

func (this *Book) WriteBook() {
	fmt.Println("write book")
}

func main() {

	// b: pair<type:Book, value:book{}地址>
	b := &Book{}

	// r: pair<type:, value:>
	var r Reader
	// r: pair<type:Book, value:book{}地址>
	r = b

	r.ReadBook()

	var w Writer
	// r: pair<type:Book, value:book{}地址>
	w = r.(Writer) //此处的断言为什么会成功？ 因为 w r 具体的type是一致的

	w.WriteBook()

}
