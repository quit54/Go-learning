package main

import "fmt"

func main() {
	var ma map[string]string
	ma = make(map[string]string, 10)
	fmt.Printf("%d\n", len(ma))

	fmt.Println("---------问世间情为何物？直教人生死相许---------")
	ma["one"] = "lv"
	ma["two"] = "xin"
	ma["three"] = "yu"
	for idx, value := range ma {
		fmt.Println(idx, ":", value)
	}
	fmt.Println(ma)
	fmt.Printf("%d\n", len(ma))
	delete(ma, "one")
	fmt.Println(ma)
}

//ma的map数据类型，在函数之间就是一个指针传递
