package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"net"
	"runtime"
	"time"
)

const (
	IP   = ""
	PORT = 1999
)

const (
	PROTOCOL_LEN  = 4
	PROTOCOL_TYPE = 4
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	listen, err := net.ListenTCP("tcp", &net.TCPAddr{net.ParseIP(IP), PORT, ""})
	if err != nil {
		fmt.Println("Server start failed")
	}
	fmt.Println("Server Started...........")
	server(listen)
}

func server(listen *net.TCPListener) {
	for {
		conn, err := listen.AcceptTCP()
		if err != nil {
			fmt.Println("Accept Failed")
			continue
		}
		fmt.Println("Connect Accepted~")
		user_conn := NewUserConnect(conn)
		go user_conn.LoopMessage()
		go TestMap()
	}
}

func TestMap() {
	for {
		fmt.Println(GetConnectMap())
		time.Sleep(time.Second * 3)
	}
}

//-------------------------------------------------------------UserConnect-------------------------------------------------------------
//init con map data
var ConnectMap = make(map[net.Conn]int)

func GetConnectMap() map[net.Conn]int {
	return ConnectMap
}

func AddConnect(conn net.Conn) {
	ConnectMap[conn] = 1
}

func RemoveConnect(conn net.Conn) {
	delete(ConnectMap, conn)
}

func NewUserConnect(conn net.Conn) *UserConnect {
	AddConnect(conn)
	return &UserConnect{conn: conn}
}

type UserConnect struct {
	disconnct bool
	conn      net.Conn
}

func (uc *UserConnect) LoopMessage() {
	defer uc.conn.Close()
	//Read Buffer
	read_buff := make([]byte, 4096)
	//Read Buff Cache
	msg_buff := bytes.NewBuffer(make([]byte, 0, 10240))
	len := 0
	for {
		c, err := uc.conn.Read(read_buff)
		//message, err := reader.ReadString('\n')
		if err == io.EOF {
			RemoveConnect(uc.conn)
			fmt.Println("Client Exit", err.Error())
			break
		}
		if err != nil {
			RemoveConnect(uc.conn)
			fmt.Println("Data read Error", err.Error())
			break
		}
		msg_buff.Write(read_buff[:c])

		uc.handMessage(msg_buff, &len)

		//disconnct flag
		if true == uc.disconnct {
			RemoveConnect(uc.conn)
			break
		}
	}
}

func (uc *UserConnect) handMessage(msg_buff *bytes.Buffer, len *int) {
	msg_head_msg_len := uint32(0)
	for {
		//Read Header
		if *len == 0 && msg_buff.Len() >= 4 {
			msg_head_msg_len = binary.LittleEndian.Uint32(msg_buff.Next(4))
			if msg_head_msg_len > 10240 {
				fmt.Println("too long message")
			}
			*len = int(msg_head_msg_len)
		}

		//Read Body
		if *len != 0 && msg_buff.Len() >= *len {
			msg_content := msg_buff.Next(*len)
			fmt.Println("Msg content is " + string(msg_content))
			*len = 0
		} else {
			break
		}
	}
}
