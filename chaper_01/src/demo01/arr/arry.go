package arr

import "fmt"

func Array(ar [6]int)  {
	for i, v := range ar {
		fmt.Printf("index=%d, value=%d\n", i, v)
	}
}

func MutiArr(a [2][2]string)  {
	for i, v1 := range a {
		i++
		for _, v2 := range v1{
			fmt.Printf("第%d行，value=%s\n",i, v2)
		}
	}
}
