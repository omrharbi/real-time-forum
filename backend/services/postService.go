package services

import (
	"context"
	"html"

	"real-time-froum/messages"
	"real-time-froum/models"
	"real-time-froum/repo"
)

type PostService interface {
	Add(ctx context.Context, p *models.Post) (m messages.Messages)
	GetPosts_Service(ctx context.Context, query string) []models.PostResponde
}

type PstService struct {
	postRepo     repo.PostRepository
	caredRepo    repo.CardRepository
	categoryRepo repo.CategoryRepository
}

func NewPostService(postRepo repo.PostRepository, caredRepo repo.CardRepository, categoryRepo repo.CategoryRepository) *PstService {
	return &PstService{
		postRepo:     postRepo,
		caredRepo:    caredRepo,
		categoryRepo: categoryRepo,
	}
}

// Add implements postService.
func (ps *PstService) Add(ctx context.Context, p *models.Post) (m messages.Messages) {
	content := html.EscapeString(p.Content)
	if content == "" {
		m.MessageError = "Content Is Null"
		return m
	}
	cards := ps.caredRepo.InsertCard(ctx, p.User_Id, content)
	p.Card_Id = cards
	id_posr := ps.postRepo.InserPost(ctx, p.Card_Id)
	for _, name := range p.Name_Category {
		m := ps.categoryRepo.PostCategory(ctx, int(id_posr), name) // sp.AddCategory(r.Context(), id, name)
		if m.MessageError != "" {
			return m
		}
	}
	return messages.Messages{}
}

// GetPosts implements postService.
func (p *PstService) GetPosts_Service(ctx context.Context, query string) []models.PostResponde {
	posts := p.postRepo.GetPosts(ctx, query)
	return posts
}
