package ma

import "fmt"
import "strconv"

func Mp(m map[int]string) map[string]int  {
	m[3] = "ddd"
	for key, value := range m{
		fmt.Printf("key=%d, value=%s\n", key, value)
	}
	mml := make(map[string]int)
	for k,v := range m{
		mml[v] = k
	}
	return mml
}

func MultiMp() map[int][]string {
	//var iStringList  map[int][]string
	iStringList := make(map[int][]string)
	for index := 1; index <= 5 ; index++  {
		ss := []string{"aa", "bb", "cc"}
		iStringList[index] = ss
	}
	return iStringList
}

// 每一次遍历后key都会变量，而不是有序的
// 遍历输出元素的顺序与填充顺序无关，不能期望 map 在遍历时返回某种期望顺序的结果
func MultiDict() map[int]map[string]string  {
	var intStrDict map[int]map[string]string
	intStrDict = make(map[int]map[string]string)
	for i := 1; i<=4; i++{
		dict := make(map[string]string)
		for in:=1;in<=3 ;in++  {
			dict[strconv.Itoa(in)] = fmt.Sprintf("index=%d",in)
		}
		intStrDict[i] = dict
	}
	return intStrDict
}
