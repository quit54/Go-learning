package main

import "fmt"

//这个代码主要说明一下切片（动态数组）的4种声明方法
func main() {
	//第一种方法，切片变量直接声明
	mt := []int{1, 2}
	fmt.Println("%T", mt)
	//第二种方法，var声明法,可以声明全局，同时进行管控
	// var my[]int
	//记住，这里是声明了slice，但并没有给slice分配空间，所以不能使用
	// my = make([]int,3)
	// fmt.Println("%T",my)
	var my []int = make([]int, 3)
	fmt.Println("长度：", len(my), "名称：", my)

}
