package function

import "fmt"

func fire()  {
	fmt.Println("fire...")
}

// 函数也是一个类型
func FunAsParams() func() {
	var f func()
	f = fire
	return f
}

