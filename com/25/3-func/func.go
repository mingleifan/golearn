package main

import "fmt"

func main() {

	fmt.Println(multiParam("hello", 1))
	fmt.Println(multiReturn())
	fmt.Println(multiReturn2())
	fmt.Println(multiParamAndReturn(1, 2))
}

func multiParam(a string, b int) int {
	fmt.Println("--func multiParam--")
	fmt.Println(a, b)
	c := 100
	return c
}

func multiReturn() (int, int) {
	fmt.Println("--func multiReturn--")
	return 1, 2
}

func multiReturn2() (x int, y int) {
	fmt.Println("--func multiReturn2--")
	x = 10
	y = 10
	return x, y
}

func multiParamAndReturn(a, b int) (x, y, z int, ff string) {
	fmt.Println("--func multiParamAndReturn--")
	fmt.Println(a, b)
	x += 1
	y -= 2
	ff = "nmlgb"
	return x, y, z, ff
}
