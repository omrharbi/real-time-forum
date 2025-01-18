package services

import "github.com/gorilla/websocket"

type Client struct {
	Conn     *websocket.Conn
	Messages chan *Message
	Id       string `json:"id_user"`
	UserName string `json:"username"`
}

type Message struct {
	Sender   string `json:"sender"`
	Receiver string `json:"receiver"`
	Content  string `json:"content"`
}

type Hub struct {
	Clients    map[string]*Client
	Register   chan *Client
	Unregister chan *Client
	Broadcast  chan *Message
}

func NewHub() *Hub {
	return &Hub{
		Clients:    make(map[string]*Client),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Broadcast:  make(chan *Message, 5),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.Register:
			h.Clients[client.UserName] = client

		case client := <-h.Unregister:
			if _, ok := h.Clients[client.UserName]; ok {
				close(client.Messages)
				delete(h.Clients, client.UserName)
			}

		case message := <-h.Broadcast:
			if recipient, ok := h.Clients[message.Receiver]; ok {
				recipient.Messages <- message
			}
		}
	}
}
