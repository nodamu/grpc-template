package domain

import "time"

type Todo struct {
	Title       string
	Description string
	IsCompleted bool
	CreatedAt   time.Time
}
