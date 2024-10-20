package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kelvinwambua/kelvinblog/internal/database"
	"github.com/kelvinwambua/kelvinblog/internal/handlers"
)

func main() {
	db, err := database.InitDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	r := mux.NewRouter()

	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	r.HandleFunc("/", handlers.IndexHandler(db)).Methods("GET")
	r.HandleFunc("/post/{id}", handlers.PostHandler(db)).Methods("GET")
	r.HandleFunc("/create", handlers.CreatePostHandler(db)).Methods("POST")

	log.Println("Server starting on :5000")
	log.Fatal(http.ListenAndServe(":5000", r))
}
