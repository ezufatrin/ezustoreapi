package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"ezustore/src/domain"
	"ezustore/src/infrastructure"
	"ezustore/src/infrastructure/database"
	"ezustore/src/interfaces"
	"ezustore/src/middleware"
	"ezustore/src/usecase"
)

func main() {
	// Load .env (ignore error if missing)
	_ = godotenv.Load()

	// Init DB
	db := database.InitMySQL()

	// AutoMigrate models
	if err := db.AutoMigrate(
		&domain.User{},
		&domain.Product{},
		&domain.Order{},
		&domain.OrderItem{},
		&domain.Transaction{},
		&domain.TransactionItem{},
	); err != nil {
		log.Fatalf("auto migrate error: %v", err)
	}

	// Init repositories
	userRepo := infrastructure.NewUserRepo(db)
	productRepo := infrastructure.NewProductRepo(db)
	orderRepo := infrastructure.NewOrderRepository(db)
	trxRepo := infrastructure.NewTransactionRepo(db)

	// Init usecases
	userUC := usecase.NewUserUsecase(userRepo)
	productUC := usecase.NewProductUsecase(productRepo)
	orderUC := usecase.NewOrderUsecase(orderRepo, productRepo)
	trxUC := usecase.NewTransactionUsecase(trxRepo, productRepo)

	// Gin router
	r := gin.Default()

	// Public routes
	interfaces.NewUserHandler(r, userUC)       // /register, /login
	interfaces.NewProductHandler(r, productUC) // GET list/detail produk publik

	// Protected routes group
	protected := r.Group("/")
	protected.Use(middleware.JWTMiddleware())
	{
		interfaces.RegisterUserProtectedRoutes(protected, userUC)
		interfaces.RegisterProductProtectedRoutes(protected, productUC) // POST/PUT/DELETE produk
		interfaces.NewOrderHandler(protected, orderUC)                  // checkout & riwayat order
		interfaces.NewTransactionHandler(protected, trxUC)              // contoh modul trx lain
	}

	// Run server
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("server error: %v", err)
	}
}
