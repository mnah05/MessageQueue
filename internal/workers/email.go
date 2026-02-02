package workers

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"jobqueue/internal/tasks"

	"github.com/hibiken/asynq"
)

func HandleWelcomeEmail(ctx context.Context, t *asynq.Task) error {
	var payload tasks.WelcomeEmailPayload
	if err := json.Unmarshal(t.Payload(), &payload); err != nil {
		return err
	}

	time.Sleep(2 * time.Second)
	fmt.Println("Thanks for signing up:", payload.Email)

	return nil
}
