package models

type Messages struct {
	Type     string `json:"sender"`
	Sender   int    `json:"sender"`
	Receiver int    `json:"receiver"`
	Content  string `json:"content"`
}
