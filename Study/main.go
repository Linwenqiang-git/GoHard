package main

import (
	//. "github.linwenqiang.com/Study/UnSafe"
	"fmt"
	//t "github.linwenqiang.com/CompileStudy"
	//_ 表示会走包的init方法，下面不调用也OK
	// . 表示调用内部公开方法时可不用写包名
	//. "/DataLogic/dataLogic"
	//_ "github.linwenqiang.com/Study/DataLogic"
	//. "github.linwenqiang.com/Study/SmallTools"
	//DL "/DataLogic/dataLogic" //DL 给包起的别名
	//_ "github.linwenqiang.com/Study/Grammar"
	//. "github.linwenqiang.com/Study/Test"
)

//声明全局变量
// var (
// 	a int
// 	b bool
// )

// func outOfRange() int {
// 	arr := [3]int{1, 2, 3}
// 	i := 4
// 	elem := arr[i]
// 	return elem
// }

// var BasePath string = "C:\\work5.0\\xbot-v3\\src\\client\\"

func main() {
	var qcrao = Person(Student{age: 18})
	fmt.Println(qcrao)
	//Visit()
	//t.PrintDynamicType()
	// a := 2 ^ 3
	// fmt.Print(a)
	//SliceCap()
	// x := outOfRange()
	// fmt.Println(x)
	// WaitDealFile := map[string]string{
	// 	"ShadowBot.UIAutomation":       "UIAutomation",
	// 	"ShadowBot.UIAutomation.Tools": "Tools",
	// }
	// for k, v := range WaitDealFile {
	// 	filePaht := BasePath + k
	// 	Fraversal_folder_all_file(filePaht, v)
	// }
	//UpdateCSSourceFile()

	//Dic_to_excel(result, "RobotApi") //写入excel
	//UpdateSourceFile(result, "cs", "DateTimeExtension", filePaht)
	//GM.BasedVriable()
	//GM.DinamicSlice()

	// nums := TwoSum([]int{9, 8, 7, 6, 5, 4, 3, 2, 1}, 9)
	// fmt.Printf("%v", nums)

	// fmt.Println(IsPalindrome(1000000001))
	// fmt.Println("罗马字转换:", RomanToInt("MCMXCIV"))
}

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
