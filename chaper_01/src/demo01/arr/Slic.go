package arr

import "fmt"

func Sli(s [4]int)  {
	fmt.Println(s[:])
	fmt.Println(s[1:3])
	fmt.Println(s[0:0])
	fmt.Println(s[:])
}
