package repo

import (
	"database/sql"

	"real-time-froum/messages"
	"real-time-froum/models"
)

type MessageRepository interface {
	AddMessage(Sender, Receiver int, Content string) (mss messages.Messages)
	GetMeessage(senderID int, receiverID int) (s []models.Messages, mss messages.Messages)
	DeleteMessage()
}

type MessageRepositoryImpl struct {
	db *sql.DB
}

func NewMessageRepository(db *sql.DB) MessageRepository {
	return &MessageRepositoryImpl{db: db}
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
func (m *MessageRepositoryImpl) GetMeessage(senderID int, receiverID int) (s []models.Messages, mss messages.Messages) {
	qury := `SELECT sender, receiver ,content  FROM messages
		WHERE
            (sender = $1 AND receiver = $2)
            OR
            (sender = $2 AND receiver = $1)

		ORDER BY created_at DESC;`

	row, err := m.db.Query(qury, senderID, receiverID)
	if err != nil {
		mss.MessageError = err.Error()
		return []models.Messages{}, mss
	}
	for row.Next() {
		message := models.Messages{}
		row.Scan(&message.Sender, &message.Receiver, &message.Content)
		s = append(s, message)

	}
	return s, messages.Messages{}
}
