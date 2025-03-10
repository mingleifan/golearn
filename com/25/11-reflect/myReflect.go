package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id   int
	Name string
	Age  int
}

func reflectNum(arg interface{}) {

	fmt.Println("type: ", reflect.TypeOf(arg))
	fmt.Println("value: ", reflect.ValueOf(arg))

}

func (user User) Call() {
	fmt.Println("user is called ...")
	fmt.Printf("%v\n", user)
}

func DoFieldAndMethod(input interface{}) {
	//获取input的type
	inputType := reflect.TypeOf(input)
	fmt.Println("inputType is:", inputType.Name())
	//获取input的value
	inputValue := reflect.ValueOf(input)
	fmt.Println("inputValue is:", inputValue)
	//通过type 获取里面的字段
	//1. 获取interface打reflect.Type, 通过Type得到NumField，遍历
	//2. 得到每个field数据类型
	//3. 用过field有一个Integer()方法得到对应的值
	fmt.Println("------fields------")
	for i := 0; i < inputType.NumField(); i++ {
		field := inputType.Field(i)
		value := inputValue.Field(i).Interface()
		fmt.Printf("%s(%v) = %v\n", field.Name, field.Type, value)
	}

	fmt.Println("------methods------")
	//通过type 获取里面的方法，调用
	for i := 0; i < inputType.NumMethod(); i++ {
		m := inputType.Method(i)
		fmt.Printf("%s: %v \n", m.Name, m.Type)
		m.Func.Call([]reflect.Value{reflect.ValueOf(input)})
	}
}

func main() {

	var num float64 = 1.2345
	reflectNum(num)

	fmt.Println("-------------")

	user := User{1, "wang", 18}

	DoFieldAndMethod(user)
}
