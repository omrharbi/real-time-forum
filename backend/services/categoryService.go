package services

import (
	"context"

	"real-time-froum/models"
	"real-time-froum/repo"
)

type CategoryService interface {
	AddCategory(ctx context.Context, post_id int, category string) error
	GetPostsByCategoryId(ctx context.Context, ategoryName string) []models.PostResponde
}

type CategoryserviceImpl struct {
	catRepo  repo.CategoryRepository
	postRepo repo.PostRepository
}

func NewcategorysService(repo repo.CategoryRepository, postRepo repo.PostRepository) CategoryService {
	return &CategoryserviceImpl{catRepo: repo, postRepo: postRepo}
}

// AddCategory implements CategoryService.
func (c *CategoryserviceImpl) AddCategory(ctx context.Context, post_id int, category string) error {
	err := c.catRepo.PostCategory(ctx, post_id, category)
	if err != nil {
		return err
	}
	return nil
}

// GetPostsByCategoryId implements CategoryService.
func (c *CategoryserviceImpl) GetPostsByCategoryId(ctx context.Context, categoryName string) []models.PostResponde {
	query := c.catRepo.GetPostsByCategor(categoryName)
	return c.postRepo.GetPosts(ctx, query)
}
