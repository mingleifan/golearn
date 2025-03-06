package main

import "fmt"

const (
	BJ = 0
	SH = 1
)

// 可以在const()中添加一个关键字 iota, 每行的iota都会累加1，第一行的iota的默认值是0
const (
	CD = iota * 10
	DY
	ZJ
	ZG
)

const (
	a, b = iota + 1, iota + 5
	c, d
	e, f

	g, h = iota + 10, iota + 100
	i, k
)

func main() {
	//常量
	const length int = 10

	fmt.Println("length = ", length)

	fmt.Println(BJ, SH)
	fmt.Println(CD, DY, ZJ, ZG)

}
