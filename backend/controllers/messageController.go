package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
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
	Name_user  string
	id_user    int
	uid        string
}

var clientsList = make(map[int]*Client)

type Manager struct {
	sync.RWMutex
	user     *UserController
	MessageS services.MessageService
	userSer  services.UserService
	Count    map[int]int
}

func NewClient(conn *websocket.Conn, man *Manager, id int, name, uid string) *Client {
	return &Client{
		connection: conn,
		Manager:    man,
		Name_user:  name,
		id_user:    id,
		uid:        uid,
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
		log.Println("WebSocket upgrade error:", err)
		return
	}
	coock, err := r.Cookie("token")
	if err != nil {
		log.Println("Error retrieving cookie:", err)
		return
	}
	remoteAddr := r.RemoteAddr

	// If you only want the IP (without the port)
	// Use net.SplitHostPort to separate IP and port
	host, _, err := net.SplitHostPort(remoteAddr)
	if err != nil {
		fmt.Fprintf(w, "Error getting remote address: %v", err)
		return
	}
	fmt.Println("host",host)
	mes, uuid := m.user.userService.UUiduser(coock.Value)
	if mes.MessageError != "" {
		fmt.Println(mes.MessageError, "jjj")
		return
	}

	client := NewClient(conn, m, uuid.Iduser, uuid.Nickname, coock.Value)
	m.broadcastOnlineUserList("online", client)

	defer func() {
		m.Count[client.id_user]--
		if m.Count[client.id_user] == 0 {
			if clientData, ok := clientsList[client.id_user]; ok && clientData != nil {
				clientData.connection.Close()
				delete(clientsList, client.id_user)
				m.broadcastOnlineUserList("offline", client)
			}
		}
	}()
	m.addClient(client)
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
	mes, mesErr := m.MessageS.GetMessages(Sender, u_ms.Receiver, u_ms.Offste)
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
		m.Username = c.Name_user
		m.Sender = c.id_user
		// m.Firstname = c.Name_user
		m.Sender = c.id_user
		c.Manager.Lock()
		Seen := 0
		if receiverClient, ok := clientsList[m.Receiver]; ok {
			receiverClient.connection.WriteJSON(m)
			Seen = 1
		}
		mg.MessageS.AddMessages(m.Sender, m.Receiver, m.Content, m.CreateAt, Seen)
		c.Manager.Unlock()
	}
}

func (m *Manager) addClient(client *Client) {
	defer m.Unlock()
	m.Lock()
	m.Count[client.id_user]++
	if clientData, ok := clientsList[client.id_user]; ok && clientData != nil && client.uid != clientData.uid {
		clientData.connection.Close()
	}
	if clientData, ok := clientsList[client.id_user]; ok && clientData != nil && client.uid != clientData.uid {
		clientData.connection.Close()
	}
	clientsList[client.id_user] = client
}

func (mu *Manager) broadcastOnlineUserList(typ string, clien *Client) {
	mu.Lock()
	defer mu.Unlock()

	message := models.OnlineUser{
		Type:        typ,
		OnlineUsers: clien.id_user,
		UserName:    clien.Name_user,
	}
	for _, connection := range clientsList {
		connection.connection.WriteJSON(&message)
	}
}
