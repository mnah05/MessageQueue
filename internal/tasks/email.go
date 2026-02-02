package tasks

import (
	"encoding/json"

	"github.com/hibiken/asynq"
)

const TaskSendWelcomeEmail = "email:welcome"

type WelcomeEmailPayload struct {
	Email string `json:"email"`
}

func NewWelcomeEmailTask(email string) (*asynq.Task, error) {
	payload, err := json.Marshal(WelcomeEmailPayload{
		Email: email,
	})
	if err != nil {
		return nil, err
	}
	return asynq.NewTask(TaskSendWelcomeEmail, payload), nil
}
