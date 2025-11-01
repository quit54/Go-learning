package main

import (
	"fmt"
)

type animal interface {
	Sleep()
	Eat()
	Walk()
	//Show(),在接口中调用这个普通函数是一种俄罗斯套娃的逻辑，很不正确
} //多态的前提是一定要有父类的接口，
// 其次全部实现了子类的接口，有指针指向了父类具体的变量
type Dog struct {
	name string
}
type Cat struct {
	name string
}

func (this *Dog) Sleep() {
	fmt.Println("小狗在睡觉")
}

// 在 Go 语言中,这种语法定义了一个方法（method），而不是普通的函数（function）。
// 这里的 this *Dog 是方法的接收者（receiver），它指定了这个方法是属于哪个类型，以及如何访问该类型的实例。
func (this *Dog) Walk() {
	fmt.Println("小狗在走路")
}
func (this *Dog) Eat() {
	fmt.Printf("%s小狗在吃饭\n", this.name)
}
func (this *Cat) Walk() {
	fmt.Println("小猫在走路")
}
func (this *Cat) Eat() {
	fmt.Printf("%s小猫在吃饭\n", this.name)
}
func (this *Cat) Sleep() {
	fmt.Println("小猫在睡觉")
}
func Show(a animal) {
	a.Sleep()
	a.Eat()
	a.Walk()
} //这就是一种多态
func main() {
	var a animal
	a = &Dog{"tian"}
	Show(a)

	a = &Cat{"yu"}
	Show(a)
	fmt.Printf("-----------------------------------")

}
