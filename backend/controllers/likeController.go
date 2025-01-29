package controllers

import (
	"encoding/json"
	"net/http"

	"real-time-froum/models"
	"real-time-froum/services"
)

type likesController struct {
	likes services.LikeServer
	user  *UserController
}

func NewLikesController(likes services.LikeServer, user *UserController) *likesController {
	return &likesController{
		likes: likes,
		user:  user,
	}
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
	liked.User_Id = r.Context().Value("id_user").(int)
	dislike := l.likes.ChecklikesUser(r.Context(), liked)
	json.NewEncoder(w).Encode(dislike)
}

func (l *likesController) HandleAddLike(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode("Status Method Not Allowed")
		return
	}
	id_user := r.Context().Value("id_user").(int)
	like := models.Like{}
	decode := DecodeJson(r)
	err := decode.Decode(&like)
	if err != nil {
		JsoneResponse(w, err.Error(), http.StatusBadRequest)
		return
	}
	like.User_Id = id_user
	m := l.likes.Addlikes(r.Context(), like)
	if m.MessageError != "" {
		JsoneResponse(w, m.MessageError, http.StatusBadRequest)
		return
	}
	JsoneResponse(w, m.MessageSucc, http.StatusCreated)
}

func (l *likesController) HandleDeletLike(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode("Status Method Not Allowed")
		return
	}
	like := models.DeletLikes{}
	id_user := r.Context().Value("id_user").(int)
	decode := DecodeJson(r)
	err := decode.Decode(&like)
	if err != nil {
		JsoneResponse(w, "Error of the Decode likes", http.StatusBadRequest)
		return
	}
	like.User_Id = id_user
	l.likes.DeletLike(r.Context(), like)
	JsoneResponse(w, "DELETED Like", http.StatusCreated)
}
