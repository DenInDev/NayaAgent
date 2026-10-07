package job

import "time"

type Job struct {
	ID            int
	CurrentStatus Status
	Type          string
	Attempts      int
	CreatedAt     time.Time
}
