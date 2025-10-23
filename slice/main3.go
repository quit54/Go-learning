package main

import "fmt"

func main() {
	var ze []int = make([]int, 1, 1)
	fmt.Printf("len=%d  cap=%d name=%v\n", len(ze), cap(ze), ze)
	//cap对于切片来说是容量的意思，容量意味着代码底层为切片留存的空间大小，是可以改变的
	//因此切片的数量多少不允许超过它的大小，超过大小类似于数组越界

	//对于切片内，已分配好空间的扩充和衰减
	//比如想要增加数据，使用append指令为切片增加，但是cap的数值不变，还是一样的
	ze = append(ze, 1)
	fmt.Printf("len=%d  cap=%d name=%v\n", len(ze), cap(ze), ze)
	ze = append(ze, 1)
	//向一个容量为1，并且已经被装满的切片传递一个数值1，这个时候切片会自动对自己的内存进行扩容
	//cap数组始终在动态的收缩和扩张，依据数据规模，所以绝大多数时候，如果不强制规定容量，就会是这样
	fmt.Printf("len=%d  cap=%d name=%v\n", len(ze), cap(ze), ze)

	// 	len=1  cap=1 name=[0]
	// len=2  cap=2 name=[0 1]
	// len=3  cap=4 name=[0 1 1]
	//这里注意观察cap的内存扩充原理，每一次扩充按照当前的内存大小乘以2
}

//这里引发我对于代码的思考：动态数组（切片）与静态数组的区别
//很明显，动态数组的内存是动态调整的，可以在每次使用的时候进行分配内存和删除内存
