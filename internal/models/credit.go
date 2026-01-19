package models

import "time"

type Credit struct {
	ID           uint64
	CostumerID   uint64
	CostumerName string
	Value        float64
	CreatedAt    time.Time
	DueDate      time.Time
}
