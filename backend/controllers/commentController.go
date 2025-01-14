package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

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
		JsoneResponse(res, req, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	id, err := strconv.Atoi(req.FormValue("target_id"))
	if err != nil {
		JsoneResponse(res, req, "Status Bad Request", http.StatusBadRequest)
		return
	}
	ctx, cancel := context.WithTimeout(req.Context(), 2*time.Second)
	defer cancel()
	comments := cn.commentService.GetAllCommentsbyTarget(ctx, id)
	if len(comments) == 0 {
		return
	}
	if comments == nil {
		JsoneResponse(res, req, "Status Not Found", http.StatusNotFound)
		return
	}
	json.NewEncoder(res).Encode(comments)
}

func (cn *CommentController) Handler_AddComment(res http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()
	if req.Method != http.MethodPost {
		JsoneResponse(res, req, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	statusCode := cn.addComment(req)
	if statusCode == http.StatusBadRequest {
		JsoneResponse(res, req, "comment Infos are wrongs!! ", http.StatusBadRequest)
		return
	}
	if statusCode == http.StatusOK {
		JsoneResponse(res, req, "comment added succesfuly", http.StatusCreated)
		return
	}
}

func (cn *CommentController) addComment(req *http.Request) int {
	iduser := cn.userController.GetUserId(req)
	comment := models.Comment{}
	comment.User_Id = iduser
	if comment.User_Id == 0 {
		fmt.Println("error")
		return -1
	}

	decoder := DecodeJson(req)
	err := decoder.Decode(&comment)
	if err != nil || (comment.Content == "") {
		return http.StatusBadRequest
	}
	if comment.Content == "" || len(comment.Content) > 1000 {
		return http.StatusBadRequest
	}
	cn.commentService.AddComment(req.Context(), &comment)
	return http.StatusOK
}
