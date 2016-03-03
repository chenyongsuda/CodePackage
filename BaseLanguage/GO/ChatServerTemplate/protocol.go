package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"net"
)

const (
	MSG_AUTH = 2
	MSG_IM   = 3
)

const (
	ERROR_READ_ERR = -1
)

const (
	INIT_LEN = -1
	INIT_CMD = -1
	INIT_SEQ = -1
)

type Message struct {
	cmd  int
	seq  int
	body interface{}
}

type Authentication struct {
	uid int64
}

type IMMessage struct {
	sender  int64
	receive int64
	msgid   int32
	content string
}

func ReceiveMessage(conn *net.TCPConn, msg_buff *bytes.Buffer) bool {
	//Read Buffer
	read_buff := make([]byte, 4096)
	c, err := conn.Read(read_buff)
	//message, err := reader.ReadString('\n')
	if err == io.EOF {
		return true
	}
	if err != nil {
		return true
	}
	msg_buff.Write(read_buff[:c])
	return false
}

func PraseMessage(msg_buff *bytes.Buffer, msg_len *int, msg_seq *int, msg_cmd *int) *Message {
	//Read Header
	if *msg_len == INIT_LEN && msg_buff.Len() >= 12 {
		//fmt.Printf("%q\n", msg_buff.Bytes())
		var cmd int32
		var len int32
		var seq int32
		binary.Read(msg_buff, binary.BigEndian, &len)
		if *msg_len > 10240 {
			fmt.Println("Too Long Message")
		}
		binary.Read(msg_buff, binary.BigEndian, &cmd)
		binary.Read(msg_buff, binary.BigEndian, &seq)

		*msg_len = int(len)
		*msg_cmd = int(cmd)
		*msg_seq = int(seq)

		fmt.Println("cmd:", *msg_cmd, "msglen:", *msg_len, "seq:", *msg_seq, "msglen", msg_buff.Len())
	}

	var ret_msg *Message
	//Do Get Data
	if *msg_len != INIT_LEN && msg_buff.Len() >= *msg_len {
		if *msg_cmd == MSG_AUTH {
			var uid int64
			binary.Read(msg_buff, binary.BigEndian, &uid)
			ret_msg = &Message{*msg_cmd, *msg_seq, &Authentication{uid}}
		} else if *msg_cmd == MSG_IM {
			var sender int64
			var receive int64
			var msgid int32
			var content string
			binary.Read(msg_buff, binary.BigEndian, &sender)
			binary.Read(msg_buff, binary.BigEndian, &receive)
			binary.Read(msg_buff, binary.BigEndian, &msgid)
			content = string(msg_buff.Next(*msg_len - 20))
			ret_msg = &Message{*msg_cmd, *msg_seq, &IMMessage{sender, receive, msgid, content}}
		} else {
			ret_msg = nil
		}

		//Reset Param
		*msg_len = INIT_LEN
		*msg_cmd = INIT_CMD
		*msg_seq = INIT_SEQ
		return ret_msg
	} else {
		return nil
	}

}

func SendMessage(conn *net.TCPConn, message *Message) {
	if message.cmd == MSG_IM {
		im_msg := message.body.(*IMMessage)
		WriteIMMessage(conn, int32(message.cmd), int32(message.seq), im_msg)
	}
}

func WriteHeader(len int32, cmd int32, seq int32, buffer *bytes.Buffer) {
	binary.Write(buffer, binary.BigEndian, len)
	binary.Write(buffer, binary.BigEndian, cmd)
	binary.Write(buffer, binary.BigEndian, seq)
}

func WriteIMMessage(conn *net.TCPConn, cmd int32, seq int32, msg *IMMessage) {
	buffer := new(bytes.Buffer)
	len := int32(len(msg.content) + 20)
	WriteHeader(len, cmd, seq, buffer)
	binary.Write(buffer, binary.BigEndian, msg.sender)
	binary.Write(buffer, binary.BigEndian, msg.receive)
	binary.Write(buffer, binary.BigEndian, msg.msgid)
	buffer.Write([]byte(msg.content))
	_, err := conn.Write(buffer.Bytes())
	if err != nil {
		fmt.Println("write socket error")
	}
}
