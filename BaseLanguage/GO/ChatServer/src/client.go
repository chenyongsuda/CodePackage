package main

import (
	//"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
	//"os"
	"time"
)

const (
	ADDR = "127.0.0.1:1999"
)

func main() {
	conn, err := net.Dial("tcp", ADDR)
	if err != nil {
		fmt.Println("连接服务器失败")
		return
	}
	fmt.Println("已经连接服务器")
	Client(conn)
	conn.Close()
}

func Client(conn net.Conn) {

	message := []byte("我是UTF-8")
	tlen := 4 + len(message)
	send_buff := bytes.NewBuffer(make([]byte, 0, 1024))

	for {
		for i := 0; i < 10; i++ {
			binary.Write(send_buff, binary.LittleEndian, uint32(len(message)))
			send_buff.Write(message)
		}

		fmt.Println("发送一整条信息：")
		conn.Write(send_buff.Next(tlen))
		time.Sleep(time.Second)

		//发送不完整head信息
		fmt.Println("发送不完整head信息：")
		conn.Write(send_buff.Next(2))
		time.Sleep(time.Second)
		conn.Write(send_buff.Next(tlen - 2))
		time.Sleep(time.Second)

		fmt.Println("发送3条信息：")
		time.Sleep(time.Second)
		conn.Write(send_buff.Next(3 * tlen))

		fmt.Println("发送不全的消息体：")
		time.Sleep(time.Second)
		conn.Write(send_buff.Next(6))
		time.Sleep(time.Second)
		conn.Write(send_buff.Next(tlen - 6))

		fmt.Println("多段发送：")
		conn.Write(send_buff.Next(tlen + 2))
		time.Sleep(time.Second)
		conn.Write(send_buff.Next(-2 + tlen - 8))
		time.Sleep(time.Second)
		conn.Write(send_buff.Next(8 + 1))
		time.Sleep(time.Second)
		conn.Write(send_buff.Next(-1 + tlen + tlen))

	}

}
