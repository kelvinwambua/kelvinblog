package main

import (
	"log"
	"net/http"
	"os"

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
	r.HandleFunc("/create", handlers.CreatePostPageHandler()).Methods("GET")
	r.HandleFunc("/create", handlers.CreatePostHandler(db)).Methods("POST")
	port := os.Getenv("PORT")
	if port == "" {
		port = "5000" // Default port if not specified
	}

	log.Printf("Server starting on :%s", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
