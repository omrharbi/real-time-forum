package controllers

import (
	"fmt"
	"log"
	"net/http"
	"sync"

	"real-time-froum/models"
	"real-time-froum/services"

	"github.com/gorilla/websocket"
)

var websocketUpgrade = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

type Client struct {
	connection *websocket.Conn
	Manager    *Manager
	egress     chan []byte
	Name_user  string
	id_user    int
}

type Manager struct {
	Clients map[int]*Client // Map user ID to their Client object
	sync.RWMutex
	user     *UserController
	MessageS services.MessageService
}

func NewClient(conn *websocket.Conn, man *Manager, id int, name string) *Client {
	return &Client{
		connection: conn,
		Manager:    man,
		egress:     make(chan []byte),
		Name_user:  name,
		id_user:    id,
	}
}

func NewManager(user *UserController, messageS services.MessageService) *Manager {
	return &Manager{
		Clients:  make(map[int]*Client),
		user:     user,
		MessageS: messageS,
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
	coock, err := r.Cookie("token")
	if err != nil {
		fmt.Println("Err", err)
		return
	}
	mes, uuid := m.user.userService.UUiduser(coock.Value)
	if mes.MessageError != "" {
		fmt.Println(mes.MessageError)
	}

	client := NewClient(conn, m, uuid.Iduser, uuid.Nickname)
	m.addClient(client)
	go client.WriteMess()

	go client.ReadMess(m)
}

func (c *Client) ReadMess(ms *Manager) {
	defer func() {
		c.Manager.removeClient(c)
	}()
	for {
		var m models.Messages
		err := c.connection.ReadJSON(&m)
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Println("error Reading Message", err)
			}
			break
		}
		c.Manager.Lock()
		if receiverClient, ok := c.Manager.Clients[m.Receiver]; ok {
			fmt.Println(receiverClient, m.Sender)
			message := fmt.Sprintf("From %s: %s", c.Name_user, m.Content)
			receiverClient.egress <- []byte(message)
			ms.MessageS.AddMessages(m)
		} else {
			log.Printf("Recipient with ID %d not connected\n", m.Receiver)
		}
		c.Manager.Unlock()
		fmt.Println("Message from", c.Name_user, "to", m.Receiver, ":", m.Content)

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
	m.Clients[client.id_user] = client // Add the client to the map using their user ID
	log.Printf("Client added: %s (ID: %d)\n", client.Name_user, client.id_user)
}

func (m *Manager) removeClient(client *Client) {
	m.Lock()
	defer m.Unlock()
	if _, ok := m.Clients[client.id_user]; ok {
		client.connection.Close()
		delete(m.Clients, client.id_user) // Remove the client from the map
		log.Printf("Client removed: %s (ID: %d)\n", client.Name_user, client.id_user)
	}
}
