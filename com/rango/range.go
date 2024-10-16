package main

import "fmt"

var rangePow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
	for i, v := range rangePow {
		fmt.Printf("2**%d = %d\n", i, v)
	}

	for i := range rangePow {
		fmt.Println(i)
	}

	for _, v := range rangePow {
		fmt.Println(v)
	}
}
