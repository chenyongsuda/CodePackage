package main

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
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
	PROTOCOL_MSG_LEN  = 4
	PROTOCOL_TYPE_LEN = 4
)

const (
	CMD_LOGIN   = 100000
	CMD_EXIT    = 100001
	CMD_MSG_P2P = 100002
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
		user_conn := NewUserConnectExt(conn)
		go user_conn.LoopMessage()
		go TestMap()
	}
}

func TestMap() {
	for {
		fmt.Println(GetUserConnectMap())
		fmt.Println(GetConnectMap())
		time.Sleep(time.Second * 8)
	}
}

//-------------------------------------------------------------UserConnect-------------------------------------------------------------
//init con map data
var ConnectMap = make(map[net.Conn]*UserConnect)

// store the User When The User Send Login Message to Server
var userConnectMap = make(map[string]*UserConnect)

func RemoveAllConnction(uc *UserConnect) {
	RemoveUserConnect(uc.user_id)
	RemoveConnect(uc.conn)
}

func GetUserConnectMap() map[string]*UserConnect {
	return userConnectMap
}

func GetUserConnectByUID(uid string) (*UserConnect, bool) {
	val, ok := userConnectMap[uid]
	if ok == true {
		return val, true
	}
	return nil, false
}

func AddUserConnect(uid string, con *UserConnect) {
	_, ok := userConnectMap[uid]
	if ok == true {
		RemoveUserConnect(uid)
	}
	userConnectMap[uid] = con
}

func RemoveUserConnect(uid string) {
	_, ok := userConnectMap[uid]
	if ok == true {
		delete(userConnectMap, uid)
	}
}

func GetConnectMap() map[net.Conn]*UserConnect {
	return ConnectMap
}

func AddConnect(conn net.Conn, val *UserConnect) {
	ConnectMap[conn] = val
}

func RemoveConnect(conn net.Conn) {
	value, ok := ConnectMap[conn]
	if ok == true {
		delete(ConnectMap, conn)
		value.Disconnect()
	}
}

func NewUserConnect(conn net.Conn) *UserConnect {
	return &UserConnect{conn: conn}
}

func NewUserConnectExt(conn net.Conn) *UserConnect {
	uc := NewUserConnect(conn)
	AddConnect(conn, uc)
	return uc
}

type UserConnect struct {
	disconnct bool
	conn      net.Conn
	user_id   string
	user_name string
}

func (uc *UserConnect) Disconnect() {
	uc.SendMessage([]byte("YOU Have DisConnected From Server"))
	uc.disconnct = true
}

func (uc *UserConnect) SendMessage(messge []byte) {
	if false == uc.disconnct {
		_, err := uc.conn.Write(messge)
		if err != nil {
			fmt.Println("Client Exit", err.Error())
		}
	}
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
			RemoveAllConnction(uc)
			fmt.Println("Client Exit", err.Error())
			break
		}
		if err != nil {
			RemoveAllConnction(uc)
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
	msg_head_msg_type := uint32(0)
	for {
		//Read Header
		if *len == 0 && msg_buff.Len() >= (PROTOCOL_MSG_LEN+PROTOCOL_TYPE_LEN) {
			msg_head_msg_len = binary.LittleEndian.Uint32(msg_buff.Next(PROTOCOL_MSG_LEN))
			if msg_head_msg_len > 10240 {
				fmt.Println("too long message")
			}
			msg_head_msg_type = binary.LittleEndian.Uint32(msg_buff.Next(PROTOCOL_TYPE_LEN))
			*len = int(msg_head_msg_len)
		}

		//Read Body
		if *len != 0 && msg_buff.Len() >= *len {
			msg_content := msg_buff.Next(*len)
			fmt.Println("Type:", msg_head_msg_type, "	Msg: ", string(msg_content))

			switch msg_head_msg_type {
			case CMD_LOGIN:
				uc.HandleLogin(msg_content)
			case CMD_MSG_P2P:
				uc.HandleP2PMessage(msg_content)
			case CMD_EXIT:
				uc.HandleLogout(msg_content)
			}
			*len = 0
		} else {
			break
		}
	}
}

func (uc *UserConnect) HandleLogin(message []byte) {
	//fmt.Println(len(message))
	//fmt.Println(string(message))
	jdata := make(map[string]interface{})
	err := json.Unmarshal(message, &jdata)
	if err != nil {
		fmt.Println("Decode Error")
		return
	}
	uc.user_id = jdata["UID"].(string)
	uc.user_name = jdata["UNAME"].(string)
	AddUserConnect(uc.user_id, uc)
	fmt.Println("HandleLogin --- UserID: ", uc.user_id, " UserName: ", uc.user_name)
}

func (uc *UserConnect) HandleLogout(message []byte) {
	jdata := make(map[string]interface{})
	err := json.Unmarshal(message, &jdata)
	if err != nil {
		fmt.Println("Decode Error")
		return
	}
	fmt.Println("HandleLogout --- UserID: ", uc.user_id, " User: ", uc.user_name, "Send LogOut")
	RemoveUserConnect(jdata["UID"].(string))
}

func (uc *UserConnect) HandleP2PMessage(message []byte) {
	jdata := make(map[string]interface{})
	err := json.Unmarshal(message, &jdata)
	if err != nil {
		fmt.Println("Decode Error")
		return
	}
	tUserConnect, ok := GetUserConnectByUID(jdata["TO"].(string))
	if ok == true {
		fmt.Println("HandleP2PMessage --- From:", jdata["FROM"], "FROM_NAME", jdata["FROMNAME"], " TO: ", jdata["TO"], " DATA: ", jdata["DATA"])
		tUserConnect.SendMessage(message)
	}
}
