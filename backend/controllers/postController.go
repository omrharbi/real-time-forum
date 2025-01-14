package controllers

import (
	"fmt"
	"net/http"

	"real-time-froum/models"
	"real-time-froum/services"
)

type PaginatedResponse struct {
	Posts        []models.Card_View `json:"posts"`
	TotalPosts   int                `json:"totalPosts"`
	TotalPages   int                `json:"totalPages"`
	CurrentPage  int                `json:"currentPage"`
	PostsPerPage int                `json:"postsPerPage"`
}

type postController struct {
	postService services.PostService
}

func NewpostController(service services.PostService) *postController {
	return &postController{
		postService: service,
	}
}

func (p *postController) HandlePost(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		JsoneResponse(w, r, "Status Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	// user:={}
	user := UserController{}
	id_user := user.GetUserId(r)
	post := &models.Post{}
	decode := DecodeJson(r)
	err := decode.Decode(&post)
	if err != nil {
		JsoneResponse(w, r, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Println(post)
	if checkdeblicat(post.Name_Category) {
		JsoneResponse(w, r, "Duplicate category: The category already exists", http.StatusConflict)
		return
	}
	if len(post.Content) > 1000 {
		JsoneResponse(w, r, "Your content is long", http.StatusBadRequest)
		return
	}
	if post.Content == "" {
		JsoneResponse(w, r, "Your content is emty", http.StatusBadRequest)
		return
	}
	for _, n := range post.Name_Category {
		if !checkGategory(n) {
			JsoneResponse(w, r, "Your category is incorrect", http.StatusBadRequest)
			return
		}
	}
	post.User_Id = id_user
	p.postService.CheckPostErr(w, post)

	id := p.postService.Add(r.Context(), post)
	fmt.Println(id)
	// for _, name := range post.Name_Category {
	// 	sp := &services.CategoryserviceImpl{}
	// 	err := sp.AddCategory(r.Context(), id, name)
	// 	if err != nil {
	// 		JsoneResponse(w, r, "Failed to add category", http.StatusBadRequest)
	// 		return
	// 	}
	// }
	JsoneResponse(w, r, "create post Seccessfuly", http.StatusCreated)
}

func checkdeblicat(cat []string) bool {
	for i := 0; i < len(cat); i++ {
		for j := i + 1; j < len(cat); j++ {
			if cat[i] == cat[j] {
				return true
			}
		}
	}
	return false
}

func checkGategory(name string) bool {
	cate := []string{
		"General",
		"Technology",
		"Sports",
		"Entertainment",
		"Science",
		"Health",
		"Food",
		"Travel",
		"Fashion",
		"Art",
		"Music",
	}
	for _, v := range cate {
		if v == name {
			return true
		}
	}
	return false
}
