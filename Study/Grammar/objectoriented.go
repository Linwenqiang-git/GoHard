package objectoriented

import (
	"fmt"
)

/*
面向对象的封装
1.首字母大写表示类或方法对外是public 小写为 private
*/
type Book struct {
	Name   string
	Author string
	Price  float32
}

//(this *Book) 表示该方法属于类的内部 *表示指针传递，不然在修改值时，只是一个拷贝，值不会修改
func (this *Book) ChangePrice(price float32) {
	this.Price = price
}
func (this Book) GetBookName() string {
	return this.Name
}

/*继承*/
type MathBook struct {
	Book //类里显示声明父类
	Page int
}

func (this *MathBook) ShowPage() {
	fmt.Print("一共有 %d页", this.Page)
}

/*接口多态*/
type Animal interface {
	Sleep(hourse int)
}

//不需要显示的继承接口，只要把接口的方法实现就OK
type Cat struct {
	Name string
}

func (this *Cat) Sleep(hourse int) {
	fmt.Println("cat sleep %d hourse", hourse)
}

func init() {
	fmt.Println("object oriented init...")
	//对象实例化
	book := Book{Name: "zhangsan"}
	fmt.Println(book.GetBookName())
	//子类实例化
	mathbook := MathBook{Book{Name: "数学", Author: "笛卡尔", Page: 230}}
	mathbook.ShowPage()
	//继承多态
	animal Animal = &Cat{Name:"cat"}
}
