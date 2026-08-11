package repository

import (
	"Tamrin/Expense_Tracker_API/internal/domain"
	"context"
)

type TransactionRepository interface {
	GetAll(ctx context.Context, userID int64, limit int, offset int) ([]domain.Transaction, error)
	GetByID(ctx context.Context, id int64, userID int64) (domain.Transaction, error)
	Create(ctx context.Context, transaction domain.CreateTransactionRequest, userID int64) (int64, error)
	Update(ctx context.Context, transaction domain.CreateTransactionRequest, id int64, userID int64) error
	Delete(ctx context.Context, id int64, userID int64) error
}
type UserRepository interface {
	Create(ctx context.Context, username string, passwordHash string) (int64, error)
	FindByUsername(ctx context.Context, username string) (domain.User, error)
	Update(ctx context.Context, userID int64, passwordHash string) error
	Delete(ctx context.Context, userID int64) error
}
