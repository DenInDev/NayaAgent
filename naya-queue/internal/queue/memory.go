package queue

import (
	"errors"
	"naya-queue/m/internal/job"
)

type MemoryQueue struct {
	jobs []job.Job
}

func (q *MemoryQueue) Push(j job.Job) error {
	q.jobs = append(q.jobs, j)
	return nil
}

func (q *MemoryQueue) Pop() (job.Job, error) {
	if len(q.jobs) == 0 {
		return job.Job{}, errors.New("queue is empty")
	}

	j := q.jobs[0]

	q.jobs = q.jobs[1:len(q.jobs)]

	return j, nil
}
