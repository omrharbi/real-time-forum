package controllers

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"real-time-froum/models"
	"real-time-froum/services"
)

type categoryController struct {
	categoryService services.CategoryService
}

func NewcategoryController(service services.CategoryService) *categoryController {
	return &categoryController{
		categoryService: service,
	}
}

func (c *categoryController) HandelCategory(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	categoryStruct := models.Category{}
	decode := DecodeJson(r)
	err := decode.Decode(&categoryStruct)
	if err != nil {
		JsoneResponse(w, r, err.Error(), http.StatusBadRequest)
		return
	}
	ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
	defer cancel()
	posts := c.categoryService.GetPostsByCategoryId(ctx, categoryStruct.Category)
	json.NewEncoder(w).Encode(posts)
}
