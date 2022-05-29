package main

import (
	algorithm "github.linwenqiang.com/LeetCode/Algorithm"
)

func main() {
	algorithm.Call_33()
	for n := 1; n < 100000; n++ {
		allocate()
	}
}

func allocate() {
    _ = make([]byte, 1<<20)
}
