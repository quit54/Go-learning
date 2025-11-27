package main

import "net"

type User struct {
	Name string
	Addr string
	C    chan string
	Conn net.Conn
}

//创建一个用户的API
func NewUser(conn net.Conn) *User {
	userAddr := conn.RemoteAddr().String()
	user := &User{
		Name: userAddr,
		Addr: userAddr,
		C:    make(chan string),
		Conn: conn,
	}
	//启动监听user channel的goroutine
	go user.ListenMessage()
	return user
}

//监听当前User channel的方法，一旦有消息，就直接发送给对端客户端。
func (this *User) ListenMessage() {
	for {
		msg := <-this.C

		this.Conn.Write([]byte(msg))
	}
}
