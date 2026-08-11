package service

import (
	"Tamrin/Expense_Tracker_API/internal/domain"
	"context"
)

type TransactionService interface {
	GetAll(ctx context.Context, userID int64, page int) ([]domain.Transaction, error)
	GetByID(ctx context.Context, userID int64, transactionID int64) (domain.Transaction, error)
	Create(ctx context.Context, userID int64, transaction domain.CreateTransactionRequest) (int64, error)
	Update(ctx context.Context, userID int64, transactionID int64, transaction domain.CreateTransactionRequest) error
	Delete(ctx context.Context, userID int64, transactionID int64) error
}
type UserService interface {
	SignUp(ctx context.Context, transaction domain.AuthRequest) (int64, error)
	Login(ctx context.Context, username string, password string) (string, error)
	Update(ctx context.Context, userID int64, transaction domain.AuthRequest) error
	Delete(ctx context.Context, userID int64) error
}
