package main

import (
	"fmt"
)

type animal interface {
	Sleep()
	Eat()
	Walk()
}
type Dog struct {
	name string
}
type Cat struct {
	name string
}

func (this *Dog, one *Cat) Sleep() {
	fmt.Println("小狗在睡觉")
}
func (this *Dog) Walk() {
	fmt.Println("小狗在走路")
}
func (this *Dog) Eat() {
	fmt.Printf("%s小狗在吃饭\n", this.name)
}
func show(a animal) {
	a.Sleep()
	a.Eat()
	a.Walk()
} //这就是一种多态
func main() {
	var a animal
	a = &Dog{"tian"}
	show(a)

}
