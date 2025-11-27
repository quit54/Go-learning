package main

import (
	"fmt"
	"io"
	"net"
	"sync"
)

type Server struct {
	Ip   string
	Port int

	//创建在线用户列表
	OnlineMap map[string]*User
	mapLock   sync.RWMutex

	//有关消息广播的channel
	Message chan string
}

func NewServer(ip string, port int) *Server {
	server := &Server{
		Ip:        ip,
		Port:      port,
		OnlineMap: make(map[string]*User),
		Message:   make(chan string),
	}
	return server
}

// 监听广播massage消息channel的goroutine，一旦有消息，就发送给全部的在线User
func (this *Server) ListenMessage() {
	for {
		msg := <-this.Message

		this.mapLock.Lock()
		for _, cli := range this.OnlineMap {
			cli.C <- msg
		}
		this.mapLock.Unlock()
	}
}

// 广播消息的方法，广播一个什么样子的消息
// 这里的广播消息，是在已经上线的用户（不包含新上线用户）中，还是在新上线的用户中，还是所有在线用户？
func (this *Server) BroadCast(user *User, msg string) {
	Sendmsg := "[" + user.Addr + "]" + user.Name + ":" + msg

	this.Message <- Sendmsg

}

func (this *Server) Handler(conn net.Conn) {
	//fmt.Println("当前业务连接成功")
	//当前用户上线了，将用户加入OnlineMap中，进行广播
	user := NewUser(conn)

	this.mapLock.Lock()

	this.OnlineMap[user.Name] = user
	this.mapLock.Unlock()

	//广播当前用户的上线消息
	this.BroadCast(user, "已上线")
	//接受客户端发来的消息
	go func() {
		buf := make([]byte, 4096)
		for {
			n, err := conn.Read(buf)
			if n == 0 {
				this.BroadCast(user, "下载")
				return
			}
			if err != nil && err != io.EOF {
				fmt.Println("Conn Read err:", err)
				return
			}
			//提取用户消息，去除"\n"，
			msg := string(buf[:n-1])
			//将得到的消息广播
			this.BroadCast(user, msg)
		}
	}()
	//当前hander堵塞
	select {}
}
func (this *Server) Start() {
	//socket listen
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", this.Ip, this.Port))
	if err != nil {
		fmt.Println("连接错误:", err)
		return
	}
	//close listen
	defer listener.Close()

	//启动监听massage的goroutine
	go this.ListenMessage()

	for {
		//accpet listen
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("连接错误:", err)
			continue
		}
		//do handler
		go this.Handler(conn)
	}
}
