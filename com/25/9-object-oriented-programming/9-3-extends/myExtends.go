package main

import "fmt"

type Human struct {
	name string
	sex  string
}

func (this *Human) Eat() {
	fmt.Println("Human.Eat()...")
}

func (this *Human) Walk() {
	fmt.Println("Human.Walk()...")
}

//===========================

type SuperMan struct {
	Human //superMan类基础了Human类的方法

	level int
}

// 重定义父类的方法
func (this *SuperMan) Eat() {
	fmt.Println("SuperMan.Eat()...")
}

// 子类的新方法
func (this *SuperMan) Fly() {
	fmt.Println("SuperMan.Fly()...")
}

func (this *SuperMan) Print() {
	fmt.Println("name =", this.name)
	fmt.Println("sex =", this.sex)
	fmt.Println("level =", this.level)
}

func main() {

	h := Human{"Tom", "Male"}

	h.Eat()
	h.Walk()

	s := SuperMan{Human{"Tom", "Male"}, 100}
	s.Eat()
	s.Walk()
	s.Fly()

	var s1 SuperMan
	s1.name = "Tom"
	s1.sex = "Male"
	s1.level = 100

	s1.Print()

}
