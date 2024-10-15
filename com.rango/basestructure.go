package main

import "fmt"

const (
	Big   = 1 << 100
	Small = Big >> 99
)

func main() {
	fmt.Println("Hello, world!")

	var i int = 64
	var f float64 = float64(i)
	var u uint = uint(f)
	fmt.Println(i, f, u)

	x := 42
	y := float64(x)
	z := uint(y)
	fmt.Println(x, y, z)
	fmt.Println()

	const fanminglei = "FML"
	fmt.Println(fanminglei)

	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}

func needInt(x int) int {
	return x*10 + 1
}

func needFloat(x float64) float64 {
	return x * 0.1
}
