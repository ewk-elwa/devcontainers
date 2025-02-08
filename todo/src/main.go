package main

import (
    "log"
    "net/http"
    "github.com/ewk-elwa/devcontainers/todo/src/routes"
)

func main() {
    router := routes.InitializeRoutes()
    log.Println("Starting server on :8080")
    if err := http.ListenAndServe(":8080", router); err != nil {
        log.Fatal(err)
    }
}