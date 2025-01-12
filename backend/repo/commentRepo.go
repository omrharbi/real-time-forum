package repo

import (
	"context"
	"database/sql"

	"real-time-froum/models"
)

type CommentRepository interface {
	insertComment(ctx context.Context, card_id, target_id int) int
	getCommentById(ctx context.Context, id int) *models.Comment_Row
	getAllCommentsbyTargetId(ctx context.Context, target int) []models.Comment_Row_View
}

type commentRepositoryImpl struct {
	db *sql.DB
}
func NewCommentRepository(db *sql.DB) CommentRepository {
	return &commentRepositoryImpl{db: db}
}

// getAllCommentsbyTargetId implements CommentRepository.
func (c *commentRepositoryImpl) getAllCommentsbyTargetId(ctx context.Context, target int) []models.Comment_Row_View {
	list_Comments := make([]models.Comment_Row_View, 0)
	query := `SELECT c.id,c.user_id,c.content,c.created_at,
	u.firstname,u.lastname, (SELECT count(*) FROM comment cm
	 WHERE cm.target_id = c.id) comments,(SELECT count(*) FROM likes l WHERE l.card_id = c.id and l.is_like = 1) likes , (SELECT count(*) FROM likes l WHERE l.card_id = c.id and l.is_like = -1) dislikes
			FROM card c  JOIN comment cm ON c.id = cm.card_id JOIN user u ON c.user_id = u.id
			WHERE cm.target_id = ? 
			GROUP BY c.id ORDER BY c.id DESC;`
	data_Rows := c.db.QueryRowContext(ctx, query, target)
	for data_Rows.Next() {
		Row := models.Comment_Row_View{}
		err := data_Rows.Scan(&Row.Id, &Row.User_Id, &Row.Content, &Row.CreatedAt, &Row.FirstName, &Row.LastName, &Row.Comments, &Row.Likes, &Row.DisLikes)
		if err != nil {
			return nil
		}
		list_Comments = append(list_Comments, Row)
	}
	return list_Comments
}

// getCommentById implements CommentRepository.
func (c *commentRepositoryImpl) getCommentById(ctx context.Context, id int) *models.Comment_Row {
	Row := models.Comment_Row{}
	query := "SELECT * FROM comment WHERE comment.id =?;"
	err := c.db.QueryRowContext(ctx, query, id).Scan(&Row.ID, &Row.Card_Id, &Row.Target_Id)
	if err != nil {
		return nil
	}
	return &Row
}

// insertComment implements CommentRepository.
func (c *commentRepositoryImpl) insertComment(ctx context.Context, card_id int, target_id int) int {
	query := "INSERT INTO comment(card_id,target_id) VALUES(?,?);"
	resl, _ := c.db.ExecContext(ctx, query, card_id, target_id)
	id, err := resl.LastInsertId()
	if err != nil {
		return -1
	}
	return int(id)
}

