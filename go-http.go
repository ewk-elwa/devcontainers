// The code implements a basic CRUD (Create, Read, Update, Delete) system for managing “items” in memory. Review this code and provide feedback on correctness, design, performance, and maintainability.

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
)

type Item struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Price float64 `json:"price"`
}

type ItemStore struct {
	mu    sync.Mutex
	items map[int]Item
	nextID int
}

func NewItemStore() *ItemStore {
	return &ItemStore{
		items: make(map[int]Item),
		nextID: 1,
	}
}

func (s *ItemStore) CreateItem(item Item) Item {
	s.mu.Lock()
	defer s.mu.Unlock()
	item.ID = s.nextID
	s.nextID++
	s.items[item.ID] = item
	return item
}

func (s *ItemStore) GetItem(id int) (Item, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	item, exists := s.items[id]
	return item, exists
}

func (s *ItemStore) UpdateItem(id int, updated Item) (Item, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	item, exists := s.items[id]
	if !exists {
		return Item{}, false
	}
	updated.ID = id
	s.items[id] = updated
	return updated, true
}

func (s *ItemStore) DeleteItem(id int) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	_, exists := s.items[id]
	if exists {
		delete(s.items, id)
		return true
	}
	return false
}

func main() {
	store := NewItemStore()
	http.HandleFunc("/items", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			var item Item
			if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
				http.Error(w, "Invalid request body", http.StatusBadRequest)
				return
			}
			createdItem := store.CreateItem(item)
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(createdItem)
		} else if r.Method == http.MethodGet {
			store.mu.Lock()
			defer store.mu.Unlock()
			items := make([]Item, 0, len(store.items))
			for _, item := range store.items {
				items = append(items, item)
			}
			json.NewEncoder(w).Encode(items)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/items/", func(w http.ResponseWriter, r *http.Request) {
		idStr := r.URL.Path[len("/items/"):]
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "Invalid ID", http.StatusBadRequest)
			return
		}

		switch r.Method {
		case http.MethodGet:
			item, exists := store.GetItem(id)
			if !exists {
				http.Error(w, "Item not found", http.StatusNotFound)
				return
			}
			json.NewEncoder(w).Encode(item)
		case http.MethodPut:
			var updated Item
			if err := json.NewDecoder(r.Body).Decode(&updated); err != nil {
				http.Error(w, "Invalid request body", http.StatusBadRequest)
				return
			}
			item, success := store.UpdateItem(id, updated)
			if !success {
				http.Error(w, "Item not found", http.StatusNotFound)
				return
			}
			json.NewEncoder(w).Encode(item)
		case http.MethodDelete:
			if !store.DeleteItem(id) {
				http.Error(w, "Item not found", http.StatusNotFound)
				return
			}
			w.WriteHeader(http.StatusNoContent)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	fmt.Println("Server is running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
