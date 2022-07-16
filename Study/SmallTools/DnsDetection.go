package smalltools

import (
	"fmt"
	"log"
	"net"
	"os"
	"time"
)

var (
	Error *log.Logger
	Info  *log.Logger
	Warn  *log.Logger
)

func DnsDetect() {
	fmt.Println("正在解析运行，[www.baidu.com][winrobot360.com],请勿关闭窗口...")
	initLog()
	Info.Print("开始解析www.baidu.com....")
	Info.Print("开始解析winrobot360.com....")
	for {
		_, err := net.LookupHost("www.baidu.com")
		_, err2 := net.LookupHost("winrobot360.com")
		if err != nil {
			//fmt.Fprintf(os.Stderr, "Err: %s", err.Error())
			Error.Printf("Err: %s", err.Error())
			goto sleep
		}
		if err2 != nil {
			//fmt.Fprintf(os.Stderr, "Err: %s", err.Error())
			Error.Printf("Err: %s", err2.Error())
			goto sleep
		}

		// for _, n := range ns {
		// 	Info.Printf("--%s\n", n)
		// 	// fmt.Fprintf(os.Stdout, )
		// }
	sleep:
		time.Sleep(time.Second * 30)
	}
}

func initLog() {
	var Outfile *os.File
	var err error
	Outfile, err = os.OpenFile("./DnsDetect.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	//fmt.Println(reflect.TypeOf(outfile))
	if err != nil {
		log.Panicf("open log file fail:%s ", err)
	}
	Error = log.New(Outfile, "[error]", log.Ldate|log.Ltime|log.Lshortfile)
	Info = log.New(Outfile, "[info]", log.Ldate|log.Ltime|log.Lshortfile)
}
