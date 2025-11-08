package main

import (
	"fmt"
	"net"
)

type Server struct {
	Ip   string
	Port int
}

func NewServer(ip string, port int) *Server {
	server := &Server{
		Ip:   ip,
		Port: port,
	}
	return server
}
func (this *Server) Handler(conn net.Conn) {
	fmt.Println("当前业务连接成功")
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
