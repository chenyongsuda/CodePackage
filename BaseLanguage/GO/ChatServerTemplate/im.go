package main

import (
	"net"
)

var route *Route

func main() {
	route = NewRoute()
	go ListenPeerClient()
	ListenClient()
}

func Listen(ip string, port int, f func(*net.TCPConn)) {
	addr := net.TCPAddr{net.ParseIP(ip), port, ""}
	listen, err := net.ListenTCP("tcp", &addr)
	if err != nil {
		return
	}

	for {
		client, err := listen.AcceptTCP()
		if err != nil {

		}
		f(client)
	}
}

func ListenClient() {
	Listen("0.0.0.0", 9081, HandleClient)
}

func ListenPeerClient() {
	Listen("0.0.0.0", 9082, HandPeerClient)
}

func HandleClient(conn *net.TCPConn) {
	client := NewClient(conn)
	client.Run()
}

func HandPeerClient(conn *net.TCPConn) {

}
