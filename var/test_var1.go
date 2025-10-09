package main

/*
	四种变量的声明方式
*/
import "fmt"

//声明全局变量，方法一，方法二，方法三是可以的
var ga int = 100
var gb = 1.00
var gc = "yu"

//用方法四声明全局变量（不可行），gC := 123,方法四变量声明只能声明在函数体内

func main() {
	var a int //方式一：声明一个变量，默认值为0
	fmt.Println("a = ", a)

	var b int = 100 //方式二：声明一个变量，赋予初始化值
	fmt.Println("b = ", b)
	var bb string = "ABCD"
	fmt.Printf("bb = %s\n", bb)

	var c = 1.0 //方式三：在初始化的时候省略类型名，通过数值自动匹配当前变量类型
	fmt.Println("c = ", c)
	fmt.Printf("c = %T\n", c) //这里%T的意思是调用变量名

	//方式四：(最常用的方法)省去var关键词，直接自动匹配，生成一个变量
	e := 100
	fmt.Println("e = ", e)
	fmt.Printf("type of f = %T\n", e)

	fmt.Println("ga = ", ga, "gb = ", gb, "gc = ", gc)
	fmt.Printf("type of f = %T\ntype of f = %T\ntype of f = %T\n", ga, gb, gc)

	//声明多个变量
	aa, kk := 100, 12
	//等效于var aa,kk int = 100,12
	fmt.Println(aa, kk)

	//多行变量声明
	var (
		x int    = 1
		y string = "bt"
	)

}
