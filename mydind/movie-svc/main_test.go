package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"encoding/json"
	"fmt"
	"math/rand"
	"time"

	"github.com/gorilla/mux"
)

func init() {
	populateMovies()
}

func populateMovies() {
	rand.Seed(time.Now().UnixNano())
	titles := []string{"Inception", "Interstellar", "The Dark Knight", "Memento", "Dunkirk", "Tenet", "The Prestige", "Batman Begins", "Insomnia", "Following"}
	statuses := []string{"stopped", "playing", "paused", "forwarding", "rewinding"}

	for i := 1; i <= 100; i++ {
		movie := Movie{
			ID:     fmt.Sprintf("%d", i),
			Title:  titles[rand.Intn(len(titles))],
			Status: statuses[rand.Intn(len(statuses))],
		}
		movies = append(movies, movie)
	}
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	for _, movie := range movies {
		if movie.ID == id {
			json.NewEncoder(w).Encode(movie)
			return
		}
	}
	http.NotFound(w, r)
}

func listMovies(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(movies)
}

func playMovie(w http.ResponseWriter, r *http.Request) {
	changeMovieStatus(w, r, "playing")
}

func stopMovie(w http.ResponseWriter, r *http.Request) {
	changeMovieStatus(w, r, "stopped")
}

func pauseMovie(w http.ResponseWriter, r *http.Request) {
	changeMovieStatus(w, r, "paused")
}

func forwardMovie(w http.ResponseWriter, r *Request) {
	changeMovieStatus(w, r, "forwarding")
}

func rewindMovie(w http.ResponseWriter, r *Request) {
	changeMovieStatus(w, r, "rewinding")
}

func changeMovieStatus(w http.ResponseWriter, r *http.Request, status string) {
	vars := mux.Vars(r)
	id := vars["id"]

	for i, movie := range movies {
		if movie.ID == id {
			movies[i].Status = status
			json.NewEncoder(w).Encode(movies[i])
			return
		}
	}
	http.NotFound(w, r)
}

func TestGetMovie(t *testing.T) {
	req, err := http.NewRequest("GET", "/movies/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := `{"id":"1","title":"Inception","status":"stopped"}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

func TestListMovies(t *testing.T) {
	req, err := http.NewRequest("GET", "/movies", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/movies", listMovies).Methods("GET")
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := `[{"id":"1","title":"Inception","status":"stopped"},{"id":"2","title":"Interstellar","status":"stopped"}]`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

func TestPlayMovie(t *testing.T) {
	req, err := http.NewRequest("POST", "/movies/1/play", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/movies/{id}/play", playMovie).Methods("POST")
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := `{"id":"1","title":"Inception","status":"playing"}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

func TestStopMovie(t *testing.T) {
	req, err := http.NewRequest("POST", "/movies/1/stop", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/movies/{id}/stop", stopMovie).Methods("POST")
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := `{"id":"1","title":"Inception","status":"stopped"}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

func TestPauseMovie(t *testing.T) {
	req, err := http.NewRequest("POST", "/movies/1/pause", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/movies/{id}/pause", pauseMovie).Methods("POST")
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := `{"id":"1","title":"Inception","status":"paused"}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

func TestForwardMovie(t *testing.T) {
	req, err := http.NewRequest("POST", "/movies/1/forward", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/movies/{id}/forward", forwardMovie).Methods("POST")
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := `{"id":"1","title":"Inception","status":"forwarding"}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

func TestRewindMovie(t *testing.T) {
	req, err := http.NewRequest("POST", "/movies/1/rewind", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/movies/{id}/rewind", rewindMovie).Methods("POST")
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := `{"id":"1","title":"Inception","status":"rewinding"}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}
