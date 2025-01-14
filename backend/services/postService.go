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
	Add(ctx context.Context, p *models.Post) 
	CheckPostErr(w http.ResponseWriter, ps *models.Post)
	GetPosts(ctx context.Context, query string) []models.PostResponde
}

type postServiceImpl struct {
	postRepo     repo.PostRepository
	caredRepo    repo.CardRepository
	categoryRepo repo.CategoryRepository
}

func NewPostService(postRepo repo.PostRepository, caredRepo repo.CardRepository, categoryRepo repo.CategoryRepository) *postServiceImpl {
	return &postServiceImpl{
		postRepo:     postRepo,
		caredRepo:    caredRepo,
		categoryRepo: categoryRepo,
	}
}

// Add implements postService.
func (ps *postServiceImpl) Add(ctx context.Context, p *models.Post) {
	content := html.EscapeString(p.Content)
	cards := ps.caredRepo.InsertCard(ctx, p.User_Id, content)
	p.Card_Id = cards
	id_posr := ps.postRepo.InserPost(ctx, p.Card_Id)
	for _, name := range p.Name_Category {
		err := ps.categoryRepo.PostCategory(ctx, int(id_posr), name) // sp.AddCategory(r.Context(), id, name)
		if err != nil {
			// JsoneResponse(w, r, "Failed to add category", http.StatusBadRequest)
			return
		}
	}
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
