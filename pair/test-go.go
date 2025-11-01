package main

import "fmt"

func main() {
	a := "小雨"
	//go语言对于变量，内置pair结构，分为两个部分：一个存储数据类型，另一个存储数据内容
	var ok interface{}
	ok = a
	value, _ := ok.(string)
	fmt.Println(value)
}
