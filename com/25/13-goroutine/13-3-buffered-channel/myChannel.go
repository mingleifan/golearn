package main

import (
	"fmt"
	"time"
)

func main() {

	c := make(chan int, 3)
	fmt.Printf("len(c) = %v, cap(c) = %v\n", len(c), cap(c))

	go func() {
		defer fmt.Println("sub goroutine end...")
		for i := 0; i < 4; i++ {
			c <- i
			fmt.Printf("sub goroutine is running, send data = %v, len(c) = %v, cap(c) = %v\n", i, len(c), cap(c))
		}
		time.Sleep(2 * time.Second)
	}()

	time.Sleep(2 * time.Second)

	for i := 0; i < 4; i++ {
		num := <-c
		fmt.Printf("num = %v\n", num)
	}

	fmt.Println("main goroutine end...")

}
