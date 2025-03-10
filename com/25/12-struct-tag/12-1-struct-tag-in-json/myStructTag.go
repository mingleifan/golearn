package main

import (
	"encoding/json"
	"fmt"
)

type Movie struct {
	Title  string   `json:"title"`
	Year   int      `json:"year"`
	Price  int      `json:"RMB"`
	Actors []string `json:"actors"`
}

// 结构体标签 json编解码；orm映射关系
func main() {
	movie := Movie{"喜剧之王", 2000, 10, []string{"周星驰", "张柏芝"}}

	//编码 结构体-> json
	marshal, err := json.Marshal(movie)
	if err != nil {
		fmt.Println("json marshal error", err)
	} else {
		fmt.Printf("json marshal result: %s\n", marshal)
	}

	//解码 json -> 结构体
	myMovie := Movie{}
	err = json.Unmarshal(marshal, &myMovie)
	if err != nil {
		fmt.Println("json unmarshal error", err)
	} else {
		fmt.Printf("%v\n", myMovie)
	}

}
