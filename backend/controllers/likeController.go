package controllers

import (
	"encoding/json"
	"net/http"

	"real-time-froum/models"
	"real-time-froum/services"
)

type likesController struct {
	likes services.LikeServer
}

func NewLikesController(likes services.LikeServer) *likesController {
	return &likesController{likes: likes}
}

func (l *likesController) LikesCheckedHandle(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	if r.Method != http.MethodPost {
		JsoneResponse(w, "Method Not Allowd", http.StatusMethodNotAllowed)
		return
	}
	liked := models.Like{}
	decode := DecodeJson(r)
	err := decode.Decode(&liked)
	if err != nil {
		JsoneResponse(w, err.Error(), http.StatusBadRequest)
		return
	}
	dislike := l.likes.ChecklikesUser(r.Context(), liked)
	json.NewEncoder(w).Encode(dislike)
}
