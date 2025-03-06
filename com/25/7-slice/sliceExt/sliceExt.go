package main

import "fmt"

func main() {

	//声明slice1是一个切片，并初始化，默认值1，2，3.长度len是3
	slice1 := []int{1, 2, 3}
	nilSlice(slice1, 1)
	fmt.Printf("slice2 len = %d, value = %v, type:%T \n", len(slice1), slice1, slice1)
	//声明slice2是一个切片，但是没有初始化，默认值是nil，长度len是0
	var slice2 []int
	nilSlice(slice2, 2)
	fmt.Printf("slice2 len = %d, value = %v \n", len(slice2), slice2)
	slice2 = make([]int, 3) //开辟三个空间
	nilSlice(slice2, 2)
	fmt.Println("slice 2 do make()")
	slice2[0] = 100
	fmt.Printf("slice2 len = %d, value = %v \n", len(slice2), slice2)
	//声明slice3是一个切片，但是没有初始化，默认值是nil，长度len是3
	//var slice3 []int = make([]int, 3)
	slice3 := make([]int, 3)
	fmt.Printf("slice3 len = %d, value = %v \n", len(slice3), slice3)

}

func nilSlice(slice []int, name int) bool {
	if slice == nil {
		fmt.Printf("slice%d is nil\n", name)
		return true
	} else {
		fmt.Printf("slice%d is not nil\n", name)
		return false
	}
}
