package main

import (
	"fmt"
	"io"
	"os"
)

/**
 * 变量 ->  1. type; 2.value
 * type -> 1. static type(int, string...); 2.concrete type(interface 所指向的具体数据类型)
 * go中的变量 包含了 type和value的这样一个pair对
 */
func main() {

	var a string

	//pair<type:string(static type), value:"abced">
	a = "abced"

	//pair<type:string, value:"abced">
	var allType interface{}
	allType = a

	str, ok := allType.(string)
	fmt.Println(str, ok)

	//tty: pair<type:*os.File, value:"/dev/tty"文件描述符>
	tty, err := os.OpenFile("/dev/tty", os.O_RDWR, 0)
	if err != nil {
		fmt.Println("open file error", err)
		return
	}

	// r: pair<type:, value:>
	var r io.Reader
	// r: pair<type:*os.File, value:"/dev/tty"文件描述符>
	r = tty

	// w: pair<type:, value:>
	var w io.Writer
	// w: pair<type:*os.File, value:"/dev/tty"文件描述符>
	w = r.(io.Writer)

	write, err := w.Write([]byte("HELLO THIS IS A TEST!!!\n"))
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("write type is %T, content: %v\n", write, write)
	}

}
