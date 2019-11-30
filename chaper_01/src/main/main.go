package main

import (
	"demo01/arr"
	"demo01/image"
	"demo01/ptr"
	"demo01/str"
	"fmt"
	"global"
	"httpweb"
	"math/rand"
	"os"
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

func Strin()  {
	var ss string="你好！ 中国，我爱你！！"
	retStr := str.Stri(ss)
	// 把字符串改为字符串数组
	angleBytes := []byte(retStr)
	fmt.Println(angleBytes)
	fmt.Println(retStr)
	fmt.Println(str.MutiStr())
}

func P()  {
	var c string="hello,world!!"
	var s string
	//var ps *string
	fmt.Println(&c)
	ps := ptr.Ptr(&c)
	fmt.Printf("%p\n", &ps)
	fmt.Println(*ps)
	s = fmt.Sprintf("%T", *ps)
	fmt.Println(s)
}
func Sw()  {
	x, y := 1, 2
	ptr.Swap(&x, &y)
	fmt.Println(x,y)
}

func Arra()  {
	var ar = [6]int{1, 7, 3, 4, 5, 6}
	q := [...]int{1, 2, 3, 4}
	fmt.Println(q)
	arr.Array(ar)
	var ma [2][2]string = [2][2]string{{"aa","bb"},{"cc","dd"}}
	arr.MutiArr(ma)
}

func Sl()  {
	q := [...]int{1, 2, 3, 4}
	arr.Sli(q)   // Slice 切片
	var s []int
	s = append(s, 11, 22, 33)
	fmt.Println(s)
	var ss []string
	ss = append(ss, "fff", "ggg", "dddd")
	fmt.Println(ss)
}

func main() {
	// 相对app位置
	appProPath, err := os.Getwd()
	global.AppProPath = appProPath
	if err == nil {
		fmt.Printf("app 相对位置: %s\n", appProPath)
	}
	Ftp() // 简单实现ftp服务
	//Image() // 实现图像输出
	//Strin()  // 字符串练习
	//P()     // 指针
	//Sw()     // 指针交换
	//Arra()     //数组
	//Sl()		//切片
}