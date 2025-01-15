package models

type Comment struct {
	ID        int    `json:"id"`
	User_Id   int    `json:"user_id"`
	Content   string `json:"content"`
	CreatedAt string `json:"createdat"`
	Card_Id   int    `json:"card_id"`
	Target_Id int    `json:"target_id"`
}

type Comment_View struct {
	Id_comment int    `json:"id_comment"`
	Id         int    `json:"card_id"`
	User_uuid  int    `json:"User_uuid"`
	Content    string `json:"content"`
	CreatedAt  string `json:"date"`
	FirstName  string `json:"firstName"`
	LastName   string `json:"lastName"`
	Likes      int    `json:"likes"`
	DisLikes   int    `json:"dislikes"`
	Comments   int    `json:"comments"`
	Nickname   string `json:"nickname"`
	Age        int    `json:"age"`
	Gender     string `json:"gender"`
}
