package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}
type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range movies {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	// json.NewEncoder(w).Encode(&Movie{})
}
func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(10000000))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movies)
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			var movie Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = params["id"]
			movies = append(movies, movie)
			return
		}
	}

}
func main() {
	router := mux.NewRouter()
	//data genarator 15 rows
	movies = append(movies, Movie{ID: "1", Isbn: "123", Title: "Cai bong su tu cua meo luoi", Director: &Director{Firstname: "Nguyen", Lastname: "Hoa"}})
	// movies = append(movies, Movie{ID: "2", Isbn: "456", Title: "Doremon", Director: &Director{Firstname: "Vjet", Lastname: "Nodejs"}})
	// movies = append(movies, Movie{ID: "3", Isbn: "789", Title: "Tom and jerry live acction", Director: &Director{Firstname: "Nguyen", Lastname: "Hoa"}})
	// movies = append(movies, Movie{ID: "4", Isbn: "101", Title: "Em la co gai anh ghet nhat", Director: &Director{Firstname: "Nguyen Hien", Lastname: "Vy"}})
	// movies = append(movies, Movie{ID: "5", Isbn: "102", Title: "End Game", Director: &Director{Firstname: "Nguyen Tuong", Lastname: "Vy"}})
	// movies = append(movies, Movie{ID: "6", Isbn: "103", Title: "Muon noi voi em anh la nguoi xau", Director: &Director{Firstname: "Nguyen Viet", Lastname: "Hoang"}})
	// movies = append(movies, Movie{ID: "7", Isbn: "104", Title: "Nguoi yeu em", Director: &Director{Firstname: "Nguyen", Lastname: "Hoa"}})
	// movies = append(movies, Movie{ID: "8", Isbn: "105", Title: "Nguoi ghet em", Director: &Director{Firstname: "John", Lastname: "Hunter"}})
	// movies = append(movies, Movie{ID: "9", Isbn: "106", Title: "Nguoi bo em", Director: &Director{Firstname: "kevin", Lastname: "Hoa"}})
	// movies = append(movies, Movie{ID: "10", Isbn: "107", Title: "Nguoi danh en", Director: &Director{Firstname: "Nguyen", Lastname: "Hoa"}})

	//fuction handler
	router.HandleFunc("/movies", getMovies).Methods("GET")
	router.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	router.HandleFunc("/movies", createMovie).Methods("POST")
	router.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	router.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")
	fmt.Printf("starting server at port 8080...\n")
	log.Fatal(http.ListenAndServe(":8080", router))
}
