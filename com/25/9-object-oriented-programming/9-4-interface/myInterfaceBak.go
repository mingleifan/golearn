package main

import "fmt"

type Panda struct {
	color string //猫的颜色
}

func (this *Panda) Sleep() {
	fmt.Println("Panda is Sleep...")
}

func (this *Panda) GetColor() string {
	return this.color
}

func (this *Panda) GetType() string {
	return "Panda"
}
