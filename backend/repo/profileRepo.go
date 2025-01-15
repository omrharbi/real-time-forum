package repo

import (
	"database/sql"
	"strconv"
)

type ProfileRepository interface {
	GetPostsProfile(user_id int) string
	GetProfileByLikes(user_id int) string
	// GetAllCards(ctx context.Context) []models.Card_View
}

type ProfileRepositoryImpl struct {
	db *sql.DB
}

func NewProfileRepository(db *sql.DB) ProfileRepository {
	return &ProfileRepositoryImpl{db: db}
}

// GetPostsProfile implements ProfileRepository.
func (p *ProfileRepositoryImpl) GetPostsProfile(user_id int) string {
	query := `SELECT
	p.card_id AS 'card_id', 
	c.user_id  AS 'user_id',
	p.id,
	c.content,
	c.created_at ,
	u.firstname, 
	u.lastname,
	u.nickname,
	u.Age,
	u.gender,
    count(cm.id) comments,
	(SELECT count(*) FROM likes l WHERE ( l.post_id =p.id  ) AND l.is_like = 1) as likes,
     (SELECT count(*) FROM likes l WHERE( l.post_id =p.id )AND l.is_like = 0) as dislikes
	FROM post p, card c, user u LEFT  JOIN comment cm
	ON c.id = cm.target_id  WHERE p.card_id=c.id 
	AND c.user_id=u.id AND u.id ="` + strconv.Itoa(user_id) + "\" GROUP BY c.id  ORDER BY c.id DESC"
	return query
}

// GetProfileByLikes implements ProfileRepository.
func (p *ProfileRepositoryImpl) GetProfileByLikes(user_id int) string {
	query := `SELECT
		p.card_id AS 'card_id', 
		c.user_id  AS 'user_id',
		p.id,
		c.content,
		c.created_at ,
		u.firstname, 
		u.lastname,
		u.nickname,
		u.Age,
		u.gender,
		count(cm.id) comments,
		(SELECT count(*) FROM likes l WHERE ( l.post_id =p.id  ) AND l.is_like = 1) as likes,
     	(SELECT count(*) FROM likes l WHERE( l.post_id =p.id )AND l.is_like = 0) as dislikes
		 
		FROM post p, card c, likes l ,user u LEFT JOIN comment cm
		ON c.id = cm.target_id  WHERE p.card_id=c.id AND l.is_like = 1
		AND c.user_id=u.id AND p.card_id = l.card_id AND l.user_id ="` + strconv.Itoa(user_id) + "\" GROUP BY c.id  ORDER BY c.id DESC"
	return query
}

// // getAllCards implements cardRepository.
// func (p *ProfileRepositoryImpl) GetAllCards(ctx context.Context) []models.Card_View {
// 	list_Cards := make([]models.Card_View, 0)
// 	query := `SELECT c.id,c.user_id,c.content,c.created_at,u.firstname,u.lastname,
// 	 count(cm.id) comments,(SELECT count(*) FROM likes l WHERE l.card_id = c.id and l.is_like = 1)
// 	  likes , (SELECT count(*) FROM likes l WHERE l.card_id = c.id and l.is_like = -1) dislikes
// 			FROM card c JOIN post p on c.id = p.card_id LEFT JOIN comment cm
// 			ON c.id = cm.target_id JOIN user u ON c.user_id = u.id
// 			GROUP BY c.id  ORDER BY c.id DESC  `
// 	data_Rows, err := p.db.QueryContext(ctx, query)
// 	if err != nil {
// 		fmt.Println("Error", err)
// 	}
// 	for data_Rows.Next() {
// 		Row := models.Card_View{}
// 		fmt.Println(Row)
// 		err := data_Rows.Scan(&Row.Id, &Row.User_Id, &Row.Age, &Row.Nickname, &Row.Gender, &Row.Content, &Row.CreatedAt, &Row.FirstName, &Row.LastName, &Row.Comments, &Row.Likes, &Row.DisLikes)
// 		if err != nil {
// 			fmt.Println(err)
// 			return nil
// 		}
// 		list_Cards = append(list_Cards, Row)
// 	}
// 	return list_Cards
// }
