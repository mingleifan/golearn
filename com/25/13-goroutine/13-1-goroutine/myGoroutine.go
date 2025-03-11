package main

import (
	"fmt"
	"runtime"
	"time"
)

func newTask() {
	i := 0
	for {
		i++
		fmt.Printf("new Goroutine : i = %d\n", i)
		time.Sleep(1 * time.Second)
	}
}

// 退出当前goroutine:[ runtime.Goexit() ]
func main() {
	//创建一个goroutine 执行newTask() 流程
	go newTask()

	//fmt.Println("main goroutine exit")

	//i := 0
	//for {
	//	i++
	//	fmt.Printf("main Goroutine : i = %d\n", i)
	//	time.Sleep(1 * time.Second)
	//}

	//用go创建承载一个形参为空，返回值为空的一个函数
	go func() {
		defer fmt.Println("A.defer")
		//return
		func() {
			defer fmt.Println("B.defer")

			//退出当前子函数
			//return

			//退出当前goroutine
			runtime.Goexit()

			fmt.Println("B")
		}()
		fmt.Println("A")
	}()

	go func(a, b int) bool {
		fmt.Printf("a = %v, b = %v\n", a, b)
		return true
	}(10, 20)

	for {
		time.Sleep(1 * time.Second)
	}

}
