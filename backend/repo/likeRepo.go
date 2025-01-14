package repo

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"

	"real-time-froum/messages"
	"real-time-froum/models"
)

type likesRepository interface {
	InserLike(ctx context.Context, user_id, card_id, is_liked int, UserLiked, Userdisliked bool) (m messages.Messages)
	GetuserLiked(ctx context.Context, card_id int) []models.ResponseUserLikeds
	GetLikes(ctx context.Context, post_id int) (int, int, int)
	DeletLike(ctx context.Context, user_id, card_id int)
	LikeExists(ctx context.Context, user_id, card_id int) bool
}

type likeRepositoryImpl struct {
	db *sql.DB
}

func NewLikesRepository(db *sql.DB) likesRepository {
	return &likeRepositoryImpl{db: db}
}

// GetLikes implements likesRepository.
func (l *likeRepositoryImpl) GetLikes(ctx context.Context, post_id int) (int, int, int) {
	querylike := `SELECT   DISTINCT c.user_id , (SELECT count(*) FROM likes l WHERE ( l.card_id = c.id ) and l.is_like = 1) 
	likes , 
	(SELECT count(*) FROM likes l WHERE ( l.card_id = c.id ) and l.is_like = 0) dislikes FROM post  p
      LEFT JOIN card  c ON p.card_id=c.id    JOIN  likes l WHERE p.card_id = l.card_id AND (l.is_like = 1 or l.is_like = 0 ) AND p.id = ? ` + strconv.Itoa(post_id)
	like := 0
	dislike := 0
	userliked := 0

	err := l.db.QueryRowContext(ctx, querylike).Scan(&userliked, &like, &userliked)
	if err != nil {
		fmt.Println(err)
		like = 0
	}

	return like, dislike, userliked
}

// GetuserLiked implements likesRepository.
func (l *likeRepositoryImpl) GetuserLiked(ctx context.Context, card_id int) []models.ResponseUserLikeds {
	querylike := `SELECT l.UserLiked , l.Userdisliked , u.UUID FROM likes l JOIN card c 
    on l.card_id=c.id JOIN user u ON u.id=l.user_id  WHERE  l.card_id =? `

	likesusers := []models.ResponseUserLikeds{}
	rows, err := l.db.QueryContext(ctx, querylike, card_id)
	if err != nil {
		fmt.Println("Error in likws get user liked", err)
	}
	for rows.Next() {
		likes := models.ResponseUserLikeds{}
		err := rows.Scan(&likes.UserLiked, &likes.UserDisliked, &likes.Uuid)
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
func (l *likeRepositoryImpl) InserLike(ctx context.Context, user_id int, card_id int, is_liked int, UserLiked bool, Userdisliked bool) (m messages.Messages) {
	if l.LikeExists(ctx, user_id, card_id) {
		query := `DELETE FROM likes WHERE user_id = ? AND card_id = ?`
		_, err := l.db.ExecContext(ctx, query, user_id, card_id)
		if err != nil {
			fmt.Println(err.Error())
		}
	}
	query := "INSERT INTO likes(user_id, card_id, is_like, UserLiked, Userdisliked) VALUES(?,?,?,?,?);"
	_, err := l.db.ExecContext(ctx, query, user_id, card_id, is_liked, UserLiked, Userdisliked)
	if err != nil {
		fmt.Println(err.Error())
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
