package main

import "fmt"

func main() {

	deferVal := 1
	//defer 结束前执行 main end2先执行，栈帧
	defer fmt.Println("main end1")
	defer fmt.Println("main end2")
	defer deferFunc1(&deferVal)

	fmt.Println("main: hello go 1")
	fmt.Println("main: hello go 2")

	returnAndDefer()

	fmt.Println("main: deferVal:", deferVal)

}

func returnAndDefer() int {
	fmt.Println("returnAndDefer func called...")
	defer deferFunc()
	return returnFunc()
}

func deferFunc() int {
	fmt.Println("defer func called...")
	return 0
}

func deferFunc1(val *int) {
	*val = 1000
	fmt.Println("deferFunc1 func called...", "*val =", *val)
}

func returnFunc() int {
	fmt.Println("return func called...")
	return 0
}
