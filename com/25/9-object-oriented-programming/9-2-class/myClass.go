package main

import "fmt"

// Hero 类名首字母大写，其他包也能够访问
type Hero struct {
	//如果累的属性首字母大写，其他包能够访问，否则仅本包内访问
	Name  string
	Ad    int
	Level int
}

func (this *Hero) GetName() string {
	fmt.Println("Name =", this.Name)
	return this.Name
}

func (this *Hero) SetName(name string) {
	this.Name = name
}

func main() {

	hero1 := Hero{
		Name:  "RanGo",
		Ad:    30,
		Level: 1,
	}
	fmt.Println(hero1)

	hero1.GetName()

	hero1.SetName("RanGo2")

	fmt.Println(hero1)
}
