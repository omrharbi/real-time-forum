package controllers

import (
	"fmt"
	"log"
	"net/http"

	"real-time-froum/services"

	"github.com/gorilla/websocket"
)

type HandlerHub struct {
	hub *services.Hub
}

func NewHubController(hub *services.Hub) *HandlerHub {
	return &HandlerHub{
		hub: hub,
	}
}

var Upgrade = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		// orgin:=r.Header.Get("Origin")
		// return orgin == "http://localhost:3333"
		return true
	},
}

func (h *HandlerHub) Messages(w http.ResponseWriter, r *http.Request) {
	conn, err := Upgrade.Upgrade(w, r, nil)
	if err != nil {
		log.Println("err", err)
	}
	defer conn.Close()
	var ms *services.Message
	// cl.Conn = conn

	err = conn.ReadJSON(&ms)
	if err != nil {
		log.Println("Error reading JSON:", err)
		return
	}
	cl := &services.Clients{
		Conn:     conn,
		Messages: make(chan *services.Message, 10),
		Id:       ms.Sender,
		UserName: ms.UserName,
	}

	h.hub.Register <- cl
	h.hub.Broadcast <- ms
	go cl.WriteMessage()
	go cl.ReadMessage(h.hub)

	fmt.Println(cl)
}
