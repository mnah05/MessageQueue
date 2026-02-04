package main

import (
	"log"

	"jobqueue/internal/redis"
	"jobqueue/internal/tasks"
	"jobqueue/internal/workers"

	// "jobqueue/internal/tasks"
	// "jobqueue/internal/workers"

	"github.com/hibiken/asynq"
)

func main() {
	server := redis.NewServer()

	mux := asynq.NewServeMux()
	mux.HandleFunc(tasks.TypeSendEmail, workers.HandleNewWelcomeEmail)
	// Pub: Event router
	mux.HandleFunc(tasks.TypeOrderCreated, workers.HandleOrderCreated)

	// Sub: Independent handlers
	mux.HandleFunc(tasks.TypeSendEmail, workers.HandleSendEmail)
	mux.HandleFunc(tasks.TypeUpdateInventory, workers.HandleUpdateInventory)

	log.Println("Worker started")
	if err := server.Run(mux); err != nil {
		log.Fatal(err)
	}
}
