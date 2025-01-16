package controllers

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"real-time-froum/messages"
	"real-time-froum/models"
	"real-time-froum/services"
)

type CommentController struct {
	commentService services.CommentService
	userController *UserController
}

func NewCommentController(comment services.CommentService, userController *UserController) *CommentController {
	return &CommentController{
		commentService: comment,
		userController: userController,
	}
}

func (cn *CommentController) Handel_GetCommet(res http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()
	if req.Method != http.MethodGet {
		JsoneResponse(res, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	id, err := strconv.Atoi(req.FormValue("target_id"))
	if err != nil {
		JsoneResponse(res, "Status Bad Request", http.StatusBadRequest)
		return
	}
	ctx, cancel := context.WithTimeout(req.Context(), 2*time.Second)
	defer cancel()
	comments := cn.commentService.GetAllCommentsbyTarget(ctx, id)
	if len(comments) == 0 {
		return
	}
	if comments == nil {
		JsoneResponse(res, "Status Not Found", http.StatusNotFound)
		return
	}
	json.NewEncoder(res).Encode(comments)
}

func (cn *CommentController) Handler_AddComment(res http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()
	if req.Method != http.MethodPost {
		JsoneResponse(res, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	ms := cn.addComment(req)
	if ms.MessageError != "" {
		JsoneResponse(res, "comment Infos are wrongs!! ", http.StatusBadRequest)
		return
	}

	JsoneResponse(res, "comment added succesfuly", http.StatusCreated)
}

func (cn *CommentController) addComment(req *http.Request) (m messages.Messages) {
	iduser := cn.userController.GetUserId(req)
	comment := models.Comment{}
	comment.User_Id = iduser
	if comment.User_Id == 0 {
		m.MessageError = "Invalid User ID"
		return m
	}

	decoder := DecodeJson(req)
	err := decoder.Decode(&comment)
	if err != nil {
		m.MessageError = err.Error()
		return m
	}

	ms := cn.commentService.AddComment(req.Context(), &comment)
	if ms.MessageError != "" {
		return ms
	}
	return messages.Messages{}
}
