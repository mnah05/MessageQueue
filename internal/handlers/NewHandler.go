package handlers

import (
	"encoding/json"
	"jobqueue/internal/redis"
	"jobqueue/internal/tasks"
	"log"
	"net/http"
	"strings"

	"github.com/hibiken/asynq"
)

type Request struct {
	Email string `json:"email"`
}

func NewUserHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "content-type must be application/json", http.StatusBadRequest)
		return
	}

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	var req Request
	if err := decoder.Decode(&req); err != nil {
		http.Error(w, "invalid JSON body", http.StatusBadRequest)
		return
	}

	// Prevent extra JSON after first object
	if decoder.More() {
		http.Error(w, "invalid JSON body", http.StatusBadRequest)
		return
	}

	// Validate email presence
	req.Email = strings.TrimSpace(req.Email)
	if req.Email == "" {
		http.Error(w, "email is required", http.StatusBadRequest)
		return
	}

	// Optional: minimal email sanity check
	if !strings.Contains(req.Email, "@") {
		http.Error(w, "invalid email", http.StatusBadRequest)
		return
	}
	client := redis.NewClient()
	defer client.Close()

	log.Println("new user email:", req.Email)
	var p tasks.TaskSendWelcomeEmail
	p.Email = req.Email
	task, err := tasks.NewUserEmail(p)
	if err != nil {
		http.Error(w, "could not create task", http.StatusInternalServerError)
		return
	}
	//now this task is retried for a max of 5 times and is in the critical queue
	if _, err := client.Enqueue(task, asynq.Queue("critical"), asynq.MaxRetry(5)); err != nil {
		http.Error(w, "could not create task", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{
		"email": req.Email,
	})
}
