package middlewares

import (
	"fmt"
	"net/http"
	"time"

	"real-time-froum/controllers"
	"real-time-froum/models"
	"real-time-froum/services"
)

type Response struct {
	Message string `json:"message"`
	Status  string `json:"status"`
}
type MeddlewireController struct {
	userService services.UserService
}

func NewMeddlewireController(service services.UserService) *MeddlewireController {
	return &MeddlewireController{
		userService: service,
	}
}

func (m MeddlewireController) AuthenticateMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookies, err := r.Cookie("token")
		// user := models.User{}
		if err != nil || cookies == nil {
			if err == http.ErrNoCookie {
				controllers.JsoneResponse(w, "Unauthorized: Cookie not presen", http.StatusUnauthorized)
				return
			}
		}
		if cookies.Value == "" {
			controllers.JsoneResponse(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		messages, expire := m.userService.AuthenticatLogin(r.Context(), cookies.Value)
		if messages.MessageError != "" {
			controllers.JsoneResponse(w, messages.MessageError, http.StatusUnauthorized)
			return
		}
		if !time.Now().Before(expire) {
			u := models.UUID{}
			m.userService.UUiduser(r.Context(), cookies.Value)
			m.userService.LogOut(r.Context(), u)
			fmt.Println("Log out")
			return
		} else {
			next.ServeHTTP(w, r)
		}
	})
}
