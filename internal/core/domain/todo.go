package domain

import "time"

type Todo struct {
	ID          string
	Title       string
	Description string
	IsCompleted bool
	CreatedAt   time.Time
}
