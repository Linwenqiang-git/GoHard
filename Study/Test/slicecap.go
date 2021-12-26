package test

import (
	"fmt"
	"sync"
)

func SliceCap() {
	var rw = new(sync.RWMutex)

	var mutex sync.Mutex
	//切片扩容
	currentcap := 0
	arr := make([]int, 0)
	for i := 0; i < 2000; i++ {
		if currentcap != cap(arr) {
			fmt.Println("len 为", len(arr), "cap 为", cap(arr))
			currentcap = cap(arr)
		}
		mutex.Lock()
		arr = append(arr, i)
		mutex.Unlock()
	}

}
