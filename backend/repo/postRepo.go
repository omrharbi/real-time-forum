package repo

import (
	"context"
	"database/sql"
	"fmt"
)

type PostRepository interface {
	inserPost(ctx context.Context, title string, card_id int) int64
}

type postRepositoryImpl struct {
	db *sql.DB
}

func NewPostRepository(db *sql.DB) PostRepository {
	return &postRepositoryImpl{db: db}
}

// inserPost implements PostRepository.
func (p *postRepositoryImpl) inserPost(ctx context.Context, title string, card_id int) int64 {
	query := "INSERT INTO post(title, card_id) VALUES(?,?);"
	row, err := p.db.ExecContext(ctx, query, title, card_id)
	if err != nil {
		fmt.Println("error to insert")
	}
	id, err := row.LastInsertId()
	if err != nil {
		fmt.Println("Error to get id ")
	}
	return id
}
