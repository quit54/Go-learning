package main

import (
	"chang" // 我这里是在go语言本身文件src里面导入我的自编辑包，chang，所以可以不假思索地引用
	//这其实相当于没有解决我的问题，是一种拐弯的，取巧的办法。
	"fmt"
)

func main() {
	fmt.Println("Hello, World!")
	chang.FWC() // 调用 chang 包中的 FWC 函数
}

//但是如今遗留两个问题：第一：GOROOT和GOPARTH的路径匹配问题，第二：module在GO语言属于什么意思？它的创建与调用
