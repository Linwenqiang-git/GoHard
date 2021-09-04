package main

import (
	"flag"
	"fmt"
	"net"
)

type Client struct {
	Name       string
	ServerIp   string
	ServerPort int
	Conn       net.Conn
}

//创建一个新的客户端
func NewClient(serverIp string, serverPort int) *Client {
	client := &Client{
		ServerIp:   serverIp,
		ServerPort: serverPort,
	}
	println("serverIp:", serverIp)
	println("serverIp:", serverPort)
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", serverIp, serverPort))
	if err != nil {
		fmt.Println("创建套接字连接失败:", err)
		return nil
	}
	println("conn.LocalAddr():", conn.LocalAddr())
	client.Conn = conn
	return client
}

//全局变量
var serverIp string
var serverPort int

func init() {
	flag.StringVar(&serverIp, "ip", "127.0.0.1", "设置服务器IP地址（默认127.0.0.1）")
	flag.IntVar(&serverPort, "port", 8888, "设置服务器端口号（默认8888）")
}

func main() {
	//从命令行解析
	flag.Parse()
	// println("serverIp:", serverIp)
	// println("serverIp:", serverPort)
	client := NewClient(serverIp, serverPort)
	if client == nil {
		fmt.Println(">>>>>>>连接服务端失败>>>>>>>")
		return
	}
	fmt.Println(">>>>>>>连接服务器成功>>>>>>>")
	select {}
}
