package main

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"os"
)

const (
	CMD_LOGIN   = 2
	CMD_EXIT    = 100001
	CMD_MSG_P2P = 100002
)

const (
	ADDR = "127.0.0.1:9081"
)

func main() {
	conn, err := net.Dial("tcp", ADDR)
	if err != nil {
		fmt.Println("连接服务器失败")
		return
	}
	fmt.Println("已经连接服务器")
	defer conn.Close()
	go LoopReceive(conn)
	Client(conn)

}

func Client(conn net.Conn) {
	read := bufio.NewReader(os.Stdin)
	for {
		data, _, _ := read.ReadLine()
		command := string(data)
		if command == "c" {
			Login(conn, "1000", "TOM")
		} else {
			Talk(conn, "1000", "TOM", "1001", string(data))
		}
	}
}

func LoopReceive(conn net.Conn) {
	read_buff := make([]byte, 4096)
	for {
		c, err := conn.Read(read_buff)
		//message, err := reader.ReadString('\n')
		if err == io.EOF {
			fmt.Println("Client Exit", err.Error())
			break
		}
		if err != nil {
			fmt.Println("Data read Error", err.Error())
			break
		}
		fmt.Println("LoopReceive	", c)
		fmt.Println(string(read_buff[:c]))
	}
}

func Login(conn net.Conn, uid string, name string) {
	cmd_type := uint32(CMD_LOGIN)

	send_buff := bytes.NewBuffer(make([]byte, 0, 1024))
	binary.Write(send_buff, binary.BigEndian, uint32(8))
	binary.Write(send_buff, binary.BigEndian, cmd_type)
	binary.Write(send_buff, binary.BigEndian, uint32(0))
	binary.Write(send_buff, binary.BigEndian, uint64(1111))
	conn.Write(send_buff.Bytes())
}

func Talk(conn net.Conn, from string, from_name string, to string, data string) {
	login_data := make(map[string]interface{})
	cmd_type := uint32(CMD_MSG_P2P)
	login_data["FROM"] = from
	login_data["FROMNAME"] = from_name
	login_data["TO"] = to
	login_data["DATA"] = data
	jData, err := json.Marshal(login_data)
	if err != nil {
		panic(err)
	}
	send_buff := bytes.NewBuffer(make([]byte, 0, 1024))
	binary.Write(send_buff, binary.LittleEndian, uint32(len(jData)))
	binary.Write(send_buff, binary.LittleEndian, cmd_type)
	send_buff.Write(jData)
	conn.Write(send_buff.Bytes())
}
