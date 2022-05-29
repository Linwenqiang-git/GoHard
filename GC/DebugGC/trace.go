package debuggc

import (
	"os"
	"runtime/trace"
)

//go tool trace trace.out
func TarceTool() {
	f, _ := os.Create("trace.out")
	defer f.Close()
	trace.Start(f)
	defer trace.Stop()
	keepalloc3()
}

//模拟goroutine泄露
func keepalloc2() {
	for i := 0; i < 100000; i++ {
		go func() {
			select {}
		}()
	}
}


var ch = make(chan struct{})
//channel泄露导致的goroutine泄露
func keepalloc3() {
	for i := 0; i < 100000; i++ {
		// 没有接收方，goroutine 会一直阻塞
		go func() { ch <- struct{}{} }()
	}
}

func allocate() {
	_ = make([]byte, 1<<20)
}
