package main

import (
	"fmt"
	"math"
)

var rangePow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
	for i, v := range rangePow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
	fmt.Println("-------")

	for i := range rangePow {
		fmt.Println(i)
	}
	fmt.Println("-------")

	for _, v := range rangePow {
		fmt.Println(v)
	}
	fmt.Println("-------")

	picture := Pic(5, 5)
	for _, row := range picture {
		for _, pixel := range row {
			fmt.Printf("%d ", pixel)
		}
		fmt.Println()
	}

}

func Pic(dx, dy int) [][]uint8 {
	picture := make([][]uint8, dy)
	for y := 0; y < dy; y++ {
		picture[y] = make([]uint8, dx)
		for x := 0; x < dx; x++ {
			picture[y][x] = uint8(math.Pow(float64(x), float64(y)))
		}
	}
	return picture
}
