package services

import (
	"real-time-froum/messages"
	"real-time-froum/models"
	"real-time-froum/repo"
)

type MessageService interface {
	AddMessages(Sender, Receiver int, Content string) (mss messages.Messages)
	GetMessages(senderID int, receiverID int) (s []models.Messages, mss messages.Messages)
	DeleteMessages(msgID int) error
}

type MessageServiceImpl struct {
	repo repo.MessageRepository
}

func NewMessageService(ms repo.MessageRepository) MessageService {
	return &MessageServiceImpl{repo: ms}
}

// AddMessages implements MessageService.
func (m *MessageServiceImpl) AddMessages(Sender, Receiver int, Content string) (mss messages.Messages) {
	err := m.repo.AddMessage(Sender, Receiver, Content)
	if err.MessageError != "" {
		return err
	}
	return messages.Messages{}
}

// DeleteMessages implements MessageService.
func (m *MessageServiceImpl) DeleteMessages(msgID int) error {
	panic("unimplemented")
}

// GetMessages implements MessageService.
func (m *MessageServiceImpl) GetMessages(senderID int, receiverID int) ([]models.Messages, messages.Messages) {
	mess, err := m.repo.GetMeessage(senderID, receiverID)
	if err.MessageError != "" {
		return []models.Messages{}, err
	}
	return mess, messages.Messages{}
}
