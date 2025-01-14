package services

import (
	"context"
	"encoding/json"
	"html"
	"net/http"

	"real-time-froum/models"
	"real-time-froum/repo"
)

type PostService interface {
	Add(ctx context.Context, p *models.Post) int
	CheckPostErr(w http.ResponseWriter, ps *models.Post)
	GetPosts(ctx context.Context, query string) []models.PostResponde
}

type postServiceImpl struct {
	postRepo repo.PostRepository
}

// Add implements postService.
func (ps *postServiceImpl) Add(ctx context.Context, p *models.Post) int {
	content := html.EscapeString(p.Content)
	cards := &CardsserviceImpl{}

	idpost := cards.Add(ctx, p.User_Id, content)

	p.Card_Id = idpost
	id_posr := ps.postRepo.InserPost(ctx, p.Card_Id)
	return int(id_posr)
}

// CheckPostErr implements postService.
func (p *postServiceImpl) CheckPostErr(w http.ResponseWriter, ps *models.Post) {
	if ps.Content == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Invalid input")
	}
}

// GetPosts implements postService.
func (p *postServiceImpl) GetPosts(ctx context.Context, query string) []models.PostResponde {
	posts := p.postRepo.GetPosts(ctx, query)
	return posts
}

func NewpostService(repo repo.PostRepository) PostService {
	return &postServiceImpl{postRepo: repo}
}
