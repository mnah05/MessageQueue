package main

import (
	"log"

	"jobqueue/internal/redis"
	"jobqueue/internal/tasks"
	"jobqueue/internal/workers"

	"github.com/hibiken/asynq"
)

func main() {
	server := redis.NewServer()

	mux := asynq.NewServeMux()
	mux.HandleFunc(tasks.TaskSendWelcomeEmail, workers.HandleWelcomeEmail)

	log.Println("Worker started")
	if err := server.Run(mux); err != nil {
		log.Fatal(err)
	}
}
