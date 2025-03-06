package main

import "fmt"

func main() {
	var a = 1

	changeValue(a)

	fmt.Println("a = ", a) // a = 1 并不为 10

	changeValue2(&a)

	fmt.Println("a = ", a) // a = 10 used pointer
}

func changeValue(p int) {
	fmt.Println("changeValue p = ", p)
	p = 10
	fmt.Println("changeValue p = ", p)
}

func changeValue2(p *int) {
	fmt.Println("changeValue2 * p = ", *p)
	*p = 10
	fmt.Println("changeValue2 * p = ", *p)
}
