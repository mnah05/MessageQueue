package workers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"jobqueue/internal/tasks"

	"github.com/hibiken/asynq"
)

func HandleSendEmail(ctx context.Context, t *asynq.Task) error {
	var p tasks.SendEmailPayload
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return err
	}

	log.Printf("📧 Sending %s email to %s for order %s", p.Type, p.Email, p.OrderID)

	if err := sendEmail(p); err != nil {
		return fmt.Errorf("send email failed: %w", err)
	}

	return nil
}

func sendEmail(p tasks.SendEmailPayload) error {

	fmt.Printf("→ Email sent: %s confirmation to %s\n", p.Type, p.Email)
	return nil
}

func HandleNewWelcomeEmail(ctx context.Context, t *asynq.Task) error {
	var payload tasks.TaskSendWelcomeEmail
	if err := json.Unmarshal(t.Payload(), &payload); err != nil {
		return err
	}
	fmt.Printf("Welcome and thanks for joining us:%s\n", payload.Email)
	return nil
}
