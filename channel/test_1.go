package main

import "fmt"

func main() {
	c := make(chan int)
	fmt.Println("main goroutine主要进程开始")
	go func() {
		defer fmt.Println("goroutinue结束")
		fmt.Println("goroutinue进行中~~~~")
		c <- 666
	}()
	//channel本身具有同步两个协程的能力
	//main.go和sub.go两个管道同步进行，当main.go进行中，遇见需要在两个管道之间传递参数的时候，它会停止，等待sub.go执行完。
	//直到传递完毕参数，才会继续运行

	num := <-c
	fmt.Println("输出", num)
	fmt.Println("main goroutine主要进程结束")

}
