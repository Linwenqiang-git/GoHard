package compilestudy

import "fmt"

type coder interface {
	code()
	debug()
}
type Gopher struct {
	language string
}

func (p Gopher) code() {
	fmt.Printf("I am coding %s language\n", p.language)
}
func (p Gopher) debug() {
	fmt.Printf("I am debuging %s language\n", p.language)
}

/*
打印出汇编语言
go tool compile -S main.go
*/
func Out() {
	x := 200
	var any interface{} = x
	fmt.Println(any)
	g := Gopher{"Go"}
	var c coder = g
	fmt.Println(c)
}
