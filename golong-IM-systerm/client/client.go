package main

import (
	"flag"
	"fmt"
	"io"
	"net"
	"os"
)

type Client struct {
	ServerIp   string
	ServerPort int
	Name       string
	conn       net.Conn
	flag       int
}

// 处理server回应的消息，直接显示到标准输出即可
func (client *Client) DealResponse() {
	//一旦conn上面有数据，就copy到stdout的标准输出上，永久阻塞监听
	io.Copy(os.Stdout, client.conn)
	// for{
	// 	buf := make()
	// 	client.conn.Read(buf)
	// 	fmt.Println(">>>>>",string(buf))
	// }
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
func (client *Client) PubliChat() {
	var ChatMsg string

	//提示用户输入信息
	fmt.Println(">>>>>>请输入聊天内容:\nexit表示退出\n")
	fmt.Scanln(&ChatMsg)

	for ChatMsg != "exit" {
		if len(ChatMsg) != 0 {
			sendMsg := ChatMsg + "\n"
			_, err := client.conn.Write([]byte(sendMsg))
			if err != nil {
				fmt.Println("conn.Write err:", err)
				break
			}
		}
		fmt.Println(">>>>>>请输入聊天内容:")
		fmt.Scanln(&ChatMsg)
	}
}
func (client *Client) SelectUsers() {
	sendMsg := "who\n"
	_, err := client.conn.Write([]byte(sendMsg))
	if err != nil {
		fmt.Println("conn.Write err:", err)
		return
	}
}
func (client *Client) PrivateChat() {
	var remoteName string
	var ChatMsg string

	client.SelectUsers()
	fmt.Println("请输入当前聊天对象的用户名,exit表示退出")
	fmt.Scanln(&remoteName)

	for remoteName != "exit" {
		fmt.Println(">>>>>>请输入聊天内容:")
		fmt.Scanln(&ChatMsg)

		for len(ChatMsg) != 0 {
			sendMsg := "to|" + remoteName + "|" + ChatMsg + "\n"
			_, err := client.conn.Write([]byte(sendMsg))
			if err != nil {
				fmt.Println("conn.Write err:", err)
				break
			}
			ChatMsg = ""
			fmt.Println(">>>>>>请输入聊天内容:")
			fmt.Scanln(&ChatMsg)
		}
		fmt.Println("请输入当前聊天对象的用户名,exit表示退出")
		fmt.Scanln(&remoteName)
	}

}
func (client *Client) UpdateName() bool {
	fmt.Println(">>>>>>请输入用户名:")
	fmt.Scanln(&client.Name)

	sendMsg := "rename|" + client.Name + "\n"
	_, err := client.conn.Write([]byte(sendMsg))
	if err != nil {
		fmt.Println("cliet.conn.Write:", err)
		return false
	}
	return true
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
			client.PubliChat()
			break
		case 2:
			client.PrivateChat()
			break
		case 3:
			client.UpdateName()
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
	go client.DealResponse()
	//单独开出来一个goroutine处理server返回的消息

	//保持客户端运行逻辑
	client.Run()
}
