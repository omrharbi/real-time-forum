package repo

import (
	"database/sql"
	"fmt"

	"real-time-froum/messages"
	"real-time-froum/models"
)

type MessageRepository interface {
	AddMessage(Sender, Receiver int, Content, CreateAt string, seen int) (mss messages.Messages)
	GetMeessage(senderID int, receiverID int, offset int) (s []models.Messages, mss messages.Messages)
	DeleteMessage()
}

type MessageRepositoryImpl struct {
	db *sql.DB
}

func NewMessageRepository(db *sql.DB) MessageRepository {
	return &MessageRepositoryImpl{db: db}
}

// AddMessage implements MessageRepository.m
func (m *MessageRepositoryImpl) AddMessage(Sender, Receiver int, Content, CreateAt string, seen int) (mss messages.Messages) {
	qury := "INSERT INTO messages (sender,receiver,created_at,content , seen ) VALUES(?,? ,?,? , ?)"
	_, err := m.db.Exec(qury, Sender, Receiver, CreateAt, Content, seen)
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
func (m *MessageRepositoryImpl) GetMeessage(senderID int, receiverID int, offste int) (s []models.Messages, mss messages.Messages) {
	if offste == 0 {
		query := `
		UPDATE messages
		SET seen = 1
		WHERE receiver = $1;
		`
		m.db.Exec(query, senderID)
	}
	qury := `SELECT m.sender, m.receiver ,m.content,m.created_at,u.firstname,u.username  FROM messages m 
		LEFT JOIN user u on m.sender = u.id
		WHERE
            (sender = $1 AND receiver = $2)
              OR
            (sender = $2 AND receiver = $1)

		ORDER BY created_at DESC
		LIMIT $3 OFFSET $4;
`

	row, err := m.db.Query(qury, senderID, receiverID, 30, offste)
	if err != nil {
		fmt.Println(err)
		mss.MessageError = err.Error()
		return []models.Messages{}, mss
	}
	for row.Next() {
		message := models.Messages{}
		row.Scan(&message.Sender, &message.Receiver, &message.Content, &message.CreateAt, &message.Firstname, &message.Username)
		s = append(s, message)

	}
	return s, messages.Messages{}
}
