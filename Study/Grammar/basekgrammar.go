//变量、语法
package basekgrammar

import (
	"fmt"
)

//变量赋值
func BasedVriable() {
	/*
		基础语法：
			值类型 ： int string float double 通过 &i 来获取变量 i 的内存地址 值类型的变量的值存储在栈中
			引用类型：一个引用类型的变量 r1 存储的是 r1 的值所在的内存地址（数字），或内存地址中第一个字所在的位置。
				这个内存地址称之为指针，这个指针实际上也被存在另外的某一个值中。
					当使用赋值语句 r2 = r1 时，只有引用（地址）被复制。
					如果 r1 的值被改变了，那么这个值的所有引用都会指向被修改后的内容，在这个例子中，r2 也会受到影响。
	*/
	var str string = "第一种写法"
	var str2 = "第二种写法"
	str3 := "第三种写法" //初始化声明 高效的创建变量

	str4, str5, str6 := "多变量声明1", "多变量声明2", "多变量声明3"

	fmt.Println(str)
	fmt.Println(str2)
	fmt.Println(str3)
	fmt.Println(str4, str5, str6)
}

func init() {

	fmt.Println("init basekgrammar...")
}
