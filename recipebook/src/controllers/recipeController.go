package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/ewk-elwa/devcontainers/recipebook/src/models"
	"github.com/ewk-elwa/devcontainers/recipebook/src/services"
)

type RecipeController struct {
	service services.RecipeService
}

func NewRecipeController(service services.RecipeService) *RecipeController {
	return &RecipeController{service: service}
}

func (rc *RecipeController) GetRecipe(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid recipe ID", http.StatusBadRequest)
		return
	}

	recipe, err := rc.service.GetRecipeByID(id)
	if err != nil {
		http.Error(w, "Recipe not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(recipe)
}

func (rc *RecipeController) GetAllRecipes(w http.ResponseWriter, r *http.Request) {
	recipes := rc.service.GetAllRecipes()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(recipes)
}

func (rc *RecipeController) CreateRecipe(w http.ResponseWriter, r *http.Request) {
	var recipe models.Recipe
	if err := json.NewDecoder(r.Body).Decode(&recipe); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	createdRecipe, err := rc.service.CreateRecipe(recipe)
	if err != nil {
		http.Error(w, "Error creating recipe", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(createdRecipe)
}