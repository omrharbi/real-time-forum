package repo

import (
	"database/sql"

	"real-time-froum/messages"
)

type MessageRepository interface {
	AddMessage(Sender, Receiver int, Content string) (mss messages.Messages)
	GetMeessage()
	DeleteMessage()
}

type MessageRepositoryImpl struct {
	db *sql.DB
}

// AddMessage implements MessageRepository.m
func (m *MessageRepositoryImpl) AddMessage(Sender, Receiver int, Content string) (mss messages.Messages) {
	qury := "INSERT INTO messages (sender,receiver,content) VALUES(?,?,?)"
	_, err := m.db.Exec(qury, Sender, Receiver, Content)
	if err != nil {
		mss.MessageError = err.Error()
		return mss
	}
	return messages.Messages{}
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
