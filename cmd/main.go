package main

import (
	"flag"
	"net/http"
	"thiagoluis88git/tech1/internal/adapters/driven/external"
	"thiagoluis88git/tech1/internal/adapters/driven/repositories"
	"thiagoluis88git/tech1/internal/adapters/driver/handler"
	"thiagoluis88git/tech1/internal/core/services"
	"thiagoluis88git/tech1/pkg/database"
	"thiagoluis88git/tech1/pkg/httpserver"
	"thiagoluis88git/tech1/pkg/responses"

	"github.com/go-chi/chi/v5"

	chiMiddleware "github.com/go-chi/chi/v5/middleware"
)

func main() {
	flag.Parse()

	db := database.ConfigDatabase()

	router := chi.NewRouter()
	router.Use(chiMiddleware.RequestID)
	router.Use(chiMiddleware.RealIP)
	router.Use(chiMiddleware.Recoverer)

	paymentRepo := repositories.NewPaymentRepository(db)
	paymentGateway := external.NewPaymentGateway()
	paymentService := services.NewPaymentService(paymentRepo, paymentGateway)

	productRepo := repositories.NewProductRepository(db)
	productService := services.NewProductService(productRepo)

	customerRepo := repositories.NewCustomerRepository(db)
	customerService := services.NewCustomerService(customerRepo)

	orderRepo := repositories.NewOrderRespository(db)
	orderService := services.NewOrderService(orderRepo, customerRepo, paymentService)

	router.Get("/api/health", func(w http.ResponseWriter, r *http.Request) {
		httpserver.SendResponseSuccess(w, &responses.BusinessResponse{
			StatusCode: 200,
			Message:    "ok",
		})
	})

	router.Post("/api/customer", handler.CreateCustomerHandler(customerService))
	router.Put("/api/customer/{id}", handler.UpdateCustomerHandler(customerService))
	router.Get("/api/customer/{cpf}", handler.GetCustomerByCPFHandler(customerService))

	router.Post("/api/products", handler.CreateProductHandler(productService))
	router.Delete("/api/products/{id}", handler.DeleteProductHandler(productService))
	router.Put("/api/products/{id}", handler.UpdateProductHandler(productService))
	router.Get("/api/products/categories", handler.GetCategoryHandler(productService))
	router.Get("/api/products/categories/{category}", handler.GetProductsByCategoryHandler(productService))

	router.Get("/api/payment/types", handler.GetPaymentTypeHandler(paymentService))
	router.Post("/api/order", handler.CreateOrderHandler(orderService))

	server := httpserver.New(router)
	server.Start()
}
