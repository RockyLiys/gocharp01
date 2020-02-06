package function

import (
	"flag"
	"fmt"
	"reflect"
)

var skillParam = flag.String("skill", "", "skill to perform")


func FunLambda() {
	// 是做命令行输入参数解析
	flag.Parse()
	var skill = map[string]func(){
		"fire": func() {
			fmt.Println("chicken fire")
		},
		"run": func() {
			fmt.Println("soldier run")
		},
		"fly": func() {
			fmt.Println("angel fly")
		},
	}
	if f, ok := skill[*skillParam]; ok {
		fmt.Println(reflect.TypeOf(f))
		fmt.Println()
		if *skillParam == "run" {
			fmt.Println(*skillParam)
			f()
		}else {
			f()
		}
	} else {
		fmt.Println("skill not found")
	}
}

