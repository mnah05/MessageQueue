// internal/handlers/order.go
package handlers

import (
	"encoding/json"
	"net/http"

	"jobqueue/internal/redis"
	"jobqueue/internal/tasks"
)

type CreateOrderRequest struct {
	OrderID    string   `json:"order_id"`
	CustomerID string   `json:"customer_id"`
	Email      string   `json:"email"`
	Amount     float32  `json:"amount"`
	Items      []string `json:"items"`
}

func CreateOrder(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req CreateOrderRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if req.OrderID == "" || req.CustomerID == "" || req.Email == "" || len(req.Items) == 0 {
		http.Error(w, "Missing required fields", http.StatusBadRequest)
		return
	}

	// Create and enqueue the event
	payload := tasks.OrderCreatedPayload{
		OrderID:    req.OrderID,
		CustomerID: req.CustomerID,
		Email:      req.Email,
		Amount:     req.Amount,
		Items:      req.Items,
	}

	task, err := tasks.NewOrderCreatedTask(payload)
	if err != nil {
		http.Error(w, "Failed to create task", http.StatusInternalServerError)
		return
	}

	client := redis.NewClient()
	defer client.Close()

	info, err := client.Enqueue(task)
	if err != nil {
		http.Error(w, "Failed to enqueue task", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message":  "Order event published",
		"order_id": req.OrderID,
		"task_id":  info.ID,
	})
}
