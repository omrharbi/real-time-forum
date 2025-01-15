package services

import (
	"context"

	"real-time-froum/models"
	"real-time-froum/repo"
)

type ProfileService interface {
	GetPostsProfile(ctx context.Context, user_id int) []models.PostResponde
	GetProfileByLikes(ctx context.Context, user_id int) []models.PostResponde
}

type ProfileserviceImpl struct {
	porRepo  repo.ProfileRepository
	postRepo repo.PostRepository
}

func NewProfilesService(repo repo.ProfileRepository, postRepo repo.PostRepository) ProfileService {
	return &ProfileserviceImpl{porRepo: repo, postRepo: postRepo}
}

// GetPostsProfile implements ProfileService.
func (p *ProfileserviceImpl) GetPostsProfile(ctx context.Context, user_id int) []models.PostResponde {
	query := p.porRepo.GetPostsProfile(user_id)
	return p.postRepo.GetPosts(ctx, query)
}

// GetProfileByLikes implements ProfileService.
func (p *ProfileserviceImpl) GetProfileByLikes(ctx context.Context, user_id int) []models.PostResponde {
	query := p.porRepo.GetProfileByLikes(user_id)
	return p.postRepo.GetPosts(ctx, query)
}
