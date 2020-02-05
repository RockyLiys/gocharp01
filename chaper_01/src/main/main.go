package main

import (
	"demo01/arr"
	"demo01/func_demo"
	"demo01/image"
	"demo01/ma"
	"demo01/ptr"
	"demo01/str"
	"fmt"
	"global"
	"httpweb"
	"math/rand"
	"os"
	"sort"
	"time"
)

// 数据生产者
func producer(header string, channel chan<- string) {
	// 无限循环, 不停地生产数据
	for {
		// 将随机数和字符串格式化为字符串发送给通道
		channel <- fmt.Sprintf("%s: %v", header, rand.Int31())
		// 等待1秒
		time.Sleep(time.Second)
	}
}

// 数据消费者
func customer(channel <-chan string) {
	// 不停地获取数据
	for {
		// 从通道中取出数据, 此处会阻塞直到信道中返回数据
		message := <-channel
		// 打印数据
		fmt.Printf("receive message = '%s'\n", message)
	}
}

func goroutines() {
	// 创建一个字符串类型的通道
	channel := make(chan string)
	// 创建producer()函数的并发goroutine
	go producer("send cat ", channel)
	go producer("send dog ", channel)
	// 数据消费函数
	customer(channel)
}

func Ftp() {
	httpweb.Webftp()
}

func Image() {
	image.OutImage()
}

func Strin() {
	var ss string = "你好！ 中国，我爱你！！"
	retStr := str.Stri(ss)
	// 把字符串改为字符串数组
	angleBytes := []byte(retStr)
	fmt.Println(angleBytes)
	fmt.Println(retStr)
	fmt.Println(str.MutiStr())
}

func P() {
	var c string = "hello,world!!"
	var s string
	//var ps *string
	fmt.Println(&c)
	ps := ptr.Ptr(&c)
	fmt.Printf("%p\n", &ps)
	fmt.Println(*ps)
	s = fmt.Sprintf("%T", *ps)
	fmt.Println(s)
}
func Sw() {
	x, y := 1, 2
	ptr.Swap(&x, &y)
	fmt.Println(x, y)
}

func Arra() {
	var ar = [6]int{1, 7, 3, 4, 5, 6}
	q := [...]int{1, 2, 3, 4}
	fmt.Println(q)
	arr.Array(ar)
	var ma [2][2]string = [2][2]string{{"aa", "bb"}, {"cc", "dd"}}
	arr.MutiArr(ma)
}

func Sl() {
	q := [...]int{1, 2, 3, 4}
	arr.Sli(q) // Slice 切片
	var s []int
	s = append(s, 11, 22, 33)
	fmt.Println(s)
	var ss []string
	ss = append(ss, "fff", "ggg", "dddd")
	fmt.Println(ss)
}

func Mp()  {
	var maL = map[int]string{1:"aaa", 2:"bbbb"}
	mp := ma.Mp(maL)
	mp["ccc"] = 4
	for k,v :=range mp{
		fmt.Printf("key=%s, value=%d\n", k,v)
	}
	intListString := ma.MultiMp()
	for ke, va := range intListString{
		for _, val := range va{
			fmt.Printf("key= %d, val=%s\n",ke, val)
		}
	}
	// 遍历输出元素的顺序与填充顺序无关，不能期望 map 在遍历时返回某种期望顺序的结果
	fmt.Println("-----------------")
	intDictString := ma.MultiDict()
	var sli []int
	for index, dict :=range intDictString{
		for ky, v := range dict{
			fmt.Printf("第%d行, key=%s, value=%s\n", index, ky, v)
		}
		sli = append(sli, index)
	}
	fmt.Println(sli)
	// 对map生成有序的key
	var slic []int
	for k := range intDictString{
		slic = append(slic, k)
	}
	sort.Ints(slic)

	fmt.Println(slic)
	for _, index := range slic{
		fmt.Println(index)
		for ky, v := range intDictString[index]{
			fmt.Printf("第%d行, key=%s, value=%s\n", index, ky, v)
		}
	}
}
func List()  {
	ls := ma.Lis()
	fmt.Printf("============== %d\n", ls.Len())
	for i:=ls.Front();i!=nil ;i=i.Next()  {
		fmt.Println(i.Value)
	}
	fmt.Printf("============== %d\n", ls.Len())
	fmt.Println(ls.Back().Value)
	fmt.Println(ls.Front().Value)
	fmt.Println(ls.Back().Prev().Value)

	ls.MoveToBack(ls.Front())
	ls.MoveToFront(ls.Back())

	fmt.Printf("============== %d\n", ls.Len())
	for i:=ls.Front();i!=nil ;i=i.Next()  {
		fmt.Println(i.Value)
	}
}

func SM()  {
	ma.SyMap()
}

func FuncDemo()  {
	// 直接打印出返回值
	fmt.Println(func_demo.ResolveTime(1000))

	// 只打印小时与分钟
	_, hour, minute := func_demo.ResolveTime(18000)
	fmt.Println(hour, minute)

	// 内要天
	day, _, _ := func_demo.ResolveTime(90000)

	fmt.Println(day)
}

func PByVal()  {
	// 准备传入函数的结构
	in := func_demo.Data{
		Complax: []int{1, 2, 3},
		Instance: func_demo.InnerData{
			A: 5,
		},

		Ptr: &func_demo.InnerData{
			A: 1,
		},
	}

	// 输入结构的成员情况
	fmt.Printf("输入初始化结构 in value: %+v\n", in)

	// 输入结构的指针地址
	fmt.Printf("in ptr: %p\n", &in)

	// 传入结构体, 返回同类型的结构
	out := func_demo.PassByValue(in)

	// 输出结构的成员情况
	fmt.Printf("out value: %+v\n", out)

	// 输出结构的指针地址
	fmt.Printf("out ptr: %p\n", &out)

}

func main() {
	// 相对app位置
	appProPath, err := os.Getwd()
	global.AppProPath = appProPath
	if err == nil {
		fmt.Printf("app 相对位置: %s\n", appProPath)
	}
	//Ftp() // 简单实现ftp服务
	//Image() // 实现图像输出
	//Strin()  // 字符串练习
	//P()     // 指针
	//Sw()     // 指针交换
	//Arra()     //数组
	//Sl() //切片
	//Mp()     // map
	//List()    // list
	//SM()      // sm
	//FuncDemo()  // 函数操作,形参,实参,返回值,与接收值的练习
	PByVal()
}
