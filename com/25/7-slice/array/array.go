package main

import "fmt"

func main() {

	//固定长度数组
	var arr1 [10]int
	arr1[1] = 1
	for i := 0; i < len(arr1); i++ {
		fmt.Println(arr1[i])
	}

	arr2 := [10]int{1, 2, 3}
	for index, val := range arr2 {
		fmt.Println("index =", index, "value =", val)
	}

	arr3 := [2]float64{1.1, 2.324}
	for i := 0; i < 2; i++ {
		fmt.Println(arr3[i])
	}

	//查看数组的数据类型
	fmt.Printf("arr1 types = %T\n", arr1)
	fmt.Printf("arr1 types = %T\n", arr3)

	arr4 := [3]int{1, 2, 3}
	println("------")
	printArrByFixedLen(arr4)
	println("------")
	for i := 0; i < len(arr4); i++ {
		fmt.Println("arr4 val", arr4[i])
	}
	println("------")

}

func printArrByFixedLen(paramArr [3]int) {
	//值拷贝，不影响实际传入的数组
	for i, v := range paramArr {
		fmt.Println("index =", i, "value =", v)
	}
	paramArr[0] = 11
}
