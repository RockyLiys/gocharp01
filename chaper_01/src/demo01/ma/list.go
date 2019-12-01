package ma

import (
	"container/list"
	"fmt"
)

func Lis() list.List {
	var l list.List
	li := list.New()
	l.PushFront("aaa")
	l.PushFront("bbbb")

	e := l.PushBack("cccc")
	l.InsertAfter("gggg", e)
	l.InsertBefore("hhh", e)

	li.PushBack("dddd")
	el := li.PushFront("fff")
	li.InsertBefore(111, el)
	li.InsertAfter(2222, el)

	for i:=l.Front(); i!=nil;i=i.Next(){
		fmt.Println(i.Value)
	}
	fmt.Println("------------")
	for i:=li.Front();i!=nil ;i=i.Next()  {
		fmt.Println(i.Value)
	}
	l.PushBackList(li)
	fmt.Println("------------")

	for i:=l.Front();i!=nil ;i=i.Next()  {
		fmt.Println(i.Value)
	}
	return l
}
