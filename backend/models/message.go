package models

type Messages struct {
	Type     string `json:"type"`
	Sender   int    `json:"sender"`
	Receiver int    `json:"receiver"`
	Content  string `json:"content"`
}

type OnlineUser struct {
	Type        string `json:"type"`
	OnlineUsers int    `json:"online_users"`
}
