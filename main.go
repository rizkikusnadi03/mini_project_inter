package main

import (
	"log"
	"os"

	"backend_golang/config"
	"backend_golang/internal/adapter/handler"
	"backend_golang/internal/adapter/repository"
	"backend_golang/internal/adapter/router"
	"backend_golang/internal/core/usecase"
	
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file, falling back to environment variables")
	}

	// Connect to Database
	db := config.ConnectDB()

	// Ensure upload directories exist
	os.MkdirAll("./uploads/products", os.ModePerm)
	os.MkdirAll("./uploads/stores", os.ModePerm)

	// Setup Repositories
	userRepo := repository.NewUserRepository(db)
	storeRepo := repository.NewStoreRepository(db)
	categoryRepo := repository.NewCategoryRepository(db)
	productRepo := repository.NewProductRepository(db)
	addressRepo := repository.NewAddressRepository(db)
	transactionRepo := repository.NewTransactionRepository(db)

	// Setup Usecases
	authUsecase := usecase.NewAuthUsecase(userRepo)
	categoryUsecase := usecase.NewCategoryUsecase(categoryRepo)
	productUsecase := usecase.NewProductUsecase(productRepo, storeRepo)
	storeUsecase := usecase.NewStoreUsecase(storeRepo)
	addressUsecase := usecase.NewAddressUsecase(addressRepo)
	transactionUsecase := usecase.NewTransactionUsecase(transactionRepo, productRepo)

	// Setup Handlers
	authHandler := handler.NewAuthHandler(authUsecase)
	categoryHandler := handler.NewCategoryHandler(categoryUsecase)
	productHandler := handler.NewProductHandler(productUsecase)
	storeHandler := handler.NewStoreHandler(storeUsecase)
	addressHandler := handler.NewAddressHandler(addressUsecase)
	transactionHandler := handler.NewTransactionHandler(transactionUsecase)
	provCityHandler := handler.NewProvCityHandler()

	// Initialize Fiber app
	app := fiber.New()

	// Enable CORS
	app.Use(cors.New())

	// Static route for serving uploaded images
	app.Static("/uploads", "./uploads")

	// Setup Routes
	router.SetupRoutes(app, authHandler, categoryHandler, productHandler, storeHandler, addressHandler, transactionHandler, provCityHandler)

	// Start Server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Fatal(app.Listen(":" + port))
}
