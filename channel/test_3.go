// 这个代码主要讲解channel的关闭机制
package main

import "fmt"

func main() {
	c := make(chan int)
	go func() {
		for i := 1; i < 6; i++ {
			c <- i
		}
		close(c)
	}()
	for {
		if a, ok := <-c; ok {
			fmt.Println(a)
		} else {
			break
		}
	}
	fmt.Println("主进程结束。")
}
