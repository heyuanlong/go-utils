package net

import (
	klog "log"
	"net/url"

	"github.com/gorilla/websocket"
)

type WSClient struct {
	scheme      string
	addr        string
	path        string
	agent       AgentInterface
	readTimeOut int
	messageType int
}

func NewWSClient(scheme, addr, path string, a AgentInterface, readTimeOutMillisecond int, messageType int) *WSClient {
	return &WSClient{
		scheme:      scheme,
		addr:        addr,
		path:        path,
		agent:       a,
		readTimeOut: readTimeOutMillisecond,
		messageType: messageType,
	}
}

func (this *WSClient) Run() ConnInterface {

	u := url.URL{Scheme: this.scheme, Host: this.addr, Path: this.path}
	klog.Println("connecting to ", u.String())

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		klog.Println("dial:", err)
	}

	var conn ConnInterface
	conn = NewWSConnStruct(c, this.agent, this.readTimeOut, this.messageType)
	err = this.agent.OnConnected(conn)
	if err != nil {
		klog.Println("OnConnected fail:", err)
		return nil
	}
	go conn.Run()
	return conn
}
