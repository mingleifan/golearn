package main

import (
	"fmt"
)

// make(chan Type)   //等价于make(chan Type, 0)
// make(chan Type, capacity)

// channel <- value  //发送value到channel
// <- channel 		 //接收并将其丢弃
// x:= <- channel 	 //从channel中接受数据，并赋值给x
// x, ok <- channel  //功能同上，同时检查通道是否已关闭或者是否为空
func main() {

	c := make(chan int)

	go func() {
		defer fmt.Println("goroutine end...")
		fmt.Println("goroutine running...")

		c <- 666 //将666 发送给c
	}()

	num := <-c //从c中接收数据，并赋值给num
	fmt.Println("num = ", num)
	fmt.Println("main goroutine end...")
}
