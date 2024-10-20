package models

import (
	"database/sql"
	"log"
	"time"
)

type Post struct {
	ID        int
	Title     string
	Content   string
	CreatedAt time.Time
}

func GetAllPosts(db *sql.DB) ([]Post, error) {
	rows, err := db.Query("SELECT id, title, content, created_at FROM posts ORDER BY created_at DESC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []Post
	for rows.Next() {
		var p Post
		err := rows.Scan(&p.ID, &p.Title, &p.Content, &p.CreatedAt)
		if err != nil {
			return nil, err
		}
		posts = append(posts, p)
	}

	return posts, nil
}

func GetPostByID(db *sql.DB, id int) (Post, error) {
	var p Post
	err := db.QueryRow("SELECT id, title, content, created_at FROM posts WHERE id = $1", id).Scan(&p.ID, &p.Title, &p.Content, &p.CreatedAt)
	if err != nil {
		return Post{}, err
	}
	return p, nil
}

func CreatePost(db *sql.DB, title, content string) (int, error) {
	var id int
	err := db.QueryRow("INSERT INTO posts (title, content) VALUES ($1, $2) RETURNING id", title, content).Scan(&id)
	if err != nil {
		log.Printf("Error inserting post into database: %v", err) // Add this line
		return 0, err
	}
	log.Printf("Post created with ID: %d", id) // Add this line
	return id, nil
}
