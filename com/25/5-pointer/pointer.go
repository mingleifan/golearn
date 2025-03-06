package main

import "fmt"

func main() {
	var a = 1

	changeValue(a)

	fmt.Println("a =", a) // a = 1 并不为 10

	changeValue2(&a)

	fmt.Println("a =", a) // a = 10 used pointer

	a, b := 1, 2

	swap(&a, &b)

	l2Pointer()
}

func changeValue(p int) {
	fmt.Println("changeValue p =", p)
	p = 10
	fmt.Println("changeValue p =", p)
}

func changeValue2(p *int) {
	fmt.Println("changeValue2 *p =", *p)
	*p = 10
	fmt.Println("changeValue2 *p =", *p)
}

func swap(a, b *int) {
	fmt.Println("exec swap func...")
	fmt.Println("a =", *a, "b =", *b)
	tmp := *a
	*a = *b
	*b = tmp
	fmt.Println("a =", *a, "b =", *b)
}

// 二级指针
func l2Pointer() {

	fmt.Println("exec show l2-pointer")

	a := 1
	p := &a
	fmt.Println("a =", a, "&a =", &a, "p =", p, "*p =", *p)

	pp := &p
	fmt.Println("pp =", pp, "&p =", &p, "*pp=", *pp)
}
