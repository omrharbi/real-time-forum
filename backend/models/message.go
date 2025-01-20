package models

type Messages struct {
	Sender   int    `json:"sender"`
	Receiver int    `json:"receiver"`
	Content  string `json:"content"`
}
