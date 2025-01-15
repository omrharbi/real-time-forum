package repo

import (
	"context"
	"database/sql"
	"fmt"

	"real-time-froum/models"
)

type CommentRepository interface {
	InsertComment(ctx context.Context, card_id, target_id int) int
	// GetCommentById(ctx context.Context, id int) *models.Comment
	GetAllCommentsbyTargetId(ctx context.Context, target int) []models.Comment_View
}

type commentRepositoryImpl struct {
	db *sql.DB
}

func NewCommentRepository(db *sql.DB) CommentRepository {
	return &commentRepositoryImpl{db: db}
}

// getAllCommentsbyTargetId implements CommentRepository.
func (c *commentRepositoryImpl) GetAllCommentsbyTargetId(ctx context.Context, target int) []models.Comment_View {
	list_Comments := make([]models.Comment_View, 0)
	query := `SELECT cm.id as id_comment,  c.id as card_id,  u.UUID,  c.content, c.created_at, u.firstname, u.lastname, u.nickname,u.Age,u.gender, 
    (SELECT count(*) FROM comment cm WHERE cm.target_id = c.id) comments,
  	(SELECT count(*) FROM likes l WHERE ( l.card_id = c.id ) and l.is_like = 1) likes , 
	(SELECT count(*) FROM likes l WHERE ( l.card_id = c.id ) and l.is_like = 0) dislikes
			FROM card c  JOIN comment cm ON c.id = cm.card_id JOIN user u ON c.user_id = u.id
			WHERE cm.target_id = ?
			GROUP BY c.id ORDER BY c.id DESC;`
	data_Rows, err := c.db.QueryContext(ctx, query, target)
	if err != nil {
		fmt.Println("Error in comment", err)
	}
	for data_Rows.Next() {
		Row := models.Comment_View{}
		err := data_Rows.Scan(&Row.Id_comment, &Row.Id, &Row.User_uuid, &Row.Content, &Row.CreatedAt, &Row.FirstName, &Row.LastName, &Row.Nickname, &Row.Age, &Row.Gender, &Row.Comments, &Row.Likes, &Row.DisLikes)
		if err != nil {
			fmt.Println(err, "GetAllCommentsbyTargetId")
			return nil
		}
		list_Comments = append(list_Comments, Row)
	}
	return list_Comments
}

// insertComment implements CommentRepository.
func (c *commentRepositoryImpl) InsertComment(ctx context.Context, card_id int, target_id int) int {
	query := "INSERT INTO comment(card_id,target_id) VALUES(?,?);"
	resl, _ := c.db.ExecContext(ctx, query, card_id, target_id)
	id, err := resl.LastInsertId()
	if err != nil {
		return -1
	}
	return int(id)
}
