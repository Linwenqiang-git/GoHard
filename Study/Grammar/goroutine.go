package basekgrammar

import (
	_ "fmt"
)

func Console() {
	defer println("Console End")
	for i := 0; i < 3; i++ {
		println("Call Console, i = ", i)
	}
}

//select 结合channel使用 能够同时监控多个channel的状态
//所有channel表达式都会被求值，所有被发送的表达式都会被求值,求值顺序：自上而下，自左向右，结果是选择一个发送或者接收的channel，
//无论是选择哪一个case进行操作，表达式（比如方法、运算等）都会被执行
//select 默认阻塞，只有监听的channel中有发送或者接受数据时才运行
//如果有一个或多个IO操作可以完成，则GO运行时会随机的选择一个执行，否则，如果有default分支，则执行default，如果没有则会一直阻塞，直到有一个IO可以进行
func selectChannel(c1, c2 chan int) {
	count := 0
	for {
		select {
		case x1 := <-c1: //C1 有数据
			println("success in c1 ,value = ", x1)
		case x2 := <-c2: //c2有数据
			println("success in c2 ,value = ", x2)
		default:
			count += 1
			if count > 20 {
				return
			} else {
				println("not found data...")
			}

		}
	}
}

type sub struct {
	closing chan chan error
}

func (this *sub) Close() error {
	errc := make(chan error)
	this.closing <- errc
	return <-errc
}

func init() {
	//go Console()                //程序是并行调用,主协程执行完，子的会被终止
	//time.Sleep(2 * time.Second) //这里需要阻塞一下，否则子协程无法执行

	/*
		channel 用于协程间的通信机制
		有缓冲的channel，设置初始容量，当通道内容量超过3时，send会阻塞,缓存空了后receive才会阻塞
		没有缓存的channel 只有sender和receiver都准备好了后它们的通讯才会发生，阻塞式的
	*/

	// c := make(chan int, 3)
	// //匿名方法
	// go func() {
	// 	for i := 0; i < 3; i++ {
	// 		time.Sleep(1 * time.Second)
	// 		println("Call Console, i = ", i)
	// 		c <- i
	// 	}
	// 	//不关闭也可，会被自动回收，显示的关闭是通知其他协程某个任务已经完成
	// 	//应该是生产者关闭通道，而不是消费者，否则生产者在channel关闭后写入，会导致panic
	// 	close(c)
	// }()
	// for i := 0; i < 3; i++ {
	// 	println("Main catch")
	// 	//已经被关闭的通道不能写数据，但可以取数据
	// 	num := <-c
	// 	println("Main Console, i = ", num)
	// }
	// //两种写法等价  当通道内没有数据时，range会阻塞
	// for data := range c {
	// 	println("Main Console, i = ", data)
	// }

	//定义单项通道
	// var send chan<- int //只写通道
	// var recv <-chan int //只读通道

	c1 := make(chan int)
	c2 := make(chan int)
	go func() {
		for i := 0; i < 6; i++ {
			c1 <- i
			c2 <- i * 10
		}

	}()
	selectChannel(c1, c2)
	print("Main End")
}
