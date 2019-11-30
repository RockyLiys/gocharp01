package ptr

import "fmt"

func Ptr(s *string) *string {
	p := s
	*p = *p + "我不hello"
	fmt.Println(p,*s,s)
	return p
}

func Swap(a, b *int)  {
	*a, *b = *b, *a
}
