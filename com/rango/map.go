package main

import (
	"fmt"
	"golang.org/x/tour/wc"
	"strings"
)

func main() {
	m = make(map[string]Vertex)
	k := "Bell Labs"
	v := Vertex{40.68433, -74.39967}
	m[k] = v
	vertex := m[k]
	// 存入m的value的值(m[k])是v的副本，取出的vertex也是(m[k])的副本
	fmt.Println(m[k], v, vertex)
	v.Long = -74.39666
	fmt.Println(m[k], v, vertex)

	//真正修改映射中的value, 需要重新存入映射
	m[k] = v

	fmt.Println(m)
	delete(m, k)
	fmt.Println(m)

	gv, ok := m[k]
	fmt.Println("值：", gv, "是否存在？", ok)

	wc.Test(WordCount)

}

func WordCount(s string) map[string]int {
	words := strings.Fields(s)
	wordsMap := make(map[string]int)
	for _, word := range words {
		wordsMap[word]++
	}
	return wordsMap
}

var m = map[string]Vertex{}

var m2 = map[string]Vertex{
	"Bell Labs": {40.68433, -74.39967},
	"Google":    {37.42202, -122.08408},
}

type Vertex struct {
	Lat, Long float64
}
