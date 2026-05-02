package router

import (
	"backend_golang/internal/adapter/handler"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, 
	authHandler *handler.AuthHandler,
	categoryHandler *handler.CategoryHandler,
	productHandler *handler.ProductHandler,
	storeHandler *handler.StoreHandler,
	addressHandler *handler.AddressHandler,
	transactionHandler *handler.TransactionHandler,
	provCityHandler *handler.ProvCityHandler) {

	api := app.Group("/api")

	// Auth
	auth := api.Group("/auth")
	auth.Post("/register", authHandler.Register)
	auth.Post("/login", authHandler.Login)

	// Province & City (Mock)
	api.Get("/provinces", provCityHandler.GetProvinces)
	api.Get("/provinces/:prov_id/cities", provCityHandler.GetCities)

	// Categories (Admin only)
	categories := api.Group("/category")
	categories.Get("/", categoryHandler.GetAll) // Public
	categories.Post("/", handler.Protected(), handler.AdminOnly(), categoryHandler.Create)
	categories.Put("/:id", handler.Protected(), handler.AdminOnly(), categoryHandler.Update)
	categories.Delete("/:id", handler.Protected(), handler.AdminOnly(), categoryHandler.Delete)

	// Products
	products := api.Group("/product")
	products.Get("/", productHandler.GetAll)
	products.Post("/", handler.Protected(), productHandler.Create) // Store owners only, implicitly verified by usecase mapping to user_id

	// Stores
	toko := api.Group("/toko")
	toko.Put("/:id_toko", handler.Protected(), storeHandler.Update)

	// Addresses
	address := api.Group("/address")
	address.Get("/", handler.Protected(), addressHandler.GetMyAddresses)
	address.Post("/", handler.Protected(), addressHandler.Create)

	// Transactions
	trx := api.Group("/trx")
	trx.Post("/checkout", handler.Protected(), transactionHandler.Checkout)
	trx.Get("/", handler.Protected(), transactionHandler.GetMyTransactions)
}
