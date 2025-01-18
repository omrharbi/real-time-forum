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
	Name_user  string
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

func NewManager(user *UserController) *Manager {
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
	fmt.Println("Add", conn.RemoteAddr())
	// coock, err := r.Cookie("token")
	// if err != nil {
	// 	fmt.Println("Err", err)
	// 	return
	// }
	// mes, uuid := m.user.userService.UUiduser(coock.Value)
	// if mes.MessageError != "" {
	// 	fmt.Println(mes.MessageError)
	// }
	client := NewClient(conn, m)
	client.Name_user = "omar"
	fmt.Println(client)
	m.addClient(client)
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
			if wsClient.connection != c.connection {
				wsClient.egress <- []byte(m.Message)
			}
		}

	}
}

func (c *Client) WriteMess() {
	defer func() {
		c.Manager.removeClient(c)
	}()
	for msg := range c.egress {

		if err := c.connection.WriteJSON(websocket.CloseMessage); err != nil {
			log.Println("Connection Closed ", err)
			return
		}

		if err := c.connection.WriteJSON(msg); err != nil {
			log.Println("Error To Send Message", err)
		}
		fmt.Println("Message Sending")

	}
}

func (m *Manager) addClient(client *Client) {
	m.Lock()
	defer m.Unlock()
	m.Client[client] = true // connected client
	fmt.Println(m.Client)
}

func (m *Manager) removeClient(client *Client) {
	m.Lock()
	defer m.Unlock()
	if _, ok := m.Client[client]; ok {
		client.connection.Close()
		delete(m.Client, client)
	}
}
