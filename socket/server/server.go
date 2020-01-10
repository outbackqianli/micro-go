package main

import (
	"fmt"
	"net"
	socket "outback/micro-go/socket/common"
)

func main() {
	// 调用net包中的dial 传入ip 端口 进行拨号连接，通过三次握手之后获取到conn
	conn, err := net.Dial(socket.Server_NetWorkType, socket.Server_Address)
	if err != nil {
		fmt.Println("Client create conn error err:", err)
	}
	defer conn.Close()
	//往服务端传递消息
	socket.Write(conn, "aaaa")
	//读取服务端返回的消息
	if str, err := socket.Read(conn); err == nil {
		fmt.Println(str)
	}
}
