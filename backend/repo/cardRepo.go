package repo

import (
	"context"
	"database/sql"
	"fmt"

	"real-time-froum/models"
)

type CardRepository interface {
	GetAllCardsForPages(ctx context.Context, page int, postsPerPage int) ([]models.Card_View, int)
	GetCard(ctx context.Context, targetID int) *models.Card_View
	GetCardById(ctx context.Context, id int) *models.Card
	InsertCard(ctx context.Context, user_id int, content string) (int, error)
}

type cardRepositoryImpl struct {
	db *sql.DB
}

func NewcardRepository(db *sql.DB) CardRepository {
	return &cardRepositoryImpl{db: db}
}

// getAllCardsForPages implements cardRepository.
func (c *cardRepositoryImpl) GetAllCardsForPages(ctx context.Context, page int, postsPerPage int) ([]models.Card_View, int) {
	list_Cards := make([]models.Card_View, 0)

	countQuery := `SELECT COUNT(DISTINCT c.id) 
                   FROM card c 
                   JOIN post p on c.id = p.card_id 
                   JOIN user u ON c.user_id = u.id`

	countRows, err := c.db.QueryContext(ctx, countQuery)
	if err != nil {
		fmt.Println("Error in get all one", err)
	}
	var totalPosts int
	if countRows.Next() {
		err := countRows.Scan(&totalPosts)
		if err != nil {
			return nil, 0
		}
	}
	defer countRows.Close()

	offset := (page - 1) * postsPerPage

	query := `SELECT c.id, u.UUID, c.content, c.created_at, u.firstname, u.lastname, u.username,u.Age,u.gender,
              count(cm.id) comments,
              (SELECT count(*) FROM likes l WHERE ( l.card_id =p.id  ) AND l.is_like = 1) as likes,
        		(SELECT count(*) FROM likes l WHERE( l.card_id =p.id )AND l.is_like = 0) as dislikes
              FROM card c 
              JOIN post p on c.id = p.card_id 
              LEFT JOIN comment cm ON c.id = cm.target_id 
              JOIN user u ON c.user_id = u.id
              GROUP BY c.id  
              ORDER BY c.id DESC
              LIMIT ? OFFSET ?`

	data_Rows, err := c.db.QueryContext(ctx, query, postsPerPage, offset)
	if err != nil {
		fmt.Println("Error in get all", err)
		return []models.Card_View{}, 0
	}
	defer data_Rows.Close()

	for data_Rows.Next() {
		Row := models.Card_View{}
		err := data_Rows.Scan(&Row.Id, &Row.User_uuid, &Row.Content, &Row.CreatedAt,
			&Row.FirstName, &Row.LastName, &Row.Nickname, &Row.Age, &Row.Gender, &Row.Comments,
			&Row.Likes, &Row.DisLikes)
		if err != nil {
			return nil, 0
		}
		list_Cards = append(list_Cards, Row)
	}

	return list_Cards, totalPosts
}

// getCard implements cardRepository.

func (c *cardRepositoryImpl) GetCard(ctx context.Context, targetID int) *models.Card_View {
	query := `SELECT c.id, u.UUID, c.content, c.created_at, u.firstname, u.lastname, u.username,u.Age,u.gender,
       (SELECT count(*) FROM comment cm WHERE cm.target_id = c.id) as comments,
        (SELECT count(*) FROM likes l WHERE ( l.card_id =p.card_id  ) AND l.is_like = 1) as likes,
        (SELECT count(*) FROM likes l WHERE( l.card_id =p.card_id )AND l.is_like = 0) as dislikes
       	FROM card c LEFT JOIN comment cm  on c.id=cm.card_id LEFT  JOIN post p on p.card_id=c.id
		JOIN user u ON c.user_id = u.id
		WHERE c.id =?;`
	Row := &models.Card_View{}
	err := c.db.QueryRowContext(ctx, query, targetID).Scan(&Row.Id, &Row.User_uuid, &Row.Content, &Row.CreatedAt, &Row.FirstName, &Row.LastName, &Row.Nickname, &Row.Age, &Row.Gender, &Row.Comments, &Row.Likes, &Row.DisLikes)
	if err != nil {
		return &models.Card_View{}
	}

	return Row
}

// getCardById implements cardRepository.
func (c *cardRepositoryImpl) GetCardById(ctx context.Context, id int) *models.Card {
	query := "SELECT * FROM card WHERE card.id =?;"
	myCard_Row := &models.Card{}
	err := c.db.QueryRowContext(ctx, query, id).Scan(&id, &myCard_Row.User_Id, &myCard_Row.Content, &myCard_Row.CreatedAt)

	if err != nil {
		return nil
	} else {
		return myCard_Row
	}
}

// insertCard implements cardRepository.
func (c *cardRepositoryImpl) InsertCard(ctx context.Context, user_id int, content string) (int, error) {
	query := "INSERT INTO card(user_id,content) VALUES(?,?)"
	resl, _ := c.db.ExecContext(ctx, query, user_id, content)
	id, err := resl.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), nil
}
