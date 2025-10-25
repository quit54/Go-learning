package main

import "fmt"

type human struct {
	name string
	sex  string
} //human作为父类，写出子类作为父类的继承
type ren struct {
	human
	height int
} //ren的类这样就继承了human类的方法和数据
func (this *human) Eat() {
	fmt.Println("human eat()...")
}
func (wok *human) Walk() {
	fmt.Println("human walk()...")
}
func (this *ren) Ok() {
	fmt.Println("你的身高是多少?", this.height)
}

//核心思想：这些human类函数首字母都大写，这样才可以在定义中被调用
func main() {
	h := human{"xiaoyu", "female"}
	h.Eat()
	h.Walk()
	h1 := ren{human{"xiaoyu", "female"}, 180}
	h1.Eat()
	h1.Walk()
	h1.Ok()
}
