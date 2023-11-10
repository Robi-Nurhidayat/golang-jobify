package jobs

import "time"

type Jobs struct {
	Id          int
	UserId      int
	Company     string
	Position    string
	Status      string
	JobType     string
	JobLocation string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
