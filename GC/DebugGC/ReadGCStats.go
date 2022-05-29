package debuggc

import (
	"fmt"
	"runtime/debug"
	"time"
)

//通过代码的方式来直接实现对感兴趣指标的监控
//每一秒监控一下GC的状态
func PrintGCStats() {
	fmt.Print("start listening")
	t := time.NewTicker(time.Second)
	s := debug.GCStats{}
	for {
		select {
		case <-t.C:
			debug.ReadGCStats(&s)
			fmt.Printf("gc %d last@%v, PauseTotal %v\n", s.NumGC, s.LastGC, s.PauseTotal)
		}
	}
}
