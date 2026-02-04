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
	log.Println("📋 Registering task handlers...")
	mux.HandleFunc(tasks.TypeNewUserEmail, workers.HandleNewWelcomeEmail)
	mux.HandleFunc(tasks.TypeOrderCreated, workers.HandleOrderCreated)
	log.Printf("   ✅ Registered: %s -> HandleOrderCreated", tasks.TypeOrderCreated)

	mux.HandleFunc(tasks.TypeSendEmail, workers.HandleSendEmail)
	log.Printf("   ✅ Registered: %s -> HandleSendEmail", tasks.TypeSendEmail)

	mux.HandleFunc(tasks.TypeUpdateInventory, workers.HandleUpdateInventory)
	log.Printf("   ✅ Registered: %s -> HandleUpdateInventory", tasks.TypeUpdateInventory)

	log.Println("Worker started")
	if err := server.Run(mux); err != nil {
		log.Fatal(err)
	}
}
