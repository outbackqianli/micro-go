package socket

import (
	"bytes"
	"net"
)

//公共函数 用来定义Socket类型 ip 端口。
const (
	Server_NetWorkType = "tcp"
	Server_Address     = "127.0.0.1:8085"
	Delimiter          = '\t'
)

// 往conn中写数据，可以用于客户端传输给服务端， 也可以服务端返回客户端
func Write(conn net.Conn, content string) (int, error) {
	var buffer bytes.Buffer
	buffer.WriteString(content)
	buffer.WriteByte(Delimiter)

	return conn.Write(buffer.Bytes())
}

// 从conn中读取字节流，以上面的结束符为标记
func Read(conn net.Conn) (string, error) {
	readBytes := make([]byte, 1)
	var buffer bytes.Buffer
	for {
		if _, err := conn.Read(readBytes); err != nil {
			return "", err
		}
		readByte := readBytes[0]
		if readByte == Delimiter {
			break
		}
		buffer.WriteByte(readByte)
	}

	return buffer.String(), nil
}
