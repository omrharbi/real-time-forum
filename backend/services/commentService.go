package services

import (
	"context"
	"fmt"
	"html"
	"strings"

	"real-time-froum/messages"
	"real-time-froum/models"
	"real-time-froum/repo"
)

type CommentService interface {
	AddComment(ctx context.Context, cm *models.Comment) (m messages.Messages)
	// GetComment(ctx context.Context, id int) *models.Comment
	GetAllCommentsbyTarget(ctx context.Context, target int) []models.Comment_View
}

type commentServiceImpl struct {
	CommentRepo repo.CommentRepository
	caredRepo   repo.CardRepository
}

func NewCommentService(repo repo.CommentRepository, caredRepo repo.CardRepository) CommentService {
	return &commentServiceImpl{CommentRepo: repo, caredRepo: caredRepo}
}

// AddComment implements CommentService.
func (c *commentServiceImpl) AddComment(ctx context.Context, cm *models.Comment) (m messages.Messages) {
	content := html.EscapeString(cm.Content)

	if strings.TrimSpace(content) == "" {
		m.MessageError = "Content is empty or only contains whitespace"
		return m
	}

	if len(content) > 1000 {
		m.MessageError = "Content exceeds the maximum allowed length of 1000 characters"
		return m
	}

	cards, err := c.caredRepo.InsertCard(ctx, cm.User_Id, content)
	if err != nil {
		fmt.Println(err)
		m.MessageError = err.Error()
		return m
	}
	cm.Card_Id = cards
	m = c.CommentRepo.InsertComment(ctx, cm.Card_Id, cm.Target_Id)
	if m.MessageError != "" {
		return m
	}
	return messages.Messages{}
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
