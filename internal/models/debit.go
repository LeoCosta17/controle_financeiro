package models

import "time"

type Debit struct {
	ID           uint64
	SupplierID   uint64
	SupplierName string
	Value        float64
	CreatedAt    time.Time
	DueDate      time.Time
}
