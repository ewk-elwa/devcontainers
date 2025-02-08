package routes

import (
    "github.com/gin-gonic/gin"
    "recipebook/src/controllers"
)

func SetupRoutes(router *gin.Engine) {
    recipeController := controllers.RecipeController{}

    router.GET("/recipes", recipeController.GetAllRecipes)
    router.GET("/recipes/:id", recipeController.GetRecipe)
    router.POST("/recipes", recipeController.CreateRecipe)
}