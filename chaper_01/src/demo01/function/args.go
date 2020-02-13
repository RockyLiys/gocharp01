package function

import (
	"bytes"
	"fmt"
)

// 定义一个函数, 参数数量为0~n, 类型约束为字符串
func joinStrings(sList ...string) string {

	fmt.Println(sList, len(sList))
	// 定义一个字节缓冲, 快速地连接字符串
	var b bytes.Buffer
	// 遍历可变参数列表sList, 类型为[]string
	for _, s := range sList {
		// 将遍历出的字符串连续写入字节数组
		b.WriteString(s)
	}

	// 将连接好的字节数组转换为字符串并输出
	return b.String()
}

func JoinStr()  {
	fmt.Println(joinStrings("pig ", "and", " rat"))
	fmt.Println(joinStrings())
}

func printTypeValue(sList ...interface{}) string {
	// 字节缓冲作为快速字符串连接
	var b bytes.Buffer
	// 遍历参数
	for _, s := range sList {
		// 将interface{}类型格式化为字符串
		str := fmt.Sprintf("%v", s)
		// 类型的字符串描述
		var typeString string
		// 对s进行类型断言
		switch s.(type) {
		case bool:    // 当s为布尔类型时
			typeString = "bool"
		case string:    // 当s为字符串类型时
			typeString = "string"
		case int:    // 当s为整型类型时
			typeString = "int"
		}
		// 写字符串前缀
		b.WriteString("value: ")
		// 写入值
		b.WriteString(str)
		// 写类型前缀
		b.WriteString(" type: ")
		// 写类型字符串
		b.WriteString(typeString)
		// 写入换行符
		b.WriteString("\n")
	}
	return b.String()
}

func PrintTypeVal()  {
	fmt.Println(printTypeValue(100, "str", true))
}

// 实际打印的函数
func rawPrint(rawList ...interface{}) {

	// 遍历可变参数切片
	for _, a := range rawList {

		// 打印参数
		fmt.Println(a)
	}
}

// 打印函数封装
func print(sList ...interface{}) {

	// 将slist可变参数切片完整传递给下一个函数
	rawPrint("fmt", sList)
	rawPrint(sList...)
}

func Print()  {
	print(1, 2, 3)
}