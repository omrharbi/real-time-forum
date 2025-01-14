package repo

import (
	"context"
	"database/sql"
	"fmt"
)

type CategoryRepository interface {
	PostCategory(ctx context.Context, postId int, category string) error
	GetCategoryId(ctx context.Context, category string) (int, error)
	GetPostsByCategor(  categoryName string) string
}

type CategoryRepositoryImpl struct {
	db *sql.DB
}

func NewCategoryRepository(db *sql.DB) CategoryRepository {
	return &CategoryRepositoryImpl{db: db}
}

// GetPostsByCategoryId implements CategoryRepository.

func (c *CategoryRepositoryImpl) GetPostsByCategor(  categoryName string) string {
	query := `
	SELECT c.id,
    c.user_id,
    p.id,
    c.content,
    c.created_at,
    u.firstname,
    u.lastname,count(cm.id) comments
			FROM card c JOIN post p on c.id = p.card_id LEFT JOIN comment cm
			ON c.id = cm.target_id JOIN user u ON c.user_id = u.id
            JOIN post_category pc on pc.post_id=p.id 
            JOIN category cat on cat.id=pc.category_id
            WHERE cat.name = "` + categoryName + "\" GROUP BY c.id  ORDER BY c.id DESC"
	return query
}

// getCategoryId implements CategoryRepository.
func (c *CategoryRepositoryImpl) GetCategoryId(ctx context.Context, category string) (int, error) {
	categoryId := 0
	query := "SELECT id FROM category WHERE name = ?"

	err := c.db.QueryRowContext(ctx, query, category).Scan(&categoryId)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	return categoryId, nil
}

// postCategory implements CategoryRepository.
func (c *CategoryRepositoryImpl) PostCategory(ctx context.Context, postId int, category string) error {
	categoryId, err := c.GetCategoryId(ctx, category)
	if err != nil {
		return err
	}
	query := "INSERT INTO post_category (post_id, category_id) VALUES(?,?)"
	_, err = c.db.ExecContext(ctx, query, postId, categoryId)
	if err != nil {
		return err
	}
	return nil
}

// getCategoryId implements CategoryRepository.
