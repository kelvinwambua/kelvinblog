package handlers

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/kelvinwambua/kelvinblog/internal/models"
)

func IndexHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("IndexHandler called")
		posts, err := models.GetAllPosts(db)
		if err != nil {
			log.Printf("Error getting posts: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		log.Println("Parsing templates...")

		// Define the function map
		funcMap := template.FuncMap{
			"truncate": func(s string, max int) string {
				if len(s) <= max {
					return s
				}
				return s[:max] + "..."
			},
		}

		// Create a new template with the function map
		tmpl := template.New("layout").Funcs(funcMap)

		// Parse the template files
		tmpl, err = tmpl.ParseFiles("templates/layout.html", "templates/index.html")
		if err != nil {
			log.Printf("Error parsing templates: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		log.Println("Executing template...")
		err = tmpl.ExecuteTemplate(w, "layout", posts)
		if err != nil {
			log.Printf("Error executing template: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		log.Println("IndexHandler completed")
	}
}

func PostHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			http.Error(w, "Invalid post ID", http.StatusBadRequest)
			return
		}

		post, err := models.GetPostByID(db, id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		tmpl, err := template.ParseFiles("templates/layout.html", "templates/post.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		err = tmpl.ExecuteTemplate(w, "layout", post)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

func CreatePostPageHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles("templates/create.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		err = tmpl.Execute(w, nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}
func CreatePostHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			http.Error(w, "Failed to parse form", http.StatusBadRequest)
			return
		}

		title := r.FormValue("title")
		content := r.FormValue("content")

		if title == "" || content == "" {
			http.Error(w, "Title and content are required", http.StatusBadRequest)
			return
		}

		_, err := models.CreatePost(db, title, content)
		if err != nil {
			http.Error(w, "Failed to create post: "+err.Error(), http.StatusInternalServerError)
			return
		}

		// Redirect to the home page after successful post creation
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}
