package unsafe

import (
	"fmt"
	"unsafe"
)

type PersonalInfo struct {
	name   string
	age    int
	weight float64
}


func Visit() {
	p := PersonalInfo{"张三", 13, 23.433}
	fmt.Print(p)
	age := (*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&p)) + unsafe.Offsetof(p.age)))
	*age = 10
	fmt.Print(p)
}
