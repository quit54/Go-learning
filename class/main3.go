package main

import "fmt"

func myfunc(a interface{}) {
	//接口函数允许接受任意类型的数据，同时添加类型断言
	fmt.Println("任意类型接口输出：")
	fmt.Println(a)
	fmt.Println("--------------------------------")

	value, ok := a.(string)
	//"类型断言"的机制
	if !ok {
		fmt.Println("当前接口的数据类型不是string\n")
	} else {
		fmt.Printf("当前的数据是string类型，内容为%s\n", value)
	}
}
func main() {
	a := "name"
	b := 123
	myfunc(a)
	myfunc(b)
}
