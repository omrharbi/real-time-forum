package services

import (
	"real-time-froum/models"
	"real-time-froum/repo"

	"github.com/gorilla/websocket"
)

type MessageService interface {
	AddMessages(msg string, senderID int, receiverID int) error
	GetMessages(senderID int, receiverID int) ([]string, error)
	DeleteMessages(msgID int) error
}

type MessageServiceImpl struct {
	mess repo.MessageRepository
}

func NewClient(conn *websocket.Conn, id int, name string) *models.Client {
	return &models.Client{
		Connection: conn,
		Egress:     make(chan []byte),
		Name_user:  name,
		Id_user:    id,
	}
}
