package structs

import "fmt"

type Command struct {
	Name string
	Var *int
	Commend string
}

var version int = 1

func initStruct() *Command  {
	var com Command
	com.Commend = "version"
	com.Var = &version
	com.Name = "show version"

	return &com
}

func initStruct1() *Command  {
	return &Command{
		Name:    "show version",
		Var:     &version,
		Commend: "version",
	}
}

func initStruct2() *Command  {
	com := new(Command)
	com.Commend = "version"
	com.Var = &version
	com.Name = "show version"

	return com
}

func InitStruct()  {
	 c := initStruct()
	 fmt.Println(c.Name, c.Var, c.Commend)
	 c1 := initStruct1()
	 fmt.Println(c1.Commend, c1.Var, c1.Name)
	 c2 := initStruct2()
	 fmt.Println(c2.Name, c2.Var, c2.Commend)
}