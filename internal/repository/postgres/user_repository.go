package postgres

import (
	"Tamrin/Expense_Tracker_API/internal/domain"
	"context"
	"database/sql"
)

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *userRepository {
	return &userRepository{db: db}
}
func (d *userRepository) Create(ctx context.Context, username string, passwordHash string) (int64, error) {
	query := `INSERT INTO users(username,password_hash) VALUES ($1,$2) RETURNING id`
	var id int64
	err := d.db.QueryRowContext(ctx, query, username, passwordHash).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}
func (d *userRepository) FindByUsername(ctx context.Context, username string) (domain.User, error) {
	query := `SELECT id,username,password_hash FROM users WHERE username = $1`
	var user domain.User
	err := d.db.QueryRowContext(ctx, query, username).Scan(&user.ID, &user.Username, &user.PasswordHash)
	if err != nil {
		return domain.User{}, err
	}
	return user, nil
}
func (d *userRepository) Update(ctx context.Context, userID int64, passwordHash string) error {
	query := `UPDATE users SET password_hash = $1, updated_at = NOW() WHERE id = $2`
	_, err := d.db.ExecContext(ctx, query, passwordHash, userID)
	if err != nil {
		return err
	}
	return nil
}
func (d *userRepository) Delete(ctx context.Context, userID int64) error {
	query := `DELETE FROM users WHERE id = $1`
	_, err := d.db.ExecContext(ctx, query, userID)
	if err != nil {
		return err
	}
	return nil
}
