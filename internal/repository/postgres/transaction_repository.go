package postgres

import (
	"Tamrin/Expense_Tracker_API/internal/domain"
	"context"
	"database/sql"
	"errors"
)

type TransactionRepo struct {
	DB *sql.DB
}

func NewTransactionRepo(db *sql.DB) *TransactionRepo {
	return &TransactionRepo{DB: db}
}
func (t *TransactionRepo) GetAll(ctx context.Context, userID int64, limit int, offset int) ([]domain.Transaction, error) {
	query := `
		SELECT id, user_id, type, amount, category, description, occurred_at, created_at, updated_at
		FROM transactions
		WHERE user_id = $1
		ORDER BY occurred_at DESC
		LIMIT $2 OFFSET $3
	`
	//بخاطر اینکه اگر transactions خالی هم باشه [] برمیگردونه و null برنمیگردونه
	//و limit هم مشخص میکنه که ظرفیت اسلایس transaction چقدر باشه
	transactions := make([]domain.Transaction, 0, limit)
	rows, err := t.DB.QueryContext(ctx, query, userID, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var transaction domain.Transaction
		err := rows.Scan(
			&transaction.ID,
			&transaction.UserID,
			&transaction.Type,
			&transaction.Amount,
			&transaction.Category,
			&transaction.Description,
			&transaction.OccurredAt,
			&transaction.CreatedAt,
			&transaction.UpdatedAt)
		if err != nil {
			return nil, err
		}
		transactions = append(transactions, transaction)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return transactions, nil
}

func (t *TransactionRepo) GetByID(ctx context.Context, id int64, userID int64) (domain.Transaction, error) {
	query := `
		SELECT id, user_id, type, amount, category, description, occurred_at, created_at, updated_at
		FROM transactions
		WHERE user_id = $1 AND id = $2
	`
	var transaction domain.Transaction
	err := t.DB.QueryRowContext(ctx, query, userID, id).Scan(
		&transaction.ID,
		&transaction.UserID,
		&transaction.Type,
		&transaction.Amount,
		&transaction.Category,
		&transaction.Description,
		&transaction.OccurredAt,
		&transaction.CreatedAt,
		&transaction.UpdatedAt)
	if errors.Is(err, sql.ErrNoRows) {
		return domain.Transaction{}, errors.New("transaction not found")
	}
	if err != nil {
		return domain.Transaction{}, err
	}
	return transaction, nil
}
func (t *TransactionRepo) Create(ctx context.Context, transaction domain.CreateTransactionRequest, userID int64) (int64, error) {
	query := `INSERT INTO transactions(user_id,type,amount, category, description, occurred_at) VALUES ($1,$2,$3,$4,$5,$6) RETURNING id`
	var id int64
	err := t.DB.QueryRowContext(ctx, query, userID, transaction.Type, transaction.Amount, transaction.Category, transaction.Description, transaction.OccurredAt).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}
func (t *TransactionRepo) Update(ctx context.Context, transaction domain.CreateTransactionRequest, id int64, userID int64) error {
	query := `UPDATE transactions SET type=$1, amount=$2, category=$3, description=$4, occurred_at=$5,updated_at = NOW() WHERE id = $6 AND user_id = $7`
	res, err := t.DB.ExecContext(ctx, query, transaction.Type, transaction.Amount, transaction.Category, transaction.Description, transaction.OccurredAt, id, userID)
	if err != nil {
		return err
	}
	affectedRows, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if affectedRows == 0 {
		return errors.New("transaction not found")
	}
	return nil
}
func (t *TransactionRepo) Delete(ctx context.Context, id int64, userID int64) error {
	query := `DELETE FROM transactions WHERE id = $1 AND user_id = $2`
	res, err := t.DB.ExecContext(ctx, query, id, userID)
	if err != nil {
		return err
	}
	affectedRows, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if affectedRows == 0 {
		return errors.New("transaction not found")
	}
	return nil
}
