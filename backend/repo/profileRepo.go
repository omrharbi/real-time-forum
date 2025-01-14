package repo

import (
	"context"
	"database/sql"
	"strconv"
)

type ProfileRepository interface {
	GetPostsProfile(ctx context.Context, user_id int) string
	GetProfileByLikes(ctx context.Context, user_id int) string
}

type ProfileRepositoryImpl struct {
	db *sql.DB
}

// GetPostsProfile implements ProfileRepository.
func (p *ProfileRepositoryImpl) GetPostsProfile(ctx context.Context, user_id int) string {
	query := `SELECT
	p.card_id AS 'card_id',
	u.id AS 'user_id',
	p.id,
	c.content,
	c.created_at ,
	u.firstname,
	u.lastname,
	u.nickname,
    u.Age,
    u.gender,
    count(cm.id) comments
	FROM post p, card c, user u LEFT  JOIN comment cm
	ON c.id = cm.target_id  WHERE p.card_id=c.id
	AND c.user_id=u.id AND u.id ="` + strconv.Itoa(user_id) + "\" GROUP BY c.id  ORDER BY c.id DESC"
	return query
}

// GetProfileByLikes implements ProfileRepository.
func (p *ProfileRepositoryImpl) GetProfileByLikes(ctx context.Context, user_id int) string {
	query := `SELECT
	p.card_id AS 'card_id',
	u.id AS 'user_id',
	p.id,
	c.content,
	c.created_at ,
	u.firstname,
	u.lastname,
	u.nickname,
    u.Age,
    u.gender,
    count(cm.id) comments
	FROM post p, card c, likes l ,user u LEFT JOIN comment cm
	ON c.id = cm.target_id  WHERE p.card_id=c.id AND l.is_like = 1
	AND c.user_id=u.id AND p.card_id = l.card_id AND l.user_id ="` + strconv.Itoa(user_id) + "\" GROUP BY c.id  ORDER BY c.id DESC"
	return query
}

func NewProfileRepository(db *sql.DB) ProfileRepository {
	return &ProfileRepositoryImpl{db: db}
}
 