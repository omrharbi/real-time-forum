package controllers

import (
	"encoding/json"
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
	egress     chan models.Messages
	Name_user  string
	id_user    int
}

var clientsList = make(map[int]*Client)

type Manager struct {
	sync.RWMutex
	user     *UserController
	MessageS services.MessageService
	userSer  services.UserService

	Count map[int]int
}

func NewClient(conn *websocket.Conn, man *Manager, id int, name string) *Client {
	return &Client{
		connection: conn,
		Manager:    man,
		egress:     make(chan models.Messages),
		Name_user:  name,
		id_user:    id,
	}
}

func NewManager(user *UserController, messageS services.MessageService, userSer services.UserService) *Manager {
	return &Manager{
		user:     user,
		MessageS: messageS,
		userSer:  userSer,
		Count:    make(map[int]int),
	}
}

func (m *Manager) ServWs(w http.ResponseWriter, r *http.Request) {
	conn, err := websocketUpgrade.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Err", err)
		return
	}
	coock, _ := r.Cookie("token")
	mes, uuid := m.user.userService.UUiduser(coock.Value)
	if mes.MessageError != "" {
		fmt.Println(mes.MessageError , "jjj")
	}

	m.broadcastOnlineUserList("online", uuid.Iduser)
	client := NewClient(conn, m, uuid.Iduser, uuid.Nickname)

	m.Count[client.id_user]++
	defer func() {
		m.Count[client.id_user]--
		if m.Count[client.id_user] == 0 {
			clientsList[client.id_user].connection.Close()
			delete(clientsList, client.id_user)
			m.broadcastOnlineUserList("offline", client.id_user)
		}
	}()
	go client.WriteMess()
	client.ReadMess(m)
}

func (m *Manager) Read(client *Client) {
	go client.WriteMess()
	client.ReadMess(m)
}

func (m *Manager) HandleGetMessages(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		JsoneResponse(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	u_ms := models.Resiver{}
	des := json.NewDecoder(r.Body)
	des.DisallowUnknownFields()
	err := des.Decode(&u_ms)
	if err != nil {
		JsoneResponse(w, err.Error(), http.StatusNotFound)
		return
	}
	Sender := r.Context().Value("id_user").(int)
	mes, mesErr := m.MessageS.GetMessages(Sender, u_ms.Receiver)
	if mesErr.MessageError != "" {
		JsoneResponse(w, mesErr.MessageError, http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(mes)
}

func (c *Client) ReadMess(mg *Manager) {
	for {
		var m models.Messages
		err := c.connection.ReadJSON(&m)
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Println("error Reading Message", err)
			}
			break
		}
		fmt.Println(m)
		c.Manager.Lock()
		if receiverClient, ok := clientsList[m.Receiver]; ok {
			receiverClient.connection.WriteJSON(m)
		}
		mg.MessageS.AddMessages(m.Sender, m.Receiver, m.Content,m.CreateAt)
		c.Manager.Unlock()

	}
}

func (c *Client) WriteMess() {
	for msg := range c.egress {
		if err := c.connection.WriteJSON(websocket.CloseMessage); err != nil {
			log.Println("Connection Closed ", err)
			return
		}

		if err := c.connection.WriteJSON(msg); err != nil {
			log.Println("Error To Send Message", err)
		}
	}
}

func (m *Manager) addClient(client *Client) {
	m.Lock()
	defer m.Unlock()
	clientsList[client.id_user] = client
}

func (mu *Manager) broadcastOnlineUserList(typ string, id_user int) {
	mu.Lock()
	defer mu.Unlock()

	message := models.OnlineUser{
		Type:        typ,
		OnlineUsers: id_user,
	}
	for _, connection := range clientsList {
		connection.connection.WriteJSON(&message)
	}
}
