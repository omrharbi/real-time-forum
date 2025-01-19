package controllers

import (
	"fmt"
	"log"
	"net/http"

	"real-time-froum/services"

	"github.com/gorilla/websocket"
)

type HandlerHub struct {
	hub            *services.Hub
	userController *UserController
}

func NewHubController(hub *services.Hub, user *UserController) *HandlerHub {
	return &HandlerHub{
		hub:            hub,
		userController: user,
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

	uuid := h.userController.GetUserId(r)

	cl := &services.Clients{
		Conn:     conn,
		Messages: make(chan *services.Message, 10),
		Id:       uuid,
	}
	go h.hub.Run(uuid)
	// h.hub.Register <- cl
	// h.hub.Broadcast <- ms
	go cl.WriteMessage()
	go cl.ReadMessage(h.hub)

	fmt.Println(cl)
}
