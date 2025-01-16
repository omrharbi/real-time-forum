package services

import (
	"context"
	"html"
	"strings"

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
	if strings.TrimSpace(content) == "" {
		m.MessageError = "Content Is Empty"
		return m
	}
	if len(p.Name_Category) == 0 {
		m.MessageError = "Your Category  Is Empty"
		return m
	}

	if checkdeblicat(p.Name_Category) {
		m.MessageError = "Duplicate category: The category already exists"
		return
	}
	if len(p.Content) > 1000 {
		m.MessageError = "Your content is long"
		return
	}
	if p.Content == "" {
		m.MessageError = "Your content is emty"
		return
	}
	for _, n := range p.Name_Category {
		if !checkGategory(n) {
			m.MessageError = "Your category is incorrect"
			return
		}
	}

	cards, err := ps.caredRepo.InsertCard(ctx, p.User_Id, content)
	if err != nil {
		m.MessageError = err.Error()
		return
	}
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
