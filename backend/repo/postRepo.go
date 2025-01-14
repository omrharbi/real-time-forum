package repo

import (
	"context"
	"database/sql"
	"fmt"

	"real-time-froum/models"
)

type PostRepository interface {
	InserPost(ctx context.Context, card_id int) int64
	GetPosts(ctx context.Context, query string) []models.PostResponde
}
type postRepositoryImpl struct {
	db *sql.DB
}

func NewPostRepository(db *sql.DB) PostRepository {
	return &postRepositoryImpl{db: db}
}

// GetPosts implements PostRepository.
func (p *postRepositoryImpl) GetPosts(ctx context.Context, query string) []models.PostResponde {
	rows, err := p.db.QueryContext(ctx, query)
	if err != nil {
		return nil
	}
	defer rows.Close()
	var posts []models.PostResponde
	for rows.Next() {
		var post models.PostResponde
		err := rows.Scan(
			&post.Card_Id,
			&post.UserID,
			&post.Post_Id,
			&post.Content,
			&post.CreatedAt,
			&post.FirstName,
			&post.LastName,
			&post.Nickname,
			&post.Age,
			&post.Gender,
			&post.Comments,
		)
		if err != nil {
			fmt.Println("er", err)
			return nil
		}
		// likes, dislikes, userliked, Userdisliked := p..GetLikes(post.Post_Id)
		// post.Likes = likes
		// post.Dislikes = dislikes
		// post.UserLiked = userliked
		// post.Userdisliked = Userdisliked
		posts = append(posts, post)
	}
	return posts
}

// inserPost implements PostRepository.
func (p *postRepositoryImpl) InserPost(ctx context.Context, card_id int) int64 {
	query := "INSERT INTO post(card_id) VALUES(?);"
	row, err := p.db.ExecContext(ctx, query, card_id)
	if err != nil {
		fmt.Println("error to insert")
	}
	id, err := row.LastInsertId()
	if err != nil {
		fmt.Println("Error to get id ")
	}
	return id
}
