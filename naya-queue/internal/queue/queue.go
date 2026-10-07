package queue

import "naya-queue/m/internal/job"

type Queue interface {
	Push(j job.Job) error
	Pop() (j job.Job, err error)
}
