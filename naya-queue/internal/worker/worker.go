package worker

import (
	"fmt"
	"naya-queue/m/internal/queue"
)

type Worker struct {
	queue queue.Queue
}

func NewWorker(q queue.Queue) *Worker {
	return &Worker{
		queue: q,
	}
}

func (w *Worker) ProcessOne() error {
	j, err := w.queue.Pop()

	if err != nil {
		return err
	}

	fmt.Println(j.ID, j.Type)

	return nil
}
