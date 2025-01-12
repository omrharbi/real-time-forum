package models
type Like struct {
	ID           int  `json:"id"`
	User_Id      int  `json:"user_id"`
	Card_Id      int  `json:"card_id"`
	Is_Liked     int  `json:"is_liked"`
	UserLiked    bool `json:"userliked"`
	Userdisliked bool `json:"userdisliked"`
}
type DeletLikes struct {
	User_Id int `json:"uuid"`
	Card_Id int `json:"card_id"`
}
type ResponseUserLikeds struct {
	UserLiked    bool
	UserDisliked bool
	Uuid         string
}