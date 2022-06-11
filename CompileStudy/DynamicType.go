package compilestudy

import (
	"fmt"
	"unsafe"
)

type iface struct {
	itab, data uintptr
}

//动态类型也就是他的具体类型
func PrintDynamicType() {
	var a interface{} = nil         //动态类型和动态值都是nil
	var b interface{} = (*int)(nil) //动态类型是*int，动态值是nil
	x := 5
	var c interface{} = (*int)(&x) //动态类型是*int，动态值是x的地址

	ia := *(*iface)(unsafe.Pointer(&a))
	ib := *(*iface)(unsafe.Pointer(&b))
	ic := *(*iface)(unsafe.Pointer(&c))

	fmt.Println(ia)
	fmt.Println(ib)
	fmt.Println(ic)
	fmt.Println(*(*int)(unsafe.Pointer(ic.data))) //获取到动态值的unsafe指针在转成int指针，在*取值
}
