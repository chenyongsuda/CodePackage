package main

import (
	"fmt"
)

type Route struct {
	clients map[int64]*Client
}

func NewRoute() *Route {
	route := new(Route)
	route.clients = make(map[int64]*Client)
	return route
}

func (r *Route) AddClient(c *Client) {
	if t, ok := r.clients[c.uid]; ok {
		if t != c {
			/*请注意在原先客户端被踢掉的时候做两件事1.去掉Map中的值2.给自己一个nil信号
			  如果这里写成先给自己一个信号wt<-nil由于这个不是立即执行的所有会有时许问题
			*/
			t.HandSocketClosed()
			r.clients[c.uid] = c
			fmt.Println("Replace The Old Client ", c.uid)
		} else {
		}
	} else {
		r.clients[c.uid] = c
		fmt.Println("Add Client uid :", c.uid)
	}
}

func (r *Route) RemoveClient(c *Client) {
	if t, ok := r.clients[c.uid]; ok {
		if t == c {
			delete(r.clients, c.uid)
			fmt.Println("Remove Client uid :", c.uid)
		}
	}
}

func (r *Route) RouteMessage(msg *Message) {
	if msg.cmd == MSG_IM {
		im_msg := msg.body.(*IMMessage)
		receiver := r.FindClient(im_msg.receive)
		if receiver != nil {
			receiver.wt <- msg
		} else {

		}
	}
}

func (r *Route) FindClient(uid int64) *Client {
	if t, ok := r.clients[uid]; ok {
		return t
	}
	return nil
}
