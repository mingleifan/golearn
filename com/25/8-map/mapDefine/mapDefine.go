package main

import (
	"fmt"
)

// Go 语言的 map 实现是高度优化的，其内部的扩容机制是自动的，并且细节对用户是隐藏的。
// 这意味着 Go 语言的设计者认为用户通常不需要关心 map 的具体容量，只需要关注 map 中存储的键值对数量即可。
// 可以使用 len 函数获取 map 中当前存储的键值对数量
func main() {

	var map1 map[string]int
	if map1 == nil {
		fmt.Println("map1 is nil")
	}
	map1 = make(map[string]int, 2)
	map1["java"] = 1
	map1["go"] = 2
	map1["python"] = 3
	fmt.Printf("map len = %d \n", len(map1))

	fmt.Println(map1)

	map2 := map[string]string{
		"1": "java",
		"2": "go",
		"3": "php",
	}
	fmt.Println(map2)

}
