package models

type Recipe struct {
    ID          int   `json:"id"`
    Title       string   `json:"title"`
    Ingredients []string `json:"ingredients"`
    Instructions string   `json:"instructions"`
}