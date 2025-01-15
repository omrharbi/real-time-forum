package services

import (
	"context"

	"real-time-froum/models"
	"real-time-froum/repo"
)

type CardsService interface {
	Add(ctx context.Context, user_id int, Content string) int
	GetCard(ctx context.Context, id int) *models.Card
	GetOneCard(ctx context.Context, id int) *models.Card_View
	GetAllCardsForPages(ctx context.Context, page int, postsPerPage int) ([]models.Card_View, int)
}

type CardsserviceImpl struct {
	catRepo repo.CardRepository
}

func NewcardssService(repo repo.CardRepository) CardsService {
	return &CardsserviceImpl{catRepo: repo}
}

// Add implements cardsService.
func (c *CardsserviceImpl) Add(ctx context.Context, user_id int, Content string) int {
	return c.catRepo.InsertCard(ctx, user_id, Content)
}

// GetAllCardsForPages implements cardsService.
func (c *CardsserviceImpl) GetAllCardsForPages(ctx context.Context, page int, postsPerPage int) ([]models.Card_View, int) {
	return c.catRepo.GetAllCardsForPages(ctx, page, postsPerPage)
}

// GetCard implements cardsService.
func (c *CardsserviceImpl) GetCard(ctx context.Context, id int) *models.Card {
	return c.catRepo.GetCardById(ctx, id)
}

// GetOneCard implements cardsService.
func (c *CardsserviceImpl) GetOneCard(ctx context.Context, id int) *models.Card_View {
	return c.catRepo.GetCard(ctx, id)
}
