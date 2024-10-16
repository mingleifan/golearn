package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var arr [2]string
	arr[0] = "hello"
	arr[1] = "world"
	fmt.Println(arr[0], arr[1])
	fmt.Println(arr)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	//slice 左包右不包
	s := primes[1:4]
	fmt.Println(s)

	//切片就像数组的引用 切片并不存储任何数据，它只是描述了底层数组中的一段。
	//更改切片的元素会修改其底层数组中对应的元素。
	//和它共享底层数组的切片都会观测到这些修改。
	names := [4]string{"John", "Paul", "George", "Ringo"}
	fmt.Println(names)

	fmt.Println("---------")
	a := names[0:2]
	b := names[1:3]
	fmt.Println(a)
	for i := 0; i < 10; i++ {
		a = append(a, strconv.Itoa(i))
	}
	fmt.Println(a)
	fmt.Println(names)
	fmt.Println("---------")

	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)

	//
	dynamicArr := []int{1, 2, 3}
	appendIntArr := append(dynamicArr, 4, 5)
	fmt.Println(dynamicArr, appendIntArr)

	//
	structs := []struct {
		x, y, z int
	}{
		{1, 2, 3},
		{2, 3, 4},
		{3, 4, 5},
	}
	ns := append(structs, struct {
		x, y, z int
	}{4, 5, 6})
	fmt.Println(structs, ns)

	//切片的默认行为
	slice1 := []int{2, 3, 4, 5, 6, 9, 10, 29, 41}
	slice1 = slice1[1:4]
	fmt.Println(slice1)

	slice1 = slice1[:2]
	fmt.Println(slice1)

	slice1 = slice1[1:]
	fmt.Println(slice1)

	//切片的长度和容量
	slice2 := []int{2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(slice2, len(slice2), cap(slice2))

	slice2 = slice2[:0]
	fmt.Println(slice2, len(slice2), cap(slice2))

	slice2 = slice2[:4]
	fmt.Println(slice2, len(slice2), cap(slice2))

	slice2 = slice2[2:]
	fmt.Println(slice2, len(slice2), cap(slice2))

	for i := 0; i < 10; i++ {
		slice2 = append(slice2, i)
	}
	fmt.Println(slice2, len(slice2), cap(slice2))

	//nil切片
	var nilSlice []int
	fmt.Println(nilSlice, len(nilSlice), cap(nilSlice))
	if nilSlice == nil {
		fmt.Println("nil!")
	}

	//make创建切片
	slice3 := make([]int, 10) //len(slice3)=5
	fmt.Println(slice3, len(slice3), cap(slice3))

	slice4 := make([]int, 0, 5) //len(slice4)=0, cap(slice4)=5
	fmt.Println(slice4, len(slice4), cap(slice4))

	slice4 = slice4[:cap(slice4)]
	fmt.Println(slice4, len(slice4), cap(slice4))
	slice4 = slice4[1:]
	fmt.Println(slice4, len(slice4), cap(slice4))

	//切片的切片
	slicesOfSlices()
}

func slicesOfSlices() {
	board := [][]string{
		[]string{"-", "-", "-"},
		[]string{"-", "-", "-"},
		[]string{"-", "-", "-"},
	}
	// 两个玩家轮流打上 X 和 O
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"
	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}
