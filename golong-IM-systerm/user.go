package main

import (
	"fmt"
	"net"
	"strings"
)

type User struct {
	Name string
	Addr string
	C    chan string
	Conn net.Conn

	server *Server
}

// 创建一个用户的API
func NewUser(conn net.Conn, server *Server) *User {
	userAddr := conn.RemoteAddr().String()
	user := &User{
		Name: userAddr,
		Addr: userAddr,
		C:    make(chan string),
		Conn: conn,

		server: server,
	}
	//启动监听user channel的goroutine
	go user.ListenMessage()
	return user
}

// 上线功能
func (this *User) Online() {

	this.server.mapLock.Lock()

	this.server.OnlineMap[this.Name] = this
	this.server.mapLock.Unlock()

	//广播当前用户的上线消息
	upmsg := fmt.Sprintf("[%s]%s:已上线\n", this.Addr, this.Name)
	this.server.BroadCast(this, upmsg)
}

// 下线功能
func (this *User) Offline() {
	this.server.mapLock.Lock()
	delete(this.server.OnlineMap, this.Name)
	this.server.mapLock.Unlock()

	//广播当前用户的上线消息
	downmsg := fmt.Sprintf("[%s]%s:已下线\n", this.Addr, this.Name)
	this.server.BroadCast(this, downmsg)
}

func (this *User) Sendmsg(msg string) {
	this.Conn.Write([]byte(msg))
}

func (this *User) DoMessage(msg string) {
	if msg == "who" {
		//查询在线用户都有哪些？

		this.server.mapLock.Lock()
		for _, user := range this.server.OnlineMap {
			onlinemsg := fmt.Sprintf("[%s][%s]在线....\n", user.Addr, user.Name)
			this.Sendmsg(onlinemsg)
		}
		this.server.mapLock.Unlock()

	} else if len(msg) > 7 && msg[:7] == "rename|" {
		//消息格式：“rename|张三”
		newName := strings.Split(msg, "|")[1]

		//判断name是否存在
		_, ok := this.server.OnlineMap[newName]
		if ok {
			this.Sendmsg("当前用户名被使用\n")
		} else {
			this.server.mapLock.Lock()
			delete(this.server.OnlineMap, this.Name)
			this.server.OnlineMap[newName] = this
			this.server.mapLock.Unlock()

			this.Name = newName
			this.Sendmsg("您已经更新用户名：" + this.Name + "\n")
		}
	} else {
		this.server.BroadCast(this, msg)
	}
}

// 监听当前User channel的方法，一旦有消息，就直接发送给对端客户端。
func (this *User) ListenMessage() {
	for {
		msg := <-this.C

		this.Conn.Write([]byte(msg))
	}
}
