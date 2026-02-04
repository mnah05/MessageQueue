package redis

import (
	"context"
	"fmt"
	"runtime"

	"github.com/hibiken/asynq"
)

func NewClient() *asynq.Client {
	return asynq.NewClient(asynq.RedisClientOpt{
		Addr: "localhost:6379",
	})
}

func NewServer() *asynq.Server {
	numOfWorkers := 2*(runtime.NumCPU()) + 1

	return asynq.NewServer(
		asynq.RedisClientOpt{
			Addr: "localhost:6379",
		},
		asynq.Config{
			Concurrency: numOfWorkers,
			Queues: map[string]int{
				"critical": 6,
				"default":  3,
				"low":      1,
			},
			ErrorHandler: asynq.ErrorHandlerFunc(
				func(ctx context.Context, t *asynq.Task, err error) {
					fmt.Println("TASK FAILED AFTER MAX RETRIES")
					fmt.Println("Type:", t.Type())
					fmt.Println("Payload:", string(t.Payload()))
					fmt.Println("Error:", err)
				},
			),
		},
	)
}
