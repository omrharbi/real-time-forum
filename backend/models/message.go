package models

type Messages struct {
	Type      string `json:"type"`
	Sender    int    `json:"sender"`
	Receiver  int    `json:"receiver"`
	Username  string `json:"username"`
	Firstname string `json:"firstname"`
	Content   string `json:"content"`
	CreateAt  string `json:"createAt"`
}
type Resiver struct {
	Receiver int `json:"receiver"`
}
type OnlineUser struct {
	Type        string `json:"type"`
	OnlineUsers int    `json:"online_users"`
}
