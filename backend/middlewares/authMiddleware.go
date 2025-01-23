package middlewares

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"real-time-froum/controllers"
	"real-time-froum/models"
	"real-time-froum/services"
)

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

		messages, expire, id_user := m.userService.AuthenticatLogin(cookies.Value)
		if messages.MessageError != "" {
			controllers.JsoneResponse(w, messages.MessageError, http.StatusUnauthorized)
			return
		}

		r = r.WithContext(context.WithValue(r.Context(), "id_user", id_user))

		if !time.Now().Before(expire) {
			u := models.UUID{}

			m.userService.LogOut(r.Context(), u)
			fmt.Println("Log out")
			return
		} else {
			next.ServeHTTP(w, r)
		}
	})
}

// func (m MeddlewireController) ContextMiddleware(ctx context.Context, next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		timeex := time.Now().Add(8 * time.Second).UTC()

// 		next.ServeHTTP(w, r.WithContext())
// 	})
// }
