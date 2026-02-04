// internal/workers/inventory.go
package workers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"jobqueue/internal/tasks"

	"github.com/hibiken/asynq"
)

func HandleUpdateInventory(ctx context.Context, t *asynq.Task) error {
	var p tasks.UpdateInventoryPayload
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return err
	}

	log.Printf("📦 Updating inventory for order %s, items: %v", p.OrderID, p.Items)

	if err := updateStock(p); err != nil {
		return fmt.Errorf("inventory update failed: %w", err)
	}

	return nil
}

func updateStock(p tasks.UpdateInventoryPayload) error {
	fmt.Printf("→ Stock reserved for items: %v\n", p.Items)
	return nil
}
