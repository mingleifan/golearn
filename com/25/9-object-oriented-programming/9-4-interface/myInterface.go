package main

import "fmt"

// Animal 本质是一个指针
type Animal interface {
	Sleep()

	GetColor() string // 获取动物的颜色
	GetType() string  // 获取动物的种类
}

type Cat struct {
	color string //猫的颜色
}

func (this *Cat) Sleep() {
	fmt.Println("Cat is Sleep...")
}

func (this *Cat) GetColor() string {
	return this.color
}

func (this *Cat) GetType() string {
	return "Cat"
}

type Dog struct {
	color string
}

func (this *Dog) Sleep() {
	fmt.Println("Dog is Sleep...")
}

func (this *Dog) GetColor() string {
	return this.color
}

func (this *Dog) GetType() string {
	return "Dog"
}

func showAnimal(animal Animal) {
	animal.Sleep()
	fmt.Println("color =", animal.GetColor())
	fmt.Println("type =", animal.GetType())
}

func main() {

	var animal Animal //接口的数据类型，父类指针
	animal = &Cat{"Green"}
	animal.Sleep()

	animal = &Dog{"Yellow"}
	animal.Sleep()

	fmt.Println("--------------")

	cat := Cat{"new Green"}
	dog := Dog{"new Yellow"}
	panda := Panda{"white and Black"}
	showAnimal(&cat)
	showAnimal(&dog)
	showAnimal(&panda)
}
