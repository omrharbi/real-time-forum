package models

import (
	"github.com/gorilla/websocket"
)

type Messages struct {
	Sender   int    `json:"sender_user"`
	Receiver int    `json:"receiver"`
	UserName string `json:"userName"`
	Content  string `json:"content"`
}

type Client struct {
	Connection *websocket.Conn
	Egress     chan []byte
	Name_user  string
	Id_user    int
}
