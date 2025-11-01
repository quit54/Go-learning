package main

import (
	"fmt"
	"time"
) //有缓冲与无缓冲channel之间存在区别：管道有没有进行channel的存储，channel具备存储的话，进程很少出现卡断，两边正常进行
func main() {
	c := make(chan int, 4)
	fmt.Println("main goroutine进程开始")
	fmt.Println("len() = ", len(c), "cap() = ", cap(c))
	go func() {
		defer fmt.Println("子进程结束")

		for i := 0; i < 3; i++ {
			c <- i
			fmt.Println("子进程输出数据 = ", i, "len() = ", len(c), "cap() = ", cap(c))
		} //len的长度，因为这在两个进程的参数传递管道中，还剩有多少余量。
	}()
	time.Sleep(2 * time.Second)
	for i := 0; i < 3; i++ {
		num := <-c
		fmt.Println("num = ", num)
	}
	fmt.Println("主进程结束")
}
