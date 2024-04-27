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

	_ "thiagoluis88git/tech1/docs"

	chiMiddleware "github.com/go-chi/chi/v5/middleware"
	httpSwagger "github.com/swaggo/http-swagger/v2"
)

// @title Tech1 API Docs
// @version 1.0
// @description This is the API for the Tech1 Fiap Project.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localshot:3210
// @BasePath /
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
	router.Post("/api/products/combo", handler.CreateComboHandler(productService))
	router.Get("/api/products/combo", handler.GetCombosHandler(productService))
	router.Delete("/api/products/{id}", handler.DeleteProductHandler(productService))
	router.Put("/api/products/{id}", handler.UpdateProductHandler(productService))
	router.Get("/api/products/categories", handler.GetCategoryHandler(productService))
	router.Get("/api/products/categories/{category}", handler.GetProductsByCategoryHandler(productService))

	router.Get("/api/payment/types", handler.GetPaymentTypeHandler(paymentService))
	router.Post("/api/order", handler.CreateOrderHandler(orderService))
	router.Put("/api/order/{id}/preparing", handler.UpdateOrderPreparingHandler(orderService))
	router.Put("/api/order/{id}/done", handler.UpdateOrderDoneHandler(orderService))
	router.Put("/api/order/{id}/delivered", handler.UpdateOrderDeliveredHandler(orderService))
	router.Put("/api/order/{id}/not-delivered", handler.UpdateOrderNotDeliveredandler(orderService))

	router.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:3210/swagger/doc.json"),
	))

	server := httpserver.New(router)
	server.Start()
}
