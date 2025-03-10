package main

import "fmt"

// interface{} 万能数据类型
func myFunc(arg interface{}) {
	fmt.Println("myFunc is called...")
	fmt.Println(arg)
	//interface{}该如何区分 此时引用的底层数据类型是什么？
	//给 interface{}提供 “断言“
	value, ok := arg.(string)
	if ok {
		fmt.Println("arg is string type")
	} else {
		fmt.Printf("arg is not string type, value type is %T\n", value)
	}

}

type Book struct {
	author string
}

func main() {

	book := Book{"author"}
	myFunc(book)
	myFunc(100)
	myFunc("abc")
	myFunc(3.14)

}
