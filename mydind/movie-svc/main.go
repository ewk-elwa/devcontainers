package main

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	"log"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/moviedb")
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
}

func getMoviesFromDB() ([]Movie, error) {
	rows, err := db.Query("SELECT id, title, status FROM movies")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var movies []Movie
	for rows.Next() {
		var movie Movie
		if err := rows.Scan(&movie.ID, &movie.Title, &movie.Status); err != nil {
			return nil, err
		}
		movies = append(movies, movie)
	}
	return movies, nil
}

func getMovieFromDB(id string) (Movie, error) {
	var movie Movie
	err := db.QueryRow("SELECT id, title, status FROM movies WHERE id = ?", id).Scan(&movie.ID, &movie.Title, &movie.Status)
	if err != nil {
		return movie, err
	}
	return movie, nil
}

func updateMovieStatusInDB(id string, status string) error {
	_, err := db.Exec("UPDATE movies SET status = ? WHERE id = ?", status, id)
	return err
}

type Movie struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Status string `json:"status"`
}

var movies = []Movie{
	{ID: "1", Title: "Inception", Status: "stopped"},
	{ID: "2", Title: "Interstellar", Status: "stopped"},
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, item := range movies {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	http.NotFound(w, r)
}

func listMovies(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(movies)
}

func playMovie(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for i, item := range movies {
		if item.ID == params["id"] {
			movies[i].Status = "playing"
			json.NewEncoder(w).Encode(movies[i])
			return
		}
	}
	http.NotFound(w, r)
}

func stopMovie(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for i, item := range movies {
		if item.ID == params["id"] {
			movies[i].Status = "stopped"
			json.NewEncoder(w).Encode(movies[i])
			return
		}
	}
	http.NotFound(w, r)
}

func pauseMovie(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for i, item := range movies {
		if item.ID == params["id"] {
			movies[i].Status = "paused"
			json.NewEncoder(w).Encode(movies[i])
			return
		}
	}
	http.NotFound(w, r)
}

func forwardMovie(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for i, item := range movies {
		if item.ID == params["id"] {
			movies[i].Status = "forwarding"
			json.NewEncoder(w).Encode(movies[i])
			return
		}
	}
	http.NotFound(w, r)
}

func rewindMovie(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for i, item := range movies {
		if item.ID == params["id"] {
			movies[i].Status = "rewinding"
			json.NewEncoder(w).Encode(movies[i])
			return
		}
	}
	http.NotFound(w, r)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/movies", listMovies).Methods("GET")
	router.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	router.HandleFunc("/movies/{id}/play", playMovie).Methods("POST")
	router.HandleFunc("/movies/{id}/stop", stopMovie).Methods("POST")
	router.HandleFunc("/movies/{id}/pause", pauseMovie).Methods("POST")
	router.HandleFunc("/movies/{id}/forward", forwardMovie).Methods("POST")
	router.HandleFunc("/movies/{id}/rewind", rewindMovie).Methods("POST")
	http.ListenAndServe(":8000", router)
}
