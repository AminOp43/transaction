package domain

import "time"

type Transaction struct {
	ID          int64
	UserID      int64
	Type        string
	Amount      int64
	Category    string
	Description string
	OccurredAt  time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
type CreateTransactionRequest struct {
	Type        string    `json:"type"`
	Amount      int64     `json:"amount"`
	Category    string    `json:"category"`
	Description string    `json:"description"`
	OccurredAt  time.Time `json:"occurred_at"`
}
