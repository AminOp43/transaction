package service

import (
	"Tamrin/Expense_Tracker_API/internal/domain"
	"Tamrin/Expense_Tracker_API/internal/repository"
	"context"
	"errors"
)

type TransactionServ struct {
	repo repository.TransactionRepository
}

func NewTransactionServ(repo repository.TransactionRepository) *TransactionServ {
	return &TransactionServ{repo: repo}
}
func (s *TransactionServ) GetAll(ctx context.Context, userID int64, page int) ([]domain.Transaction, error) {
	if userID < 1 {
		return nil, errors.New("invalid user id")
	}
	if page <= 0 {
		return nil, errors.New("page must be greater than zero")
	}
	limit := 10
	offset := (page - 1) * limit
	transactions, err := s.repo.GetAll(ctx, userID, limit, offset)
	if err != nil {
		return nil, err
	}
	return transactions, nil
}
func (s *TransactionServ) GetByID(ctx context.Context, userID int64, transactionID int64) (domain.Transaction, error) {
	if userID < 1 || transactionID < 1 {
		return domain.Transaction{}, errors.New("invalid id or transaction id")
	}
	transaction, err := s.repo.GetByID(ctx, transactionID, userID)
	if err != nil {
		return domain.Transaction{}, err
	}
	return transaction, nil
}
func (s *TransactionServ) Create(ctx context.Context, userID int64, transaction domain.CreateTransactionRequest) (int64, error) {
	if userID < 1 {
		return 0, errors.New("invalid id")
	}
	if transaction.Type != "income" && transaction.Type != "expense" {
		return 0, errors.New("invalid transaction type")
	}

	if transaction.Amount <= 0 {
		return 0, errors.New("amount must be greater than zero")
	}

	if transaction.Category == "" {
		return 0, errors.New("category is required")
	}

	if transaction.OccurredAt.IsZero() {
		return 0, errors.New("occurred_at is required")
	}
	id, err := s.repo.Create(ctx, transaction, userID)
	if err != nil {
		return 0, err
	}
	return id, nil
}
func (s *TransactionServ) Update(ctx context.Context, userID int64, transactionID int64, transaction domain.CreateTransactionRequest) error {
	if userID < 1 || transactionID < 1 {
		return errors.New("invalid id or transaction id")
	}
	if transaction.Type != "income" && transaction.Type != "expense" {
		return errors.New("invalid transaction type")
	}

	if transaction.Amount <= 0 {
		return errors.New("amount must be greater than zero")
	}

	if transaction.Category == "" {
		return errors.New("category is required")
	}

	if transaction.OccurredAt.IsZero() {
		return errors.New("occurred_at is required")
	}
	err := s.repo.Update(ctx, transaction, transactionID, userID)
	if err != nil {
		return err
	}
	return nil
}
func (s *TransactionServ) Delete(ctx context.Context, userID int64, transactionID int64) error {
	if userID < 1 || transactionID < 1 {
		return errors.New("invalid id or transaction id")
	}
	err := s.repo.Delete(ctx, transactionID, userID)
	if err != nil {
		return err
	}
	return nil
}
