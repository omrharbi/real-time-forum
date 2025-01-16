package services

import (
	"context"

	"real-time-froum/messages"
	"real-time-froum/models"
	"real-time-froum/repo"
)

type LikeServer interface {
	Addlikes(ctx context.Context, like models.Like) messages.Messages
	DeletLike(ctx context.Context, like models.Like)
	ChecklikesUser(ctx context.Context, like models.Like) []models.ResponseUserLikeds
}

type likesServiceImpl struct {
	like repo.LikesRepository
}

func NewLikesServer(like repo.LikesRepository) LikeServer {
	return &likesServiceImpl{like: like}
}

func (l *likesServiceImpl) Addlikes(ctx context.Context, like models.Like) messages.Messages {
	m := l.like.InserLike(ctx, like.User_Id, like.Card_Id, like.Is_Liked)
	return m
}

func (l *likesServiceImpl) DeletLike(ctx context.Context, like models.Like) {
	l.like.DeletLike(ctx, like.User_Id, like.Card_Id)
}

func (l *likesServiceImpl) ChecklikesUser(ctx context.Context, like models.Like) []models.ResponseUserLikeds {
	likes := l.like.GetuserLiked(ctx, like.Card_Id)
	// fmt.Println(likes)
	return likes
}
