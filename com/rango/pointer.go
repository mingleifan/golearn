package main

import "fmt"

func main() {
	x, y := 42, 2701

	p := &x         // 指向 x
	fmt.Println(p)  // 查看 p指针 的值
	fmt.Println(*p) // 通过指针读取 x 的值

	*p = 21        // 通过指针设置 x 的值
	fmt.Println(x) // 查看 x 的值

	p = &y         // 指向 y
	fmt.Println(p) // 查看 p指针 的值
	*p = *p / 37   // 通过指针对 y 进行除法运算
	fmt.Println(y) // 查看 y 的值

	fmt.Println(x, y)

}
