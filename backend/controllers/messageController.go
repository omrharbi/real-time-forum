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
	egress     chan []byte
	Name_user  string
	id_user    int
}

type Manager struct {
	Clients map[int]*Client // Map user ID to their Client object
	sync.RWMutex
	user     *UserController
	MessageS services.MessageService
	userSer  services.UserService
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

func NewManager(user *UserController, messageS services.MessageService, userSer services.UserService) *Manager {
	return &Manager{
		Clients:  make(map[int]*Client),
		user:     user,
		MessageS: messageS,
		userSer:  userSer,
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
	m.broadcastOnlineUserList(client)
	m.addClient(client)
	go client.WriteMess()

	go client.ReadMess(m)
}

func (m *Manager) HandleGetMessages(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		JsoneResponse(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	u_ms := models.Messages{}
	des := json.NewDecoder(r.Body)
	des.DisallowUnknownFields()
	err := des.Decode(&u_ms)
	if err != nil {
		fmt.Println(err)
	}
	mes, mesErr := m.MessageS.GetMessages(u_ms.Sender, u_ms.Receiver)
	if mesErr.MessageError != "" {
		JsoneResponse(w, mesErr.MessageError, http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(mes)
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
			message := fmt.Sprintf("From %s: %s", c.Name_user, m.Content)
			receiverClient.egress <- []byte(message)
			ms.MessageS.AddMessages(m.Sender, m.Receiver, m.Content)
		} else {

			err := ms.userSer.UpdateStatus("online", m.Sender)
			fmt.Println(err, "Error")
			log.Printf("Recipient with ID %d not connected\n %v %v %v", m.Receiver, m.Type, "offline", c.id_user, c.Name_user)
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
	if existingClient, ok := m.Clients[client.id_user]; ok {
		existingClient.connection.Close()
		delete(m.Clients, client.id_user)
	}
	err := m.userSer.UpdateStatus("online", client.id_user)
	fmt.Println(err, "Error")
	m.Clients[client.id_user] = client
	log.Printf("Client added: %s (ID: %d)\n", client.Name_user, client.id_user)
}

func (m *Manager) removeClient(client *Client) {
	m.Lock()
	defer m.Unlock()
	if _, ok := m.Clients[client.id_user]; ok {
		client.connection.Close()
		delete(m.Clients, client.id_user)
		log.Printf("Client removed: %s (ID: %d)\n", client.Name_user, client.id_user)
	}
}

func (mu *Manager) broadcastOnlineUserList(client *Client) {
	mu.Lock()
	defer mu.Unlock()
	connectedIDs := make([]int, 0, len(mu.Clients))
	for id := range mu.Clients {
		connectedIDs = append(connectedIDs, id)
	}
	fmt.Println(connectedIDs)

	// Broadcast to all connections
	for userID, connection := range mu.Clients {
		// Marshal the filtered list into JSON format and send it to the current connection
		var m models.Messages
		err := client.connection.ReadJSON(&m)
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Println("error Reading Message", err)
			}
			break
		}

		if err != nil {
			// if there was an error it means that the user is disconnected
			connection.connection.Close()
			delete(mu.Clients, userID)
		}
	}
}

// func broadcastOnlineUserList() {
//     mu.Lock()
//     defer mu.Unlock()
//     connectedIDs := make([]int, 0, len(ConnectedUsers))
//     for id := range ConnectedUsers {
//         connectedIDs = append(connectedIDs, id)
//     }

//     // Broadcast to all connections
//     for userID, connection := range ConnectedUsers {
//         // Marshal the filtered list into JSON format and send it to the current connection
//         data := map[string]interface{}{
//             "type":  "users-status",
//             "users": connectedIDs,
//         }
//         message, err := json.Marshal(data)
//         if err != nil {
//             log.Printf("Error marshalling user list for user %d: %v\n", userID, err)
//             continue
//         }

//         err = connection.WriteMessage(websocket.TextMessage, message)
//         if err != nil {
//             // if there was an error it means that the user is disconnected
//             connection.Close()
//             delete(ConnectedUsers, userID)
//         }
//     }
// }
