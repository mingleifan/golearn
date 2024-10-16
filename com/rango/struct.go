package main

import "fmt"

var (
	v1 = Person{9, 8, "fml", 26}
	v2 = Person{X: 10, Y: 20}
	v3 = Person{}
	vp = &Person{2, 3, "rg", 8}
)

func main() {
	person := Person{1, 2, "rango", 26}
	fmt.Println(person)
	fmt.Println(person.Name, person.Age)

	person.X = 10
	fmt.Println(person.X)

	//结构体指针 如果我们有一个指向结构体的指针 p 那么可以通过 (*p).X 来访问其字段 X。 不过这么写太啰嗦了，所以语言也允许我们使用隐式解引用
	p := &person
	p.X = 12
	fmt.Println(person.X)

	fmt.Println(v1, v2, v3, vp)
}

type Person struct {
	X    int
	Y    int
	Name string
	Age  uint8
}
