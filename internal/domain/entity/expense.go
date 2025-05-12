package entity

import (
	"SplitWise/internal/observer"
	"time"
)

type ExpenseState string

const (
	Created ExpenseState = "CREATED"
	Pending ExpenseState = "PENDING"
	Settled ExpenseState = "SETTLED"
)

type ExpenseHistory struct {
	Timestamp time.Time
	Message   string
}

type Expense struct {
	ID           string
	Title        string
	CreatedBy    string
	State        ExpenseState
	TotalAmount  float64
	Contributors map[string]*Contribution
	CreatedAt    time.Time
	SettledAt    *time.Time
	Observers    []observer.Observer // for event notification
	History      []ExpenseHistory
}
