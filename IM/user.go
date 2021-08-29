package main

import "net"

type User struct {
	Name string
	Addr string
	C    chan string //每个用户对应监听的channel 数据流是string类型
	conn net.Conn    //用户的连接

	BelongServer *Server // 所属server
}

//用户上线
func (this *User) OnLine() {
	this.BelongServer.mapLock.Lock()
	this.BelongServer.OnlineMap[this.Name] = this
	this.BelongServer.mapLock.Unlock()
}

//用户下线
func (this *User) OutLine() {

	this.BelongServer.mapLock.Lock()
	delete(this.BelongServer.OnlineMap, this.Name)
	this.BelongServer.mapLock.Unlock()
}

//用户发送消息
func (this *User) DoMsg(buf []byte, n int) {
	//提取用户信息
	msg := string(buf[:n])
	//再将信息广播
	this.BelongServer.BroadCast(this, msg)
}

//监听有没有消息
func (this *User) ListenMessage() {
	for {
		//监听通道的消息 阻塞式
		msg := <-this.C
		//将消息写进连接
		this.conn.Write([]byte(msg + "\n"))
	}
}

//创建一个用户连接
func CreateNewUser(conn net.Conn, server *Server) *User {
	userAddr := conn.RemoteAddr().String()

	user := &User{
		Name:         userAddr,
		Addr:         userAddr,
		C:            make(chan string),
		conn:         conn,
		BelongServer: server,
	}

	//开启一个携程监听消息
	go user.ListenMessage()
	return user
}
