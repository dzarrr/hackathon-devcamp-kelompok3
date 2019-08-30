package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	_ "github.com/lib/pq"
	"yukhaji/internal"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "\"\""
	dbname   = "yukhaji"
  )

func initFlags(args *internal.Args) {
	port := flag.Int("port", 3000, "port number for your apps")
	args.Port = *port
}

func initHandler(handler *internal.Handler) error {

	// Initialize SQL DB
	// db, err := sql.Open("postgres", "postgres://postgres:postgres@localhost:5432/?sslmode=disable")
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return err
	}
	handler.DB = db

	return nil
}

func initRouter(router *httprouter.Router, handler *internal.Handler) {

	router.GET("/", handler.Index)

	// Single user API
	router.GET("/user/:userID", handler.GetUserByID)
	router.POST("/user", handler.InsertUser)
	router.PUT("/user/:userID", handler.EditUserByID)
	router.DELETE("/user/:userID", handler.DeleteUserByID)

	router.GET("/saving/:userID", handler.GetSavingByUserID)
	router.POST("/saving", handler.InsertSaving)
	router.PUT("/saving/:userID", handler.EditSaving)
	router.PUT("/addbalance/:userID", handler.AddBalance)
	router.PUT("/editenddate/:userID", handler.EditEndDate)
	router.DELETE("/saving/:userID", handler.DeleteSavingByUserID)

	// Single book API
	router.GET("/book/:bookID", handler.GetBookByID)
	router.POST("/book", handler.InsertBook)
	router.PUT("/book/:bookID", handler.EditBook)
	router.DELETE("/book/:bookID", handler.DeleteBookByID)

	// Batch book API
	router.POST("/books", handler.InsertMultipleBooks)

	// Lending API
	router.POST("/lend", handler.LendBook)

	// `httprouter` library uses `ServeHTTP` method for it's 404 pages
	router.NotFound = handler
}

func main() {
	args := new(internal.Args)
	initFlags(args)

	handler := new(internal.Handler)
	if err := initHandler(handler); err != nil {
		panic(err)
	}

	router := httprouter.New()
	initRouter(router, handler)

	fmt.Printf("Apps served on :%d\n", args.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", args.Port), router))
}
