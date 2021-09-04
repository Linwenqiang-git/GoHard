package main

import (
	"fmt"
	"io"
	"net"
	"sync"
	"time"
)

type Server struct {
	Ip   string
	Port int
	//连接用户的map
	OnlineMap map[string]*User
	mapLock   sync.RWMutex

	//广播的channel
	Message chan string
}

func (this *Server) Handle(conn net.Conn) {
	user := CreateNewUser(conn, this)
	//将用户添加到上线表汇总
	user.OnLine()
	//广播消息，用户上线
	this.BroadCast(user, "on line")

	//判断用户是否一直活跃
	isActive := make(chan bool)

	//接收当前连接用户发来的消息
	go func() {
		buf := make([]byte, 4096)
		for {
			n, err := conn.Read(buf)
			if n == 0 {
				this.BroadCast(user, "out line")
				user.OutLine()
				return
			}
			if err != nil && err != io.EOF { //io.EOF 文件结束末尾
				println("read conn data err:", err)
			}
			msg := string(buf[:n])
			user.DoMsg(msg, n)
			isActive <- true
		}
	}()

	//阻塞当前handle
	//超时强踢功能
	for {
		select {
		case <-isActive:
			//do nothing
		case <-time.After(time.Second * 10):
			//off line
			this.BroadCast(user, "outtime already outline")
			//删除用户列表
			this.mapLock.Lock()
			delete(this.OnlineMap, user.Name)
			this.mapLock.Unlock()
			//关闭连接
			conn.Close()
		}
	}

}

//广播消息
func (this *Server) BroadCast(user *User, message string) {
	msg := "[" + user.Addr + "]" + user.Name + ":" + message
	//将这个消息添加到广播通道
	this.Message <- msg
}

//监听广播channel 一旦有消息 就发送 给所有在线用户
func (this *Server) ListenMessager() {
	for {
		msg := <-this.Message
		this.mapLock.Lock()
		for _, v := range this.OnlineMap {
			v.C <- msg //向在线用户写入消息
		}
		this.mapLock.Unlock()
	}
}

//启动一个服务
func (this *Server) Start() {
	//1. socket listen
	//创建套接字连接
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", this.Ip, this.Port))
	if err != nil {
		print("create net.Listen error:", err)
	}
	//4.close listen socket 最后记得关闭连接
	defer listener.Close()
	//监听处理广播消息
	print(">>>>>>>>>>>server start>>>>>>>>>>>")
	go this.ListenMessager()
	//循环的接收消息
	for {
		//2.accept msg
		conn, err := listener.Accept() //监听客户端连接
		if err != nil {
			println("creat conn error:", err)
			continue
		}

		//3.handle msg 开启一个协程去处理，程序继续监听连接
		go this.Handle(conn)
	}

}

//返回一个server类型的指针
func NewServer(ip string, port int) *Server {
	server := &Server{
		Ip:        ip,
		Port:      port,
		OnlineMap: make(map[string]*User),
		Message:   make(chan string),
	}
	return server
}
