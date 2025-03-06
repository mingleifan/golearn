package main

import "fmt"

func main() {

	//动态数组, 切片 slice
	arr1 := []int{1, 2, 3, 4, 5}

	fmt.Printf("arr1 type is %T\n", arr1)

	printArr(arr1)
	arr1[0] = 100
	printArr(arr1)
	//error -> panic: runtime error: index out of range [9] with length 5
	//arr1[9] = 20
	//printArr(arr1)

	// 创建一个长度为 3，容量为 5 的整数切片
	arr2 := make([]int, 3, 5)
	fmt.Printf("切片: %v, 长度: %d, 容量: %d\n", arr2, len(arr2), cap(arr2))
	printArr(arr2)

	fmt.Println("----------")

	sliceAppend()
}

func printArr(arr []int) {
	//引用传递
	//_表示匿名的变量
	for _, val := range arr {
		fmt.Println(val)
	}
}

func sliceAppend() {
	// 创建一个长度为 2，容量为 3 的整数切片
	slice := make([]int, 2, 3)
	fmt.Printf("初始状态: 长度 %d, 容量 %d\n", len(slice), cap(slice))

	// 追加一个元素，此时长度小于容量
	slice = append(slice, 1)
	fmt.Printf("追加一个元素后: 长度 %d, 容量 %d\n", len(slice), cap(slice))

	// 再追加一个元素，此时长度等于容量，会触发扩容
	slice = append(slice, 2)
	fmt.Printf("再追加一个元素后: 长度 %d, 容量 %d\n", len(slice), cap(slice))
}
