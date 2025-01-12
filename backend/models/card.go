package models

type Card struct {
	Id        int
	User_Id   int
	Content   string
	CreatedAt string
}

type Card_View struct {
	Id        int    `json:"id"`
	User_Id   int    `json:"userid"`
	Content   string `json:"content"`
	CreatedAt string `json:"date"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Likes     int    `json:"likes"`
	DisLikes  int    `json:"dislikes"`
	Comments  int    `json:"comments"`
}
