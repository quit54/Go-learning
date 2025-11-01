package main

import "fmt"

type Reader interface {
	Readbook()
}
type Writer interface {
	Writebook()
}
type Book struct {
}

func (this *Book) Readbook() {
	fmt.Println("在读书")
}
func (this *Book) Writebook() {
	fmt.Println("在写书")
}
func main() {
	a := &Book{}
	//a存储的是一个地址

	//在命名这个b的变量时，b的pair<>内容为空，并且Reader的接口决定它可以接受任何数据类型
	var b Reader
	//这个B的pair存储的是pair<type:Book ,value:book{}地址>
	b = a
	b.Readbook() //接收地址完毕后，可以通过接口地址来访问接口内部的方法
	//以下同理
	fmt.Println("--------------------")

	var c Writer
	c = a
	c.Writebook()
}
