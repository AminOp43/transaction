package main

import (
	"Tamrin/Expense_Tracker_API/internal/handler"
	"Tamrin/Expense_Tracker_API/internal/middleware"
	"Tamrin/Expense_Tracker_API/internal/repository/postgres"
	"Tamrin/Expense_Tracker_API/internal/service"
	"Tamrin/Expense_Tracker_API/pkg/db"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"time"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("error loading .env file")
	}
	database := db.InitDB()
	defer database.Close()
	transactionRepo := postgres.NewTransactionRepo(database)
	userRepo := postgres.NewUserRepository(database)
	transactionService := service.NewTransactionServ(transactionRepo)
	userService := service.NewUserService(userRepo)
	transactionHandler := handler.NewTransactionHandler(transactionService)
	userHandler := handler.NewUserHandler(userService)
	router := gin.Default()
	router.Use(func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c.Request.Context(), time.Second*5)
		defer cancel()
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	})
	transaction := router.Group("/transaction")
	transaction.Use(middleware.AuthMiddleware)
	transaction.GET("/", transactionHandler.GetAll)
	transaction.GET("/:id", transactionHandler.GetByID)
	transaction.POST("/", transactionHandler.Create)
	transaction.PUT("/:id", transactionHandler.Update)
	transaction.DELETE("/:id", transactionHandler.Delete)

	protectedUser := router.Group("/user")
	protectedUser.Use(middleware.AuthMiddleware)
	protectedUser.PUT("/", userHandler.Update)
	protectedUser.DELETE("/", userHandler.Delete)

	//signup doesnt want jwt and login at the the start dont use jwt
	user := router.Group("/user")
	user.POST("/login", userHandler.Login)
	user.POST("/signup", userHandler.Signup)

	router.Run(":8080")
}
