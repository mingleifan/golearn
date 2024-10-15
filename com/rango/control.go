package main

import (
	"fmt"
	"math"
	"runtime"
	"time"
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

	switch os := runtime.GOOS; os {
	case "linux":
		fmt.Println("LINUX.")
	default:
		//default
		fmt.Printf("%s.\n", os)
	}

	todayIsWhatDay()

	// defer语句 会将函数推迟到外层函数返回之后执行。
	defer toConditionSwitch()

	deferStack()

}

// defer栈
func deferStack() {
	fmt.Println("counting")
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("done...")
}

func toConditionSwitch() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		println("早上好")
	case t.Hour() < 18:
		println("下午好")
	default:
		println("晚上好")
	}
}

func todayIsWhatDay() {
	fmt.Println("周五是哪天？")
	today := time.Now().Weekday()
	switch time.Friday {
	case today + 0:
		fmt.Println("今天。")
	case today + 1:
		fmt.Println("明天。")
	case today + 2:
		fmt.Println("后天。")
	case today + 3:
		fmt.Println("外天。")
	default:
		fmt.Println("很多天后。")
	}
}

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else if v == lim {

	}
	return lim
}
