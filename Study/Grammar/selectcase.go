package main

import (
	"fmt"
	"time"
)

/*select case 死锁分析和理解*/

func talk(msg string, sleep int) <-chan string {
	ch := make(chan string)
	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println("成功写入：", msg, i)
			ch <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(sleep) * time.Millisecond)
		}
	}()
	return ch
}

/*
现象：本应该输出10次，但是只输出5次，程序发生死锁  all goroutines are asleep - deadlock!
规范：select语句运行:
	(1)在进入该语句时，会按源码的顺序对每一个 case 子句进行求值：这个求值只针对发送或接收操作的额外表达式，自上而下 自左向右。
	所以在这个死锁的例子中会一次执行 <-input1  <-input2
	(2)case的进入规则：
		a.若有一个case的接收或发送可执行，则选择该case执行;
		b.若所有case都不能执行，则有default就执行default，否则就阻塞；
		c.若有多个case同时触发，则随机选择一个去执行；
结论：
	5次输出的原因：每次进来，都会执行 <-input1、<-input2从chan中取出一个值，然而每次随机选择一个执行，必然会导致有一个丢失
	死锁的原因：只有五次输出，但是main里面输出10次，导致chan里没有数据，发生死锁
*/
func fanIn(input1, input2 <-chan string) <-chan string {
	ch := make(chan string)
	go func() {
		for {
			fmt.Println("~~~~~~~~~~~~~~~~~~~~~~进入下一次循环啦~~~~~~~~~~~~~~~~~~~~~~~~`")
			select {
			case ch <- <-input1:
			case ch <- <-input2:
			default:
				println("￥￥￥￥￥￥￥￥￥￥￥￥￥没有执行，走default￥￥￥￥￥￥￥￥￥￥￥￥")
			}

			//该写发不会死锁
			//原因：t := <-input1 里没有只针对发送或接收操作的额外表达式，这是一个发送的表达式，因此不会先计算一遍 <-input1、<-input2
			// select {
			// case t := <-input1:
			// 	ch <- t
			// case t := <-input2:
			// 	ch <- t
			// }
		}
	}()
	return ch
}

/*
经典问题 time.After 导致内存泄露
传递给 time.After 的时间参数越大，泄露会越厉害
*/
func timeafter() {
	ch := make(chan int, 10)

	go func() {
		var i = 1
		for {
			i++
			ch <- i
		}
	}()

	for {
		select {
		case x := <-ch:
			println(x)
		case <-time.After(30 * time.Second):
			println(time.Now().Unix())
		}
	}
}
func main() {
	// ch := fanIn(talk("A", 10), talk("B", 1000))
	// for i := 0; i < 5; i++ {
	// 	fmt.Printf("===================主程序输出：%q=============\n", <-ch)
	// }
	timeafter()
	select {}
}
