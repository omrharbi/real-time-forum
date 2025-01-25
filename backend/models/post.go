package models

type Post struct {
	ID            int      `json:"id"`
	User_Id       int      `json:"user_id"`
	Title         string   `json:"title"`
	Content       string   `json:"content"`
	Name_Category []string `json:"name"`
	CreatedAt     string   `json:"createdat"`
	Card_Id       int      `json:"card_id"`
}

type PostResponde struct {
	Card_Id      int    `json:"id"`
	Post_Id      int    `json:"post_id"`
	User_uuid    string    `json:"user_uuid"`
	FirstName    string `json:"firstName"`
	LastName     string `json:"lastName"`
	Nickname     string `json:"nickname"`
	Age          int    `json:"age"`
	Gender       string `json:"gender"`
	Content      string `json:"content"`
	Likes        int    `json:"likes"`
	DisLikes     int    `json:"dislikes"`
	Userdisliked int    `json:"userdisliked"`
	Comments     string `json:"comments"`
	CreatedAt    string `json:"date"`
}
