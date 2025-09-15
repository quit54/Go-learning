package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	// Set properties of the predefined Logger, including，设置预定义的Logger属性，包括日志条目前缀和禁用打印的标志
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// Request a greeting message.
	message, err := greetings.Hello("樊浩成") //这里调用模块的时候，返回了两个价值，一个字符串一个错误
	// If an error was returned, print it to the console and
	// exit the program.
	if err != nil {
		log.Fatal(err)
	} //可对错误进行一个条件判定，没问题就输出，有问题就更新训练日志

	// If no error was returned, print the returned message
	// to the console.
	fmt.Println(message)
}
