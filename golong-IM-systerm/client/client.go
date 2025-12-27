package main

import (
	"flag"
	"fmt"
	"net"
)

type Client struct {
	ServerIp   string
	ServerPort int
	Name       string
	conn       net.Conn
	flag       int
}

func NewClient(serverip string, serverport int) *Client {
	client := &Client{
		ServerIp:   serverip,
		ServerPort: serverport,
		flag:       999,
	}
	Conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", serverip, serverport))
	if err != nil {
		fmt.Println("连接服务器失败:", err)
		return nil
	}
	client.conn = Conn

	return client
}
func (client *Client) menu() bool {
	var flag int
	fmt.Println("1.公聊模式")
	fmt.Println("2.私聊模式")
	fmt.Println("3.更新用户列表")
	fmt.Println("0.退出系统")

	fmt.Scanln(&flag)

	if flag >= 0 && flag <= 3 {
		client.flag = flag
		return true
	} else {
		fmt.Println(">>>>请输入合法数字")
		return false
	} //检查flag输入是否合法
}

var serverIp string
var serverPort int

func init() {
	//命令行解析，调用库函数
	flag.StringVar(&serverIp, "ip", "127.0.0.1", "设置服务器IP地址:(默认是127.0.0.1)")
	flag.IntVar(&serverPort, "port", 8888, "设置服务器端口(默认是8888")
}
func (client *Client) Run() {
	for client.flag != 0 {
		for client.menu() != true {
		}
		switch client.flag {
		case 1:
			fmt.Println("公聊模式选择--------")
			break
		case 2:
			fmt.Println("私聊模式选择---------")
			break
		case 3:
			fmt.Println("更新用户列表选择--------")
			break
		}
	}
}
func main() {
	flag.Parse()

	client := NewClient(serverIp, serverPort)
	if client == nil {
		fmt.Println(">>>>>>客户端连接失败。。。")
		return
	}
	fmt.Println(">>>>>>连接服务器成功。。。")

	//保持客户端运行逻辑
	client.Run()
}
