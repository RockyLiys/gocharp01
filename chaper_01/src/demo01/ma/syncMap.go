package ma

import (
	"fmt"
	"sync"
)

func SyMap()  {
	var sm sync.Map
	sm.Store(111, "aaaaa")
	sm.Store(222, "bbbbb")
	sm.Store("dddd", "ffff")

	sma, o := sm.LoadOrStore("hhh", 88888)
	fmt.Println(o)
	if o{
		fmt.Println(sma)
	}
	val, ok := sm.Load("dddd")
	if ok{
		fmt.Println(val)
	}
	sm.Range(func(k, v interface{}) bool{
		fmt.Printf("key=%v, value=%v\n", k, v)
		return true
	})

}
