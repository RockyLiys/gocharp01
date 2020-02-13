package function

import (
	"fmt"
	"time"
)

func test()  {
	start := time.Now()
	sum := 0
	for i := 0 ; i < 10000000 ; i++  {
		sum++
	}

	elapsed := time.Since(start)
	fmt.Println("该函数执行 完成耗时:", elapsed)

}

func test1()  {
	start := time.Now()
	sum := 0
	for i := 0 ; i < 10000000 ; i++  {
		sum++
	}

	elapsed1 := time.Now().Sub(start)
	fmt.Println("该函数执行 完成耗时:", elapsed1)

}

func Time()  {
	test()
	test1()
}
