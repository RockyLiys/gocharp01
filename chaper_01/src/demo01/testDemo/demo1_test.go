package testDemo

import (
	"fmt"
	"testing"
)

func TestGetArea(t *testing.T) {
	area := GetArea(40, 50)

	if area != 2000{
		t.Error("测试失败!")
	}
}

func BenchmarkGetArea(b *testing.B) {
	fmt.Println(b.N)
	for i := 0; i < b.N ; i++  {
		GetArea(40, 50)
	}
}