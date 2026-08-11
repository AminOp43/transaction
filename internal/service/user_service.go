package service

import (
	"Tamrin/Expense_Tracker_API/internal/domain"
	"Tamrin/Expense_Tracker_API/internal/repository"
	"context"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"os"
	"time"
)

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *userService {
	return &userService{repo: repo}
}
func (r *userService) SignUp(ctx context.Context, transaction domain.AuthRequest) (int64, error) {
	if transaction.Username == "" {
		return 0, errors.New("username is required")
	}
	if transaction.Password == "" {
		return 0, errors.New("password is required")
	}
	passwordByte, err := bcrypt.GenerateFromPassword([]byte(transaction.Password), bcrypt.DefaultCost)
	if err != nil {
		return 0, err
	}
	id, err := r.repo.Create(ctx, transaction.Username, string(passwordByte))
	if err != nil {
		return 0, err
	}
	return id, nil
}
func (r *userService) Login(ctx context.Context, username string, password string) (string, error) {
	if username == "" {
		return "", errors.New("username is required")
	}
	if password == "" {
		return "", errors.New("password is required")
	}
	user, err := r.repo.FindByUsername(ctx, username)
	if err != nil {
		return "", err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return "", errors.New("invalid username or password")
	}
	expirationTime := time.Now().Add(24 * time.Hour)
	newJwt := jwt.MapClaims{"user_id": user.ID, "exp": expirationTime.Unix()}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, newJwt)
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
func (r *userService) Update(ctx context.Context, userID int64, transaction domain.AuthRequest) error {
	if transaction.Password == "" {
		return errors.New("password is required")
	}
	if userID < 1 {
		return errors.New("id must be greater than zero")
	}
	passwordByte, err := bcrypt.GenerateFromPassword([]byte(transaction.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	err = r.repo.Update(ctx, userID, string(passwordByte))
	if err != nil {
		return err
	}
	return nil
}
func (r *userService) Delete(ctx context.Context, userID int64) error {
	if userID < 1 {
		return errors.New("user_id must be greater than zero")
	}
	err := r.repo.Delete(ctx, userID)
	if err != nil {
		return err
	}
	return nil
}
