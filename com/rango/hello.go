package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

var c, python, java bool
var fml bool

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	var stockcode = 123
	var enddate = "2020-12-31"
	var url = "Code=%d&endDate=%s"
	var target_url = fmt.Sprintf(url, stockcode, enddate)
	fmt.Printf(url, stockcode, enddate)
	fmt.Println()
	fmt.Println(target_url)

	var av, bv, cv = "xx", 1, true
	fmt.Println(av, bv, cv)

	a := 1
	fmt.Println(a, c, python, java)

	var n1, n2 int
	n1, n2 = a, a
	fmt.Println(n1, n2)

	fmt.Println(&a)
	fmt.Println(&n1)
	fmt.Println(&n2)

	var ax = "abc"
	fmt.Println("hello world", ax)

	arr := []int{1, 2, 3}
	fmt.Println(arr)
	var arr1 [4]int
	arr1[2] = 1
	fmt.Println(arr1)

	println(math.Mod(1, 2))
	println(math.Pi)

	fmt.Println(sum(1, 10))

	sx, sy := swap(ax, enddate)
	fmt.Println(sx, sy)

	fmt.Println(split(17))

	fmt.Printf("类型：%T 值：%v\n", ToBe, ToBe)
	fmt.Printf("类型：%T 值：%v\n", MaxInt, MaxInt)
	fmt.Printf("类型：%T 值：%v\n", z, z)

	fmt.Printf("%%b val: %b;\n"+
		"%%c val: %c;\n"+
		"%%d val: %d;\n"+
		"%%o val: %o;\n", 1, 'c', 12, 12)
}

/**
 * (x int, y int) int
 */
func sum(x, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}
