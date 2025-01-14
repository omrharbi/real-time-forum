package services

import (
	"context"
	"fmt"

	"real-time-froum/models"
	"real-time-froum/repo"
)

type cardsService interface {
	Add(ctx context.Context, user_id int, Content string) int
	GetCard(ctx context.Context, id int) *models.Card
	GetOneCard(ctx context.Context, id int) models.Card_View
	GetAllCardsForPages(ctx context.Context, page int, postsPerPage int) ([]models.Card_View, int)
}

type CardsserviceImpl struct {
	catRepo repo.CardRepository
}

func NewcardssService(repo repo.CardRepository) cardsService {
	return &CardsserviceImpl{catRepo: repo}
}

// Add implements cardsService.
func (c *CardsserviceImpl) Add(ctx context.Context, user_id int, Content string) int {
	fmt.Println(user_id, Content)
	return 0
}

// GetAllCardsForPages implements cardsService.
func (c *CardsserviceImpl) GetAllCardsForPages(ctx context.Context, page int, postsPerPage int) ([]models.Card_View, int) {
	panic("unimplemented")
}

// GetCard implements cardsService.
func (c *CardsserviceImpl) GetCard(ctx context.Context, id int) *models.Card {
	panic("unimplemented")
}

// GetOneCard implements cardsService.
func (c *CardsserviceImpl) GetOneCard(ctx context.Context, id int) models.Card_View {
	panic("unimplemented")
}
