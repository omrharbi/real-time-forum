package controllers

import (
	"net/http"
	"time"

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
	card           services.CardsService
}

func NewpostController(service services.PostService, userController *UserController, card services.CardsService) *postController {
	return &postController{
		postService:    service,
		userController: userController,
		card:           card,
	}
}

func (p *postController) HandlePost(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		JsoneResponse(w, "Status Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	id_user := r.Context().Value("id_user").(int)
	post := models.Post{}
	decode := DecodeJson(r)
	err := decode.Decode(&post)
	if err != nil {
		JsoneResponse(w, err.Error(), http.StatusBadRequest)
		return
	}
	post.User_Id = id_user
	ms := p.postService.Add(r.Context(), &post)
	post.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
	if ms.MessageError != "" {
		JsoneResponse(w, ms.MessageError, http.StatusBadRequest)
		return
	}
	Card_View := p.card.GetOneCard(r.Context(), post.Card_Id)
	// card := models.Card_View{}
	JsoneResponse(w, Card_View, http.StatusCreated)
}
