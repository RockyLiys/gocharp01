package structs

import "fmt"

type Cat struct {
	Color string
	Name string
}

type BlackCat struct {
	Cat
}

func newCat(color, name string) *Cat  {
	return &Cat{
		Color: color,
		Name:  name,
	}
}

func newBlackCat(cat Cat) *BlackCat  {
	blackCat := &BlackCat{}
	cat.Color = "Black"
	blackCat.Cat = cat
	return blackCat
}

func Input()  {
	cat := newCat("Yellow", "Hello")
	block := newBlackCat(*cat)
	fmt.Println(cat)
	fmt.Println(block)
}
