package main

import (
	handlers "jobqueue/internal/handlers"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/new-user", handlers.NewUserHandler)

	log.Println("server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
