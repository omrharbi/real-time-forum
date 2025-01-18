package services

import "github.com/gorilla/websocket"

type Clients struct {
	Conn     *websocket.Conn
	Messages chan *Messages
	Id       string `json:"id_user"`
	UserName string `json:"username"`
	RoomID   string `json:"roomid"`
}

type Room struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Clients map[string]*Clients
}
type Messages struct {
	UserName string `json:"username"`
	Message  string `json:"message"`
	RoomID   string `json:"roomid"`
}

type Hub struct {
	Room       map[string]*Room
	Register   chan *Clients
	Unregister chan *Clients
	Broadcast  chan *Messages
}

func NewHub() *Hub {
	return &Hub{
		Room:       make(map[string]*Room),
		Register:   make(chan *Clients),
		Unregister: make(chan *Clients),
		Broadcast:  make(chan *Messages, 5),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case cl := <-h.Register:
			if _, ok := h.Room[cl.RoomID]; ok {
				r := h.Room[cl.Id]
				if _, ok := r.Clients[cl.Id]; ok {
					r.Clients[cl.Id] = cl
				}
			}
		case cl := <-h.Unregister:
			if _, ok := h.Room[cl.RoomID]; ok {
				if _, ok := h.Room[cl.RoomID].Clients[cl.Id]; ok {

					if len(h.Room[cl.RoomID].Clients) != 0 {
						h.Broadcast <- &Messages{
							Message:  "user Left the chat",
							RoomID:   cl.RoomID,
							UserName: cl.UserName,
						}
					}

					delete(h.Room[cl.RoomID].Clients, cl.Id)
					close(cl.Messages)
				}
			}
		case m := <-h.Broadcast:
			if _, ok := h.Room[m.RoomID]; ok {
				for _, cl := range h.Room[m.RoomID].Clients {
					cl.Messages <- m
				}
			}
		}
	}
}
