package main

import "fmt"

var myArray = []int{1, 2, 3}

func p(myArray []int) {

	for _, value := range myArray {
		fmt.Println("输出数值：", value)
	}
}
func main() {
	p(myArray)
}

//切记，数组在函数中的传递是值拷贝，切片在函数中数值的传递是引用传递拷贝，动态数组在传参上是引用传递的
