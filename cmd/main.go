package main

import (
	"fmt"
	"net/http"
	"thiagoluis88git/tech1/internal/adapters/driven/entities"
	"thiagoluis88git/tech1/internal/adapters/driven/repositories"
	"thiagoluis88git/tech1/internal/adapters/driver/handler"
	"thiagoluis88git/tech1/internal/core/services"
	"thiagoluis88git/tech1/pkg/httpserver"
	"thiagoluis88git/tech1/pkg/responses"

	"github.com/go-chi/chi/v5"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	chiMiddleware "github.com/go-chi/chi/v5/middleware"
)

func main() {
	dsn := "host=database user=fastfood password=fastfood1234 dbname=fastfood_db port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(fmt.Sprintf("could not open database: %v", err.Error()))
	}

	db.AutoMigrate(
		&entities.Customer{},
		&entities.Order{},
		&entities.OrderProduct{},
		&entities.PaymentOutbox{},
		&entities.Product{},
		&entities.ProducImage{},
		&entities.ProductCombo{},
	)

	router := chi.NewRouter()
	router.Use(chiMiddleware.RequestID)
	router.Use(chiMiddleware.RealIP)
	router.Use(chiMiddleware.Recoverer)

	productRepo := repositories.NewProductRepository(db)
	productService := services.NewProductService(productRepo)

	customerRepo := repositories.NewCustomerRepository(db)
	customerService := services.NewCustomerService(customerRepo)

	router.Get("/api/health", func(w http.ResponseWriter, r *http.Request) {
		httpserver.SendResponseSuccess(w, &responses.BusinessResponse{
			StatusCode: 200,
			Message:    "ok",
		})
	})

	router.Post("/api/customer", handler.CreateCustomerHandler(customerService))
	router.Post("/api/product", handler.CreateProductHandler(productService))
	router.Get("/api/category", handler.GetCategoryHandler(productService))

	server := httpserver.New(router)
	server.Start()
}
