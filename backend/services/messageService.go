package services

import "real-time-froum/repo"

type MessageService interface {
	AddMessages()
	GetMeessages()
	DeleteMessages()
}

type MessageServiceImpl struct {
	mess repo.MessageRepository
}

func NewMessageService(mess repo.MessageRepository) MessageService {
	return &MessageServiceImpl{mess: mess}
}

// AddMessages implements MessageService.
func (m *MessageServiceImpl) AddMessages() {
	panic("unimplemented")
}

// DeleteMessages implements MessageService.
func (m *MessageServiceImpl) DeleteMessages() {
	panic("unimplemented")
}

// GetMeessages implements MessageService.
func (m *MessageServiceImpl) GetMeessages() {
	panic("unimplemented")
}
