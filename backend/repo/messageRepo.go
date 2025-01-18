package repo

import "database/sql"

type MessageRepository interface {
	AddMessage()
	GetMeessage()
	DeleteMessage()
}

type MessageRepositoryImpl struct {
	db *sql.DB
}

// AddMessage implements MessageRepository.
func (m *MessageRepositoryImpl) AddMessage() {
	panic("unimplemented")
}

// DeleteMessage implements MessageRepository.
func (m *MessageRepositoryImpl) DeleteMessage() {
	panic("unimplemented")
}

// GetMeessage implements MessageRepository.
func (m *MessageRepositoryImpl) GetMeessage() {
	panic("unimplemented")
}

func NewMessageRepository(db *sql.DB) MessageRepository {
	return &MessageRepositoryImpl{db: db}
}
