package services

import (
	"errors"
	"sync"
	"recipebook/src/models"
)

type RecipeService struct {
	recipes []models.Recipe
	mu      sync.Mutex
}

func NewRecipeService() *RecipeService {
	return &RecipeService{
		recipes: []models.Recipe{},
	}
}

func (rs *RecipeService) GetAllRecipes() []models.Recipe {
	rs.mu.Lock()
	defer rs.mu.Unlock()
	return rs.recipes
}

func (rs *RecipeService) CreateRecipe(recipe models.Recipe) error {
	rs.mu.Lock()
	defer rs.mu.Unlock()

	if recipe.ID == 0 {
		return errors.New("recipe ID must be greater than 0")
	}

	for _, r := range rs.recipes {
		if r.ID == recipe.ID {
			return errors.New("recipe with this ID already exists")
		}
	}

	rs.recipes = append(rs.recipes, recipe)
	return nil
}

func (rs *RecipeService) GetRecipeByID(id int) (*models.Recipe, error) {
	rs.mu.Lock()
	defer rs.mu.Unlock()

	for _, recipe := range rs.recipes {
		if recipe.ID == id {
			return &recipe, nil
		}
	}
	return nil, errors.New("recipe not found")
}