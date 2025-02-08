package main

import (
    "github.com/gin-gonic/gin"
    "recipebook/src/routes"
)

func main() {
    r := gin.Default()
    routes.SetupRoutes(r)
    r.Run(":8080")
}