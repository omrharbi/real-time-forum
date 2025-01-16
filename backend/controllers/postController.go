package controllers

import (
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
	postService    services.PostService
	userController *UserController
}

func NewpostController(service services.PostService, userController *UserController) *postController {
	return &postController{
		postService:    service,
		userController: userController,
	}
}

func (p *postController) HandlePost(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		JsoneResponse(w, "Status Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	id_user := p.userController.GetUserId(r)
	post := &models.Post{}
	decode := DecodeJson(r)
	err := decode.Decode(&post)
	if err != nil {
		JsoneResponse(w, err.Error(), http.StatusBadRequest)
		return
	}
	post.User_Id = id_user
	ms := p.postService.Add(r.Context(), post)

	if ms.MessageError != "" {
		JsoneResponse(w, ms.MessageError, http.StatusBadRequest)
		return
	}

	JsoneResponse(w, "create post Seccessfuly", http.StatusCreated)
}
