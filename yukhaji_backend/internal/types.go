package internal

import (
	"database/sql"
	"time"
)

// Args used for this application
type Args struct {

	// Port used by this service
	Port int
}

// Handler object used to handle the HTTP API
type Handler struct {

	// DB object that'll be used
	DB *sql.DB
}

// User struct for database query
type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
}

// Book struct for database query
type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	ISBN   string `json:"isbn"`
	Stock  int64  `json:"stock"`
}

type Saving struct {
	ID     	int	`json:"id"`
	UserID 	int	`json:"user_id"`
	Balance int	`json:"balance"`
	Target 	int `json:"target"`
	StartDate 	time.Time `json:"start_date"`
	EndDate 	time.Time `json:"end_date"`
}

// LendRequest struct for receiving lend http request
type LendRequest struct {
	UserID int `json:"user_id"`
	BookID int `json:"book_id"`
}
