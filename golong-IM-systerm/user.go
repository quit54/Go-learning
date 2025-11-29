package main

import "net"

type User struct {
	Name string
	Addr string
	C    chan string
	Conn net.Conn

	server *Server
}

//创建一个用户的API
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

//上线功能
func (this *User) Online() {

	this.server.mapLock.Lock()

	this.server.OnlineMap[this.Name] = this
	this.server.mapLock.Unlock()

	//广播当前用户的上线消息
	this.server.BroadCast(this, "已上线")
}

//下线功能
func (this *User) Offline() {
	this.server.mapLock.Lock()
	delete(this.server.OnlineMap, this.Name)
	this.server.mapLock.Unlock()

	//广播当前用户的上线消息
	this.server.BroadCast(this, "下线")
}

func (this *User) Sendmsg(msg string) {
	this.Conn.Write([]byte(msg))
}

func (this *User) DoMessage(msg string) {
	if msg == "who" {
		//查询在线用户都有哪些？

		this.server.mapLock.Lock()
		for _, user := range this.server.OnlineMap {
			onlinemsg := "[" + user.Addr + "]" + user.Name + "在线...."
			this.Sendmsg(onlinemsg)
		}
		this.server.mapLock.Unlock()

	}
	this.server.BroadCast(this, msg)
}

//监听当前User channel的方法，一旦有消息，就直接发送给对端客户端。
func (this *User) ListenMessage() {
	for {
		msg := <-this.C

		this.Conn.Write([]byte(msg))
	}
}
