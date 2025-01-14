package controllers

import (
	"encoding/json"
	"net/http"

	"real-time-froum/services"
)

type profileController struct {
	profileService services.ProfileService
	userController *UserController
}

func NewprofileController(service services.ProfileService, userController *UserController) *profileController {
	return &profileController{
		profileService: service,
		userController: userController,
	}
}

func (p *profileController) HandleProfilePosts(w http.ResponseWriter, r *http.Request) {
	id_user := p.userController.GetUserId(r)
	posts := p.profileService.GetPostsProfile(r.Context(), id_user)
	json.NewEncoder(w).Encode(posts)
}

func (p *profileController) HandleProfileLikes(w http.ResponseWriter, r *http.Request) {
	id_user := p.userController.GetUserId(r)
	posts := p.profileService.GetProfileByLikes(r.Context(), id_user)
	json.NewEncoder(w).Encode(posts)
}
