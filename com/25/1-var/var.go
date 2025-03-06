package main

import (
	"fmt"
	"math"
	"strconv"
)

var (
	z int = 1
)

var (
	x int = 3
)

var fff = 4

func main() {
	fmt.Println(math.Pi)

	fmt.Println(fff)

	fmt.Println(add(1, 2))
	fmt.Println(addIntString(1, "x"))
	r1, r2 := multiReturn("x", "y")
	fmt.Println(r1, r2)

	fmt.Println(z, x)

	bb := 1
	fmt.Println(bb)

	// := 只能在函数体内使用
	c := 3.14
	fmt.Printf("c = %g, type of c = %T\n", c, c)

	var xx, yy, zz = 100, 100, 100
	fmt.Println(xx, yy, zz)

	xxx, yyy, zzz := 10, 10, 10
	fmt.Println(xxx, yyy, zzz)

}

func add(x, y int) int {
	return x + y
}

func addIntString(x int, y string) int {
	yi, err := strconv.Atoi(y)
	if err != nil {
		return x
	} else {
		return x + yi
	}
}

// 多返回值
func multiReturn(x, y string) (string, string) {
	return y, x
}
