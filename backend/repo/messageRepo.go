package repo

import (
	"database/sql"

	"real-time-froum/messages"
	"real-time-froum/models"
)

type MessageRepository interface {
	AddMessage(m models.Messages) (mss messages.Messages)
	GetMeessage()
	DeleteMessage()
}

type MessageRepositoryImpl struct {
	db *sql.DB
}

// AddMessage implements MessageRepository.
func (m *MessageRepositoryImpl) AddMessage(ms models.Messages) (mss messages.Messages) {
	qury := "INSERT INTO messages (sender,receiver,content) VALUES(?,?,?)"
	_, err := m.db.Exec(qury, ms.Sender, ms.Receiver, ms.Content)
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
