package services

import (
	"log"
	"sync"

	"github.com/gorilla/websocket"
)

type Clients struct {
	Conn     *websocket.Conn
	Messages chan *Message
	Id       int    `json:"id_user"`
	UserName string `json:"username"`
	sync.RWMutex
}

type Message struct {
	Sender   int    `json:"sender_user"`
	Receiver int    `json:"receiver"`
	UserName string `json:"userName"`
	Content  string `json:"content"`
}

type Hub struct {
	Clients    map[int]*Clients
	Register   chan *Clients
	Unregister chan *Clients
	Broadcast  chan *Message
}

func NewHub() *Hub {
	return &Hub{
		Clients:    make(map[int]*Clients),
		Register:   make(chan *Clients),
		Unregister: make(chan *Clients),
		Broadcast:  make(chan *Message, 5),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.Register:

			h.Clients[client.Id] = client

		case client := <-h.Unregister:
			if _, ok := h.Clients[client.Id]; ok {
				close(client.Messages)
				delete(h.Clients, client.Id)
			}

		case message := <-h.Broadcast:
			if recipient, ok := h.Clients[message.Receiver]; ok {
				recipient.Messages <- message
			}
		}
	}
}

func (c *Clients) WriteMessage() {
	defer func() {
		c.Conn.Close()
	}()
	for {
		mess, ok := <-c.Messages
		if !ok {
			return
		}
		c.Conn.WriteJSON(mess)
	}
}

func (c *Clients) ReadMessage(h *Hub) {
	defer func() {
		h.Unregister <- c
		c.Conn.Close()
	}()

	for {

		_, m, err := c.Conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Println("error Reading Message", err)
			}
			break
		}
		msg := &Message{
			Content:  string(m),
			UserName: c.UserName,
			Sender:   c.Id,
		}
		h.Broadcast <- msg
	}
}
