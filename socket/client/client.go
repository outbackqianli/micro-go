package main

import (
	"fmt"
	"net"
	socket "outback/micro-go/socket/common"
	"time"
)

func main() {
	// net listen 函数 传入socket类型和ip端口，返回监听对象
	listener, err := net.Listen(socket.Server_NetWorkType, socket.Server_Address)
	if err == nil {
		// 循环等待客户端访问
		for {
			conn, err := listener.Accept()
			if err == nil {
				// 一旦有外部请求，并且没有错误 直接开启异步执行
				go handleConn(conn)
			}
		}
	} else {
		fmt.Println("server error", err)
	}
	defer listener.Close()
}

func handleConn(conn net.Conn) {
	for {
		// 设置读取超时时间
		conn.SetReadDeadline(time.Now().Add(time.Second * 2))
		// 调用公用方法read 获取客户端传过来的消息。
		if str, err := socket.Read(conn); err == nil {
			fmt.Println("client:", conn.RemoteAddr(), str)
			// 通过write 方法往客户端传递一个消息
			socket.Write(conn, "server got:"+str)
		}
	}
}
