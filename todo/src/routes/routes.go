package routes

import (
    // "net/http"
    "github.com/gorilla/mux"
    "github.com/ewk-elwa/devcontainers/todo/src/handlers"
)

func InitializeRoutes() *mux.Router {
    router := mux.NewRouter()

    router.HandleFunc("/todos", handlers.CreateTodo).Methods("POST")
    router.HandleFunc("/todos", handlers.GetTodos).Methods("GET")
    router.HandleFunc("/todos/{id}", handlers.UpdateTodo).Methods("PUT")
    router.HandleFunc("/todos/{id}", handlers.DeleteTodo).Methods("DELETE")

    return router
}