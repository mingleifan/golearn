package main

import "fmt"

func main() {

	fmt.Println("leaning append.....")

	slice1 := make([]int, 3, 4)
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(slice1), cap(slice1), slice1)

	slice1 = append(slice1, 1)
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(slice1), cap(slice1), slice1)

	slice2 := append(slice1, 2)
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(slice1), cap(slice1), slice1)
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(slice2), cap(slice2), slice2)

	slice3 := make([]int, 3)
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(slice3), cap(slice3), slice3)
	slice3 = append(slice3, 1)
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(slice3), cap(slice3), slice3)

	fmt.Println("leaning substring.....")

	s1 := []int{1, 2, 3, 4, 5}
	//[0, 2] 左闭右开
	s2 := s1[0:2]
	fmt.Println(s2)
	s1[0] = 10
	fmt.Println(s2)

	s3 := s1[:]
	fmt.Println(s3)

	s4 := s1[3:]
	fmt.Println(s4)

	s5 := s1[:2]
	fmt.Println(s5)

	//copy
	s6 := make([]int, 3)
	//将s1中的值，依次拷贝到s6中
	copy(s6, s1)
	s1[0] = 100
	fmt.Println(s1)
	fmt.Println(s6)

}
