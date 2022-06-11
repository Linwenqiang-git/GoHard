package compilestudy

import "fmt"

type Person interface {
	growUp()
}
type Student struct {
	age int
}

func (p Student) growUp() {
	p.age += 1
	return
}
func Struct() {
	var qcrao = Person(Student{age: 18})
	fmt.Println(qcrao)
}
