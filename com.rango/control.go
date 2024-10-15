package main

import (
	"fmt"
	"math"
)

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
	//初始化语句和后置语句是可选的。
	for sum < 100 {
		sum += sum
	}
	fmt.Println(sum)
	for {
		fmt.Println(sum + 1)
		if sum > 100 {
			break
		}
	}

	num := 16
	squareRoot := math.Sqrt(float64(num))
	fmt.Println("The square root of", num, "is", squareRoot)

	fmt.Println(pow(3, 2, 10), pow(3, 3, 20))
}

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else if v == lim {

	}
	return lim
}
