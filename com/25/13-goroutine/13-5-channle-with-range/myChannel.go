package main

import "fmt"

func main() {

	c := make(chan int)

	go func() {
		for i := 0; i < 5; i++ {
			c <- i
		}
		close(c)
	}()

	//for {
	//	if data, ok := <-c; ok {
	//		fmt.Println(data)
	//	} else {
	//		fmt.Println("channel closed, data", data)
	//		break
	//	}
	//}
	//也可以用下面这种方式
	for data := range c {
		fmt.Println(data)
	}

	fmt.Println("main finished...")

}
