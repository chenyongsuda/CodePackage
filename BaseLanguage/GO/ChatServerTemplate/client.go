package main

import (
	"bytes"
	"fmt"
	"net"
)

type Client struct {
	conn *net.TCPConn
	wt   chan *Message
	uid  int64
}

func NewClient(conn *net.TCPConn) *Client {
	client := new(Client)
	client.conn = conn
	client.wt = make(chan *Message)
	return client
}

func (c *Client) Run() {
	fmt.Println("New Socket In")
	go c.Read()
	go c.Write()
}

func (c *Client) Read() {
	msg_buff := bytes.NewBuffer(make([]byte, 0, 10240))
	var msg_len int = INIT_LEN
	var msg_cmd int = INIT_CMD
	var msg_seq int = INIT_SEQ
	for {
		err := ReceiveMessage(c.conn, msg_buff)
		if err == true {
			c.wt <- nil
			//fmt.Println("Receive Message err")
			break
		}
		message := PraseMessage(msg_buff, &msg_len, &msg_cmd, &msg_seq)
		if message == nil {
			//fmt.Println("Prase Message err")
			continue
		}
		//Get Message
		if message.cmd == MSG_AUTH {
			c.HandleAuth(message.body.(*Authentication))
		} else if message.cmd == MSG_IM {
			c.HandleMessage(message.body.(*IMMessage), message.seq)
		} else {

		}
	}
}

func (c *Client) Write() {
	for {
		message := <-c.wt
		if message == nil {
			route.RemoveClient(c)
			c.conn.Close()
			fmt.Println("Socket Closed")
			break
		}
		//SendMessage is implement by protocol
		SendMessage(c.conn, message)
	}
}

func (c *Client) HandleAuth(login *Authentication) {
	c.uid = login.uid
	route.AddClient(c)
}

func (c *Client) HandSocketClosed() {
	route.RemoveClient(c)
	c.wt <- nil
}

func (c *Client) HandleMessage(msg *IMMessage, seq int) {
	//Send Message
	route.RouteMessage(&Message{cmd: MSG_IM, body: msg})
	//Send Ack
}
