package function

import "fmt"

//提供一个值, 每次调用函数会指定对值进行累加

func accumulate(value int) func() int {
	// 返回一个闭包
	return func() int {
		value++
		//返回一个累加值
		return value
	}
}

func Accumulate()  {
	// 创建一个累加器, 初始值为1
	accumulator := accumulate(1)

	// 累加1并打印
	fmt.Println(accumulator())

	fmt.Println(accumulator())

	// 打印累加器的函数地址
	fmt.Printf("%p\n", &accumulator)

	//创建一个累加器, 初始化为1
	accumulator2 := accumulate(1)
	//累加1并打印
	fmt.Println(accumulator2())

	// 打印累加器的函数地址
	fmt.Printf("%p\n", &accumulator2)


}


// 创建一个玩家生成器, 输入名称, 输出生成器
func playerGen(name string) func() (string, int) {

	// 血量一直为150
	hp := 150

	// 返回创建的闭包
	return func() (string, int) {
		hp = 400
		// 将变量引用到闭包中
		return name, hp
	}
}

func PlayerGenerator()  {
	// 创建一个玩家生成器
	generator := playerGen("high noon")

	// 返回玩家的名字和血量
	name, hp := generator()

	// 打印值
	fmt.Println(name, hp)
}