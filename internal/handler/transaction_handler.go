package handler

import (
	"Tamrin/Expense_Tracker_API/internal/domain"
	"Tamrin/Expense_Tracker_API/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type TransactionHandler struct {
	service service.TransactionService
}

func NewTransactionHandler(service service.TransactionService) *TransactionHandler {
	return &TransactionHandler{service: service}
}
func (t *TransactionHandler) GetAll(c *gin.Context) {
	pageNumberStr := c.Query("page")
	pageNumber, err := strconv.Atoi(pageNumberStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if pageNumber <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "pageNumber must be positive integer"})
		return
	}
	value, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "user id not found",
		})
		return
	}
	valueInt, ok := value.(int64)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "incorrect user id"})
		return
	}
	transactions, err := t.service.GetAll(c.Request.Context(), valueInt, pageNumber)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": transactions})
}
func (t *TransactionHandler) GetByID(c *gin.Context) {
	transactionId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil || transactionId <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid transaction id"})
		return
	}
	value, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "user id not found",
		})
		return
	}
	valueInt, ok := value.(int64)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "incorrect user id"})
		return
	}
	transaction, err := t.service.GetByID(c.Request.Context(), valueInt, int64(transactionId))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": transaction})
}
func (t *TransactionHandler) Create(c *gin.Context) {
	userIdAny, ok := c.Get("user_id")
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user id not found"})
		return
	}
	var transaction domain.CreateTransactionRequest
	err := c.ShouldBindJSON(&transaction)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userIdInt, ok := userIdAny.(int64)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "incorrect user id"})
		return
	}
	transactionId, err := t.service.Create(c.Request.Context(), userIdInt, transaction)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"transaction_id": transactionId})
}
func (t *TransactionHandler) Update(c *gin.Context) {
	userIdAny, ok := c.Get("user_id")
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user id not found"})
		return
	}
	userIdInt, ok := userIdAny.(int64)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "incorrect user id"})
		return
	}
	var transaction domain.CreateTransactionRequest
	err := c.ShouldBindJSON(&transaction)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	transactionId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil || transactionId <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid transaction id"})
		return
	}
	err = t.service.Update(c.Request.Context(), userIdInt, transactionId, transaction)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": transaction})
}
func (t *TransactionHandler) Delete(c *gin.Context) {
	userIdAny, ok := c.Get("user_id")
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user id not found"})
		return
	}
	userIdInt, ok := userIdAny.(int64)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "incorrect user id"})
		return
	}
	transactionId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil || transactionId <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid transaction id"})
		return
	}
	err = t.service.Delete(c.Request.Context(), userIdInt, transactionId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}
