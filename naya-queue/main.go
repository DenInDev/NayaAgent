package main

import (
	"fmt"
	"naya-queue/m/internal/job"
	"naya-queue/m/internal/queue"
	"naya-queue/m/internal/worker"
	"time"
)

func main() {
	fmt.Println("Forge queue starting...")

	newJob1 := job.Job{
		ID:            1,
		CurrentStatus: job.StatusRunning,
		Type:          "echo",
		Attempts:      2,
		CreatedAt:     time.Now(),
	}

	newJob2 := job.Job{
		ID:            2,
		CurrentStatus: job.StatusRunning,
		Type:          "echo",
		Attempts:      2,
		CreatedAt:     time.Now(),
	}

	newJob3 := job.Job{
		ID:            3,
		CurrentStatus: job.StatusRunning,
		Type:          "echo",
		Attempts:      2,
		CreatedAt:     time.Now(),
	}

	firstQueue := queue.MemoryQueue{}
	firstQueue.Push(newJob1)
	firstQueue.Push(newJob2)
	firstQueue.Push(newJob3)

	fmt.Println(firstQueue)

	for i := 0; i < 4; i++ {
		j, err := firstQueue.Pop()

		if err != nil {
			fmt.Println("error:", err)
			continue
		}

		fmt.Println(j.ID)
	}

	w := worker.NewWorker(firstQueue)
}
