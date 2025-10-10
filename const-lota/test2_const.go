package main

import "fmt"

//首先const属性和C语言一样，是只读属性
//其次，iota的意思是按照顺序叠加,与const表示枚举类型，默认值为零，被应用在const里面
//iota是一种叠加的数据概念，在不断+1改变
const (
	BEIJING  = iota //iota = 0
	SHANGHAI        //iota = 1
	SHENZHEN        //iota = 2
)

//这里的常量是不允许修改的
const (
	a, b = iota + 1, iota + 2
	c, d
	e, f
	//1,2   2,3   3,4

	g, k = iota * 1, iota * 2
	//iota=3，g=3,k=6
	m, n
	//iota=4，m=4,n=8
)

func main() {
	fmt.Println(BEIJING, SHANGHAI, SHENZHEN)
	fmt.Println(a, b, c, d, e, f, g, k, m, n)
}
