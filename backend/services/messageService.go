package services

import (
	"real-time-froum/messages"
	"real-time-froum/models"
	"real-time-froum/repo"
)

type MessageService interface {
	AddMessages(m models.Messages) (mss messages.Messages)
	GetMessages(senderID int, receiverID int) ([]string, error)
	DeleteMessages(msgID int) error
}

type MessageServiceImpl struct {
	mess repo.MessageRepository
}

func NewMessageService(ms repo.MessageRepository) MessageService {
	return &MessageServiceImpl{mess: ms}
}

// AddMessages implements MessageService.
func (m *MessageServiceImpl) AddMessages(ms models.Messages) (mss messages.Messages) {
	err := m.mess.AddMessage(ms)
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
func (m *MessageServiceImpl) GetMessages(senderID int, receiverID int) ([]string, error) {
	panic("unimplemented")
}
