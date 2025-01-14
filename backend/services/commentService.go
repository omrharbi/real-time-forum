package services

import (
	"context"
	"html"

	"real-time-froum/models"
	"real-time-froum/repo"
)

type CommentService interface {
	AddComment(ctx context.Context, cm *models.Comment)
	GetComment(ctx context.Context, id int) *models.Comment
	GetAllCommentsbyTarget(ctx context.Context, target int) []models.Comment_View
}

type commentServiceImpl struct {
	CommentRepo repo.CommentRepository
	caredRepo   repo.CardRepository
}



func NewCommentService(repo repo.CommentRepository, caredRepo repo.CardRepository) CommentService {
	return &commentServiceImpl{CommentRepo: repo, caredRepo: caredRepo}
}

// GetComment implements CommentService.
func (c *commentServiceImpl) GetComment(ctx context.Context, id int) *models.Comment {
	panic("unimplemented")
}
// AddComment implements CommentService.
func (c *commentServiceImpl) AddComment(ctx context.Context, cm *models.Comment) {
	content := html.EscapeString(cm.Content)
	cards := c.caredRepo.InsertCard(ctx, cm.User_Id, content)
	cm.Card_Id = cards
	cm.ID = c.CommentRepo.InsertComment(ctx, cm.Card_Id, cm.Target_Id)
}

// GetAllCommentsbyTarget implements CommentService.
func (c *commentServiceImpl) GetAllCommentsbyTarget(ctx context.Context, target int) []models.Comment_View {
	list_Comments := make([]models.Comment_View, 0)
	list := c.CommentRepo.GetAllCommentsbyTargetId(ctx, target)
	size := len(list)
	if size == 0 {
		return nil
	}
	for index := 0; index < size; index++ {
		list_Comments = append(list_Comments, list[index])
	}
	return list_Comments
}

// GetComment implements CommentService.
// func (c *commentServiceImpl) GetComment(ctx context.Context, id int) *models.Comment {
// 	data_Row := c.CommentRepo.GetCommentById(ctx, id)
// 	if data_Row == nil {
// 		return nil
// 	}
// 	card := c.caredRepo.GetCard(ctx, data_Row.Card_Id)
// 	if card == nil {
// 		return nil
// 	}
// 	// newComment := &Comment{
// 	// 	ID:        data_Row.ID,
// 	// 	Card_Id:   data_Row.Card_Id,
// 	// 	Target_Id: data_Row.Target_Id,
// 	// 	Content:   card.Content,
// 	// 	User_Id:   card.User_Id,
// 	// 	CreatedAt: card.CreatedAt,
// 	// }
// 	return newComment
// }
