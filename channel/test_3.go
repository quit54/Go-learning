// 这个代码主要讲解channel的关闭机制
package main

import "fmt"

func main() {
	c := make(chan int)
	go func() {
		for i := 1; i < 6; i++ {
			c <- i
		}
		// close(c)
		//这里close（）语句的进行与否，主要查看c的传递数值仍然存在，如果不存在就会结束
		//如果不存在close（）这样的语句，会有死锁发生
	}()
	// for {
	// 	if a, ok := <-c; ok {
	// 		fmt.Println(a)
	// 	} else {
	// 		break
	// 	}
	// }
	for data := range c {
		fmt.Println(data)
	}
	fmt.Println("主进程结束。")
}
