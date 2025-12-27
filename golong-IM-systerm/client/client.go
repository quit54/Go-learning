package main

import (
	"fmt"
	"net"
)

type Client struct {
	ServerIp   string
	ServerPort int
	Name       string
	conn       net.Conn
}

func NewClient(serverip string, serverport int) *Client {
	client := &Client{
		ServerIp:   serverip,
		ServerPort: serverport,
	}
	Conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", serverip, serverport))
	if err != nil {
		fmt.Println("连接服务器失败:", err)
		return nil
	}
	client.conn = Conn

	return client
}
func main() {
	client := NewClient("127.0.0.1", 8888)
	if client == nil {
		fmt.Println(">>>>>>客户端连接失败。。。")
		return
	}
	fmt.Println(">>>>>>连接服务器成功。。。")

	//保持客户端运行逻辑
	for {
		select {
		default:
		}
	}
}
