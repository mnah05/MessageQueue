package tasks

import (
	"encoding/json"

	"github.com/hibiken/asynq"
)

const (
	TypeOrderCreated    = "order:created"
	TypeSendEmail       = "email:send"
	TypeUpdateInventory = "inventory:update"
	TypeNewUserEmail    = "email:new_user"
)

// OrderCreatedPayload represents data for order creation task
type OrderCreatedPayload struct {
	OrderID    string   `json:"order_id"`
	CustomerID string   `json:"customer_id"`
	Email      string   `json:"email"`
	Amount     float32  `json:"amount"`
	Items      []string `json:"items"`
}

// SendEmailPayload represents data for email sending task
type SendEmailPayload struct {
	OrderID    string  `json:"order_id"`
	CustomerID string  `json:"customer_id"`
	Email      string  `json:"email"`
	Type       string  `json:"type"`
	Amount     float32 `json:"amount"`
}

// UpdateInventoryPayload represents data for inventory update task
type UpdateInventoryPayload struct {
	OrderID string   `json:"order_id"`
	Items   []string `json:"items"`
}

type TaskSendWelcomeEmail struct {
	Email string `json:"emai"`
}

func NewUserEmail(payload TaskSendWelcomeEmail) (*asynq.Task, error) {
	data, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}
	return asynq.NewTask(TypeNewUserEmail, data), nil
}

// NewOrderCreatedTask creates a task for order creation event
func NewOrderCreatedTask(payload OrderCreatedPayload) (*asynq.Task, error) {
	data, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}
	return asynq.NewTask(TypeOrderCreated, data), nil
}

// NewSendEmailTask creates a task for sending email
func NewSendEmailTask(payload SendEmailPayload) (*asynq.Task, error) {
	data, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}
	return asynq.NewTask(TypeSendEmail, data), nil
}

// NewUpdateInventoryTask creates a task for inventory update
func NewUpdateInventoryTask(payload UpdateInventoryPayload) (*asynq.Task, error) {
	data, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}
	return asynq.NewTask(TypeUpdateInventory, data), nil
}
