package main

import "fmt"

// channel不像文件一样需要经常去关闭，只有当你确实没有任何发送数据了，或者你想显式的结束range循环之类的，才去关闭channel
// 关闭channel后，无法向channel再发送数据(引发 panic 错误后导致接受立即返回零值)
// 关闭channel后，无法继续从channel接受数据
// 对于nil channel，无论收发都会被阻塞
func main() {

	c := make(chan int)

	go func() {
		for i := 0; i < 5; i++ {
			c <- i
		}
		close(c)
		//close可以关闭一个channel
	}()

	for {
		//ok为true:意味着通道处于打开状态，并且成功从通道中接收到了数据。也就是说，在接收数据的时候，通道没有被关闭，还有数据可以接收。
		//ok为false:表示通道已经被关闭，并且通道里没有数据了。
		//当通道关闭后，若尝试从通道接收数据，会得到通道元素类型的零值，同时 ok 会被设置为 false。
		if data, ok := <-c; ok {
			fmt.Println(data)
		} else {
			fmt.Println("channel closed, data", data)
			break
		}
	}

	fmt.Println("main finished...")

}
