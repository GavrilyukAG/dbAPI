package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/GavrilyukAG/dbAPI/handlers"

	"github.com/gorilla/mux"
)

func handleRequests(handlers *handlers.Handler) {
	router := mux.NewRouter()

	router.HandleFunc("/api/forum/create", handlers.CreateForum).Methods("POST")
	router.HandleFunc("/api/forum/{slug}/create", handlers.CreateThread).Methods("POST")
	router.HandleFunc("/api/forum/{slug}/details", handlers.GetForumDetails).Methods("GET")
	router.HandleFunc("/api/forum/{slug}/threads", handlers.GetThreadsList).Methods("GET")
	router.HandleFunc("/api/forum/{slug}/users", handlers.GetUsersList).Methods("GET")

	router.HandleFunc("/api/post/{id}/details", handlers.GetPostDetails).Methods("GET")
	router.HandleFunc("/api/post/{id}/details", handlers.UpdatePost).Methods("POST")

	router.HandleFunc("/api/service/clear", handlers.EraseDB).Methods("POST")
	router.HandleFunc("/api/service/status", handlers.GetStatus).Methods("GET")

	router.HandleFunc("/api/thread/{slug_or_id}/create", handlers.CreatePost).Methods("POST")
	router.HandleFunc("/api/thread/{slug_or_id}/details", handlers.GetThreadDetails).Methods("GET")
	router.HandleFunc("/api/thread/{slug_or_id}/details", handlers.UpdateThread).Methods("POST")
	router.HandleFunc("/api/thread/{slug_or_id}/posts", handlers.GetThreadPosts).Methods("GET")
	router.HandleFunc("/api/thread/{slug_or_id}/vote", handlers.Vote).Methods("POST")

	router.HandleFunc("/api/user/{nickname}/create", handlers.CreateUser).Methods("POST")
	router.HandleFunc("/api/user/{nickname}/profile", handlers.GetUserProfile).Methods("GET")
	router.HandleFunc("/api/user/{nickname}/profile", handlers.UpdateUser).Methods("POST")

	fmt.Println("Server is listening port 5000")
	http.ListenAndServe(":5000", router)
}

func main() {
	// dataSource := "host=127.0.0.1 port=5432 user=tpuser password=password dbname=tpforumdb sslmode=disable"
	dataSource := "host=127.0.0.1 port=5432 user=docker password=docker dbname=docker sslmode=disable"
	db, err := sql.Open("postgres", dataSource)
	if err != nil {
		// log.Fatal(err)
	}

	handlers := &handlers.Handler{
		DB: db,
	}

	handleRequests(handlers)
}
