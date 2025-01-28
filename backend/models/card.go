package models

type Card struct {
	Id        int
	User_Id   int
	Content   string
	CreatedAt string
}

type Card_View struct {
	Id         int    `json:"id"`
	User_uuid  string `json:"user_uuid"`
	Content    string `json:"content"`
	CreatedAt  string `json:"date"`
	FirstName  string `json:"firstName"`
	Nickname   string `json:"nickname"`
	Age        int    `json:"age"`
	Gender     string `json:"gender"`
	LastName   string `json:"lastName"`
	Likes      int    `json:"likes"`
	DisLikes   int    `json:"dislikes"`
	UserLiked  any    `json:"userLiked"`
	Comments   int    `json:"comments"`
	Categories string `json:"categories"`
}

type PaginatedResponse struct {
	Posts        []Card_View `json:"posts"`
	TotalPosts   int         `json:"totalPosts"`
	TotalPages   int         `json:"totalPages"`
	CurrentPage  int         `json:"currentPage"`
	PostsPerPage int         `json:"postsPerPage"`
}
