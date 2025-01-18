package controllers

import (
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

type ClientList map[*Client]bool

var websocketUpgrade = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type Message struct {
	Id      int    `json:"id_user"`
	Message string `json:"message"`
}

type Client struct {
	connection *websocket.Conn
	Manager    *Manager
	egress     chan []byte
}
type Manager struct {
	Client ClientList
	sync.RWMutex
}

func NewClient(conn *websocket.Conn, man *Manager) *Client {
	return &Client{
		connection: conn,
		Manager:    man,
		egress:     make(chan []byte),
	}
}

func NewManager() *Manager {
	return &Manager{
		Client: make(ClientList),
	}
}

func (m *Manager) ServWs(w http.ResponseWriter, r *http.Request) {
	log.Println("Connected")
	conn, err := websocketUpgrade.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Err", err)
		return
	}
	client := NewClient(conn, m)
	m.addClient(client)
	/// start Messages
	go client.ReadMess()
	go client.WriteMess()
}

func (c *Client) ReadMess() {
	defer func() {
		c.Manager.removeClient(c)
	}()
	for {
		var m Message
		err := c.connection.ReadJSON(&m)
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Println("error Reading Message", err)
			}
			break
		}
		for wsClient := range c.Manager.Client {
			wsClient.egress <- []byte(m.Message)
		}
		fmt.Println(m.Id, "message")
		fmt.Println(m.Message, "message")
	}
}

func (c *Client) WriteMess() {
	defer func() {
		c.Manager.removeClient(c)
	}()
	for {
		select {
		case mess, ok := <-c.egress:
			if !ok {
				if err := c.connection.WriteJSON(websocket.CloseMessage); err != nil {
					log.Println("Connection Closed ", err)
				}
				return
			}

			if err := c.connection.WriteJSON(mess); err != nil {
				log.Println("Error To Send Message", err)
			}
			fmt.Println("Message Sending")
		}
	}
}

func (m *Manager) addClient(client *Client) {
	m.Lock()
	defer m.Unlock()
	m.Client[client] = true // connected client
}

func (m *Manager) removeClient(client *Client) {
	m.Lock()
	defer m.Unlock()
	if _, ok := m.Client[client]; ok {
		client.connection.Close()
		delete(m.Client, client)
	}
}
