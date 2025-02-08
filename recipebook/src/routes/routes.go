package routes

import (
    "github.com/gorilla/mux"
    "github.com/ewk-elwa/devcontainers/recipebook/src/controllers"
)

func InitializeRoutes() *mux.Router {
    router := mux.NewRouter()
    recipeController := controllers.RecipeController{}

    router.HandleFunc("/recipes", recipeController.GetAllRecipes).Methods("GET")
    router.HandleFunc("/recipes/:id", recipeController.GetRecipe).Methods("GET")
    router.HandleFunc("/recipes", recipeController.CreateRecipe).Methods("POST")
    return router
}