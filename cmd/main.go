package main

import (
	"flag"
	"fmt"
	"net/http"
	"thiagoluis88git/tech1/internal/adapters/driven/entities"
	"thiagoluis88git/tech1/internal/adapters/driven/repositories"
	"thiagoluis88git/tech1/internal/adapters/driver/handler"
	"thiagoluis88git/tech1/internal/core/services"
	"thiagoluis88git/tech1/pkg/environment"
	"thiagoluis88git/tech1/pkg/httpserver"
	"thiagoluis88git/tech1/pkg/responses"

	"github.com/go-chi/chi/v5"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	chiMiddleware "github.com/go-chi/chi/v5/middleware"
)

func main() {
	flag.Parse()

	dsn := fmt.Sprintf("host=%v user=fastfood password=fastfood1234 dbname=fastfood_db port=5432 sslmode=disable", *environment.DbHost)
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
		&entities.ProductImage{},
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

	orderRepo := repositories.NewOrderRespository(db)
	orderService := services.NewOrderService(orderRepo)

	router.Get("/api/health", func(w http.ResponseWriter, r *http.Request) {
		httpserver.SendResponseSuccess(w, &responses.BusinessResponse{
			StatusCode: 200,
			Message:    "ok",
		})
	})

	router.Post("/api/customer", handler.CreateCustomerHandler(customerService))
	router.Get("/api/customer/{cpf}", handler.GetCustomerByCPFHandler(customerService))

	router.Post("/api/products", handler.CreateProductHandler(productService))
	router.Get("/api/products/categories", handler.GetCategoryHandler(productService))
	router.Get("/api/products/categories/{category}", handler.GetProductsByCategory(productService))

	router.Post("/api/order", handler.CreateOrderHandler(orderService))

	server := httpserver.New(router)
	server.Start()
}
