package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"real-time-froum/models"
	"real-time-froum/services"
)

type UserController struct {
	userService services.UserService
	ctx         context.Context
}

func NewUserController(service services.UserService, ctx context.Context) *UserController {
	return &UserController{
		userService: service,
		ctx:         ctx,
	}
}

func (uc *UserController) HandleRegister(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		JsoneResponse(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	defer r.Body.Close()
	user := models.User{}
	decode := DecodeJson(r)
	decode.DisallowUnknownFields()
	err := decode.Decode(&user)
	if err != nil {
		JsoneResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(uc.ctx, 2*time.Second)
	defer cancel()
	fmt.Println(user)
	timeex := time.Now().Add(5 * time.Hour).UTC()
	userRegiseter, message, uuid := uc.userService.Register(ctx, timeex, &user)
	fmt.Println(message.MessageError)
	if message.MessageError != "" {
		JsoneResponse(w, message.MessageError, http.StatusBadRequest)
		return
	}

	SetCookie(w, "token", uuid, timeex)
	JsoneResponse(w, userRegiseter, http.StatusOK)
}

func (uc *UserController) HandleLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		JsoneResponse(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	defer r.Body.Close()
	var user models.Login
	decode := DecodeJson(r)
	err := decode.Decode(&user)
	if err != nil {
		JsoneResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	timeex := time.Now().Add(5 * time.Hour).UTC()
	loged, message, uuid := uc.userService.Authentication(uc.ctx, timeex, &user)

	if message.MessageError != "" {
		JsoneResponse(w, message.MessageError, http.StatusBadRequest)
		return
	}

	SetCookie(w, "token", uuid.String(), timeex)
	JsoneResponse(w, loged, http.StatusOK)
}

func (uc *UserController) HandleLogOut(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		JsoneResponse(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	defer r.Body.Close()
	var logout models.Logout
	decode := DecodeJson(r)
	decode.DisallowUnknownFields()
	err := decode.Decode(&logout)
	if err != nil {
		JsoneResponse(w, "Invalid request format", http.StatusBadRequest)
		return
	}
	logout.Id = int64(uc.GetUserId(r))

	ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
	defer cancel()
	message, uuid := uc.userService.UUiduser(logout.Uuid)
	if message.MessageError != "" {
		JsoneResponse(w, "Missing or invalid Uuid", http.StatusBadRequest)
		return
	}
	message = uc.userService.LogOut(ctx, uuid)
	if message.MessageError != "" {
		JsoneResponse(w, message.MessageError, http.StatusBadRequest)
		return
	}
	clearCookies(w)
	w.WriteHeader(http.StatusOK)
}

func SetCookie(w http.ResponseWriter, name string, value string, time time.Time) {
	user := http.Cookie{
		Name:    name,
		Value:   value,
		Expires: time,
		Path:    "/",
	}
	http.SetCookie(w, &user)
}

func (uc *UserController) GetUserId(r *http.Request) int {
	cookie, err := r.Cookie("token")
	if err != nil {
		return 0
	}
	// ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
	// defer cancel()
	m, uuid := uc.userService.UUiduser(cookie.Value)
	if m.MessageError != "" {
		fmt.Println(m.MessageError)
	}
	fmt.Println(uuid, "uuid ")
	return uuid.Iduser
}

func clearCookies(w http.ResponseWriter) {
	SetCookie(w, "token", "", time.Now())
	SetCookie(w, "user_id", "", time.Now())
}

func (uc *UserController) HandleIsLogged(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		JsoneResponse(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	cookies, err := r.Cookie("token")
	if err != nil {
		JsoneResponse(w, err.Error(), http.StatusUnauthorized)
		fmt.Println(err)
		return
	}
	is, expire, _ := uc.userService.CheckAuth(r.Context(), cookies.Value)
	if !time.Now().Before(expire) {
		u := models.UUID{}
		uc.userService.UUiduser(cookies.Value)
		uc.userService.LogOut(r.Context(), u)
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Println("Log out")
		return
	} else {
		if is {
			return
		}
		w.WriteHeader(http.StatusUnauthorized)
	}
}

func (uc *UserController) HandleUserConnected(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		JsoneResponse(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	id := r.Context().Value("id_user").(int)
	id_usr := uc.userService.UserConnect(id)

	// for i := 0; i < len(id_usr); i++ {
	// 	id := id_usr[i].Iduser
	// 	_, ok := clientsList[id]
	// 	if ok {
	// 		id_usr[i].Status = "online"
	// 	}else {
	// 		id_usr[i].Status = "offline"
	// 	}
	// }
	json.NewEncoder(w).Encode(id_usr)
}
