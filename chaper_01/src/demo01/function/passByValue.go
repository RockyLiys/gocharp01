package function

import "fmt"

// 测试值传递效果的结构体
type Data struct {
	Complax  []int      // 测试切片在参数传递中的效果
	Instance InnerData  // 实例分配的innerData
	Ptr      *InnerData // 将ptr 声明为InnerData的指针类型
}

type InnerData struct {
	A int
}

// 传值测试参数
func PassByValue(inFunc Data) Data  {

	inFunc.Instance.A = 6

	// 输出参数的成员情况
	fmt.Printf("inFunc value %+v\n", inFunc)

	// 打印inFunc的指针
	fmt.Printf("inFunc ptr: %p\n", &inFunc)



	return inFunc
}
