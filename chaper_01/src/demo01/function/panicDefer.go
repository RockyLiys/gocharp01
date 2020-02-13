package function

import "fmt"

func PanicDefer()  {

	defer fmt.Println("宕机后要做的事情1")
	defer fmt.Println("宕机后要做的事情2")

	panic("宕机")
}






