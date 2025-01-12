package models

type Comment_Row struct {
	ID        int
	User_Id   int
	Content   string
	CreatedAt string
	Card_Id   int
	Target_Id int
}

type Comment_Row_View struct {
	Id        int
	User_Id   int
	Content   string
	CreatedAt string
	FirstName string
	LastName  string
	Likes     int
	DisLikes  int
	Comments  int
}
