package main

import (
	// to log out the error while connecting the server
	//allows to create server in golang
	//to encode data to json while using postman
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux" //external library
)

type Movie struct {
	ID    string `json:"id"`
	Isbn  string `json:"isbn"`
	Title string `json:"title"`
}

var movies []Movie //slice of the type movie

func getMovies(w http.ResponseWriter, r *http.Request) {
	//passing the pointer of the request from postman to this function
	w.Header().Set("Content-Type", "application/json")
	//w is the response header..when we send a respobse frm this fn it will be w
	json.NewEncoder(w).Encode(movies)
	// encoding data to json..to send json of all the movies of slice..returnd the movies
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	//id passed from postman will move as a param to this funct and id will be present inside
	//mux.vars and inside the request r ...extract the `id` that we wish to delete
	params := mux.Vars(r)

	//ranges over the movies..we then need to loop through all our movies
	for index, item := range movies {

		if item.ID == params["id"] {

			//updates our movies array to remove the  movie
			//it appends all other movies except the id selected to be removed
			//movies[:index] is the currently selected ids..this wont exixts..inplace of that next index movies comes
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	// using blnk identifier as index is not used
	for _, item := range movies {

		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}

	}
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// definin a variable movie of the type Movie
	var movie Movie
	//decode the body sent [sending the enitr body contents like id,title..new movie]
	//after decoding gets the value in the movie var
	_ = json.NewDecoder(r.Body).Decode(&movie)
	//creating a random id -- Intn-> to xreate random value btw 1 and this number
	//then format it into string-- Itoa->converts int to sring
	//movie.ID = strconv.Itoa(rand.Intn(100000))
	//append the new movie to the set of movies
	//the new movie come from the body is inside the movie that we decode from json
	movies = append(movies, movie)
	//returns the created movie
	json.NewEncoder(w).Encode(movie)
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	//set json content type
	w.Header().Set("Content-Type", "application/json")
	//setting the params in mux to pass the id
	params := mux.Vars(r)

	// looping over the movies
	for index, item := range movies {
		if item.ID == params["id"] {
			// delete the movie with the id you hve sent
			movies = append(movies[:index], movies[index+1:]...)
			// creating a new movie..decoding the body that contains the new movie to the var movie
			// add new movie..the movie that is sent to postman
			var movie Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			// as need to use the same id
			movie.ID = params["id"]
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movie)
		}
	}
}

func main() {
	r := mux.NewRouter()

	movies = append(movies, Movie{ID: "1", Isbn: "456783", Title: "Movie one"})
	movies = append(movies, Movie{ID: "2", Isbn: "456678", Title: "Movie two"})

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Printf("STARTING SERVER AT PORT 8080\n")
	log.Fatal(http.ListenAndServe(":8080", r)) //to start the server in go and log for in case server not getting started
}
