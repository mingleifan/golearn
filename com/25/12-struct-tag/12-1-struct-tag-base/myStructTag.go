package main

import (
	"fmt"
	"reflect"
)

// 结构体标签

type resume struct {
	Name string `info:"name" doc:"我的名字"`
	Sex  string `info:"sex"`
}

func findTag(str interface{}) {
	elem := reflect.TypeOf(str).Elem()
	for i := 0; i < elem.NumField(); i++ {
		field := elem.Field(i)
		tag := field.Tag
		fmt.Printf("tag(info): %v, tag(doc): %v \n", tag.Get("info"), tag.Get("doc"))
	}
}

func findTag2(str interface{}) {
	inputType := reflect.TypeOf(str)
	for i := 0; i < inputType.NumField(); i++ {
		tag := inputType.Field(i).Tag
		fmt.Printf("tag(info): %v, tag(doc): %v \n", tag.Get("info"), tag.Get("doc"))
	}
}

func main() {

	re := resume{"name-v", "sex-v"}
	findTag(&re)
	findTag2(re)

}
