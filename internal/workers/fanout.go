// internal/workers/fanout.go
package workers

import (
	"context"
	"encoding/json"
	"log"
	"sync"

	"jobqueue/internal/redis"
	"jobqueue/internal/tasks"

	"github.com/hibiken/asynq"
)

// HandleOrderCreated receives the event and fans out to subscribers
func HandleOrderCreated(ctx context.Context, t *asynq.Task) error {
	var payload tasks.OrderCreatedPayload
	if err := json.Unmarshal(t.Payload(), &payload); err != nil {
		return err
	}

	client := redis.NewClient()
	defer client.Close()

	var wg sync.WaitGroup
	wg.Add(2)
	// Fan out to all subscribers concurrently
	go func() {
		defer wg.Done()
		enqueueEmail(client, payload)
	}()

	go func() {
		defer wg.Done()
		enqueueInventory(client, payload)
	}()

	wg.Wait()

	log.Printf("📢 OrderCreated event fanned out: %s", payload.OrderID)
	return nil
}

func enqueueEmail(client *asynq.Client, p tasks.OrderCreatedPayload) {
	task, _ := tasks.NewSendEmailTask(tasks.SendEmailPayload{
		OrderID:    p.OrderID,
		CustomerID: p.CustomerID,
		Email:      p.Email,
		Type:       "confirmation",
		Amount:     p.Amount,
	})
	if _, err := client.Enqueue(task); err != nil {
		log.Printf("❌ failed to enqueue email task: %v", err)
	}
}

func enqueueInventory(client *asynq.Client, p tasks.OrderCreatedPayload) {
	task, _ := tasks.NewUpdateInventoryTask(tasks.UpdateInventoryPayload{
		OrderID: p.OrderID,
		Items:   p.Items,
	})
	if _, err := client.Enqueue(task); err != nil {
		log.Printf("❌ failed to enqueue inventory task: %v", err)
	}
}
