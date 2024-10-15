package main

import "fmt"

func main() {
	person := Person{1, 2, "rango", 26}
	fmt.Println(person)
	fmt.Println(person.Name, person.Age)
}

type Person struct {
	X    int
	Y    int
	Name string
	Age  uint8
}
