package main

import (
	"time"
	tarce "github.linwenqiang.com/gc/debuggc"
)

func main() {
	// trace.TarceTool()
	// go trace.PrintGCStats()
	go tarce.TarceTool()
	time.Sleep(1000 * time.Second)

}
