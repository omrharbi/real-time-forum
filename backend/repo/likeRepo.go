package repo

import (
	"context"
	"database/sql"
	"fmt"

	"real-time-froum/messages"
	"real-time-froum/models"
)

type LikesRepository interface {
	InserLike(ctx context.Context, user_id, card_id int, is_liked bool) (m messages.Messages)
	GetuserLiked(ctx context.Context, card_id int, userid int) []models.ResponseUserLikeds
	// GetLikes(ctx context.Context, card_id int) int
	DeletLike(ctx context.Context, user_id, card_id int)
	LikeExists(ctx context.Context, user_id, card_id int) bool
}

type likeRepositoryImpl struct {
	db *sql.DB
}

func NewLikesRepository(db *sql.DB) LikesRepository {
	return &likeRepositoryImpl{db: db}
}

// GetuserLiked implements likesRepository. // mybe is no working
func (l *likeRepositoryImpl) GetuserLiked(ctx context.Context, card_id int, userid int) []models.ResponseUserLikeds {
	querylike := `SELECT l.is_like=1 , l.is_like=0 , u.UUID,u.id as user_id  FROM likes l JOIN card c 
    on l.card_id=c.id JOIN user u ON u.id=l.user_id  WHERE  l.card_id =? and l.user_id=? `

	likesusers := []models.ResponseUserLikeds{}
	rows, err := l.db.QueryContext(ctx, querylike, card_id, userid)
	if err != nil {
		fmt.Println("Error in likws get user liked", err)
	}
	for rows.Next() {
		likes := models.ResponseUserLikeds{}
		err := rows.Scan(&likes.UserLiked, &likes.UserDisliked, &likes.Uuid, &likes.Id_user)
		if err != nil {
			fmt.Println(err)
		}
		likesusers = append(likesusers, likes)
	}
	return likesusers
}

// deletLike implements likesRepository.
func (l *likeRepositoryImpl) DeletLike(ctx context.Context, user_id int, card_id int) {
	query := "DELETE FROM likes WHERE user_id=? AND card_id=?"
	_, err := l.db.ExecContext(ctx, query, user_id, card_id)
	if err != nil {
		fmt.Println(err.Error(), "test")
	}
}

// inserLike implements likesRepository.
func (l *likeRepositoryImpl) InserLike(ctx context.Context, user_id int, card_id int, is_liked bool) (m messages.Messages) {
	if l.LikeExists(ctx, user_id, card_id) {
		l.DeletLike(ctx, user_id, card_id)
	}
	query := "INSERT INTO likes(user_id, card_id, is_like) VALUES(?,?,?);"
	_, err := l.db.ExecContext(ctx, query, user_id, card_id, is_liked)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	m.MessageSucc = "is liked"
	return m
}

// likeExists implements likesRepository.
func (l *likeRepositoryImpl) LikeExists(ctx context.Context, user_id int, card_id int) bool {
	var exists bool
	query := "SELECT EXISTS (select is_like from likes where user_id = ? AND card_id = ?)"
	err := l.db.QueryRowContext(ctx, query, user_id, card_id).Scan(&exists)
	if err != nil {
		fmt.Println("Error exist Like", err)
	}
	return exists
}
