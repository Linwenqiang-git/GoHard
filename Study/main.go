package main

import (

	//. "/DataLogic/dataLogic"  //_ 表示会走包的init方法，下面不调用也OK
	"fmt"

	_ "github.linwenqiang.com/Study/DataLogic" // . 表示调用内部公开方法时可不用写包名
	//DL "/DataLogic/dataLogic" //DL 给包起的别名
	//_ "github.linwenqiang.com/Study/Grammar"
	//. "github.linwenqiang.com/Study/Test"
)

//声明全局变量
var (
	a int
	b bool
)

func outOfRange() int {
	arr := [3]int{1, 2, 3}
	i := 4
	elem := arr[i]
	return elem
}
func main() {
	//SliceCap()
	x := outOfRange()
	fmt.Println(x)
	//GM.BasedVriable()
	//GM.DinamicSlice()

	// nums := TwoSum([]int{9, 8, 7, 6, 5, 4, 3, 2, 1}, 9)
	// fmt.Printf("%v", nums)

	// fmt.Println(IsPalindrome(1000000001))
	// fmt.Println("罗马字转换:", RomanToInt("MCMXCIV"))
}
