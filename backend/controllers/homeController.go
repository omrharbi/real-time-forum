package controllers

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"real-time-froum/services"
)

type HomeController struct {
	cardService services.CardsService
}

func NewHomeController(card services.CardsService) *HomeController {
	return &HomeController{
		cardService: card,
	}
}

func (h *HomeController) HomeHandle(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	if r.Method != http.MethodGet {
		JsoneResponse(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	page := 1
	pageStr := r.URL.Query().Get("page")
	if pageStr != "" {
		if p, err := strconv.Atoi(pageStr); err == nil && p > 0 {
			page = p
		}
	}

	postsPerPage := 10
	ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
	defer cancel()
	posts, totalPosts := h.cardService.GetAllCardsForPages(ctx, page, postsPerPage)

	totalPages := (totalPosts + postsPerPage - 1) / postsPerPage

	response := PaginatedResponse{
		Posts:        posts,
		TotalPosts:   totalPosts,
		TotalPages:   totalPages,
		CurrentPage:  page,
		PostsPerPage: postsPerPage,
	}

	json.NewEncoder(w).Encode(response)
	// fmt.Println(response)
}

func (h *HomeController) GetCard_handler(res http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()

	if req.Method != http.MethodGet {
		JsoneResponse(res, "Status Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	id, err := strconv.Atoi(req.FormValue("id"))
	if err != nil {
		JsoneResponse(res, "Status Bad Request ID Uncorect", http.StatusBadRequest)
		return
	}
	card := h.cardService.GetOneCard(req.Context(), id)
	if card.Id == 0 {
		JsoneResponse(res, "Status Bad Request Not Have any card ", http.StatusBadRequest)
		return
	}
	json.NewEncoder(res).Encode(card)
}
