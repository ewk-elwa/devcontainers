package handlers

import (
	"encoding/json"
	"net/http"
	"sync"

	"todo/src/models"
)

var (
	todos  = make(map[int]models.Todo)
	nextID = 1
	mu     sync.Mutex
)

// CreateTodo handles the creation of a new todo item
func CreateTodo(w http.ResponseWriter, r *http.Request) {
	var todo models.Todo
	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	mu.Lock()
	todo.ID = nextID
	nextID++
	todos[todo.ID] = todo
	mu.Unlock()

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(todo)
}

// GetTodos handles fetching all todo items
func GetTodos(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	todoList := make([]models.Todo, 0, len(todos))
	for _, todo := range todos {
		todoList = append(todoList, todo)
	}

	json.NewEncoder(w).Encode(todoList)
}

// UpdateTodo handles updating an existing todo item
func UpdateTodo(w http.ResponseWriter, r *http.Request) {
	var todo models.Todo
	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	mu.Lock()
	defer mu.Unlock()

	if _, exists := todos[todo.ID]; !exists {
		http.Error(w, "Todo not found", http.StatusNotFound)
		return
	}

	todos[todo.ID] = todo
	json.NewEncoder(w).Encode(todo)
}

// DeleteTodo handles deleting a todo item
func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	mu.Lock()
	defer mu.Unlock()

	for k := range todos {
		if k == id {
			delete(todos, k)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}

	http.Error(w, "Todo not found", http.StatusNotFound)
}