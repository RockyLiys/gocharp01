package function

import (
	"fmt"
	"strings"
)

func stringProcess(list []string, chain []func(string) string)  {
	for index,  str := range list {
		result := str

		for _, proc := range chain{

			result = proc(result)
		}
		list[index] = result
	}
}

// 自定义的移除前缀的处理函数
func removePrefix(str string) string  {
	return strings.TrimPrefix(str, "go")
}

func StringProcess()  {
	// 待处理的字符串列表
	list := []string{
		"go scanner",
		"go parser",
		"go compiler",
		"go printer",
		"go formater",
	}
	fmt.Println(list, len(list))

	// 处理函数链
	chain := []func(string) string{
		removePrefix,
		strings.TrimSpace,
		strings.ToUpper,
	}

	fmt.Println(chain, len(chain))
	// 处理字符串
	stringProcess(list, chain)

	// 输出处理的字符串
	for _, str := range list{
		fmt.Println(str)
	}
}