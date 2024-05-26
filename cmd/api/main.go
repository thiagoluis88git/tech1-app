package main

import (
	"flag"
	"net/http"

	"github.com/thiagoluis88git/tech1/internal/adapters/driven/external"
	"github.com/thiagoluis88git/tech1/internal/adapters/driven/repositories"
	"github.com/thiagoluis88git/tech1/internal/adapters/driver/handler"
	"github.com/thiagoluis88git/tech1/internal/core/services"
	"github.com/thiagoluis88git/tech1/pkg/database"
	"github.com/thiagoluis88git/tech1/pkg/environment"
	"github.com/thiagoluis88git/tech1/pkg/httpserver"
	"github.com/thiagoluis88git/tech1/pkg/responses"

	"github.com/mvrilo/go-redoc"

	"github.com/go-chi/chi/v5"

	_ "github.com/thiagoluis88git/tech1/docs"

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

	doc := redoc.Redoc{
		Title:       "Example API",
		Description: "Example API Description",
		SpecFile:    *environment.RedocFolderPath,
		SpecPath:    "/docs/swagger.json",
		DocsPath:    "/docs",
	}

	db := database.ConfigDatabase()

	router := chi.NewRouter()
	router.Use(chiMiddleware.RequestID)
	router.Use(chiMiddleware.RealIP)
	router.Use(chiMiddleware.Recoverer)

	paymentRepo := repositories.NewPaymentRepository(db)
	paymentGateway := external.NewPaymentGateway()
	paymentService := services.NewPaymentService(paymentRepo, paymentGateway)

	productRepo := repositories.NewProductRepository(db)
	validateProductCategoryUseCase := services.NewValidateProductCategoryUseCase()
	productService := services.NewProductService(validateProductCategoryUseCase, productRepo)

	customerRepo := repositories.NewCustomerRepository(db)
	validateCPFUseCase := services.NewValidateCPFUseCase()
	customerService := services.NewCustomerService(validateCPFUseCase, customerRepo)

	orderRepo := repositories.NewOrderRespository(db)
	orderService := services.NewOrderService(orderRepo, customerRepo)

	router.Get("/api/health", func(w http.ResponseWriter, r *http.Request) {
		httpserver.SendResponseSuccess(w, &responses.BusinessResponse{
			StatusCode: 200,
			Message:    "ok",
		})
	})

	router.Post("/api/customers", handler.CreateCustomerHandler(customerService))
	router.Put("/api/customers/{id}", handler.UpdateCustomerHandler(customerService))
	router.Get("/api/customers/{id}", handler.GetCustomerByIdHandler(customerService))
	router.Post("/api/customers/login", handler.GetCustomerByCPFHandler(customerService))

	router.Post("/api/products", handler.CreateProductHandler(productService))
	router.Delete("/api/products/{id}", handler.DeleteProductHandler(productService))
	router.Put("/api/products/{id}", handler.UpdateProductHandler(productService))
	router.Get("/api/products/{id}", handler.GetProductsByIdHandler(productService))
	router.Get("/api/products/categories", handler.GetCategoriesHandler(productService))
	router.Get("/api/products/categories/{category}", handler.GetProductsByCategoryHandler(productService))

	router.Get("/api/payments/types", handler.GetPaymentTypeHandler(paymentService))
	router.Post("/api/payments", handler.CreatePaymentHandler(paymentService))

	router.Post("/api/orders", handler.CreateOrderHandler(orderService))
	router.Get("/api/orders/{id}", handler.GetOrderByIdHandler(orderService))
	router.Get("/api/orders/to-prepare", handler.GetOrdersToPrepareHandler(orderService))
	router.Get("/api/orders/follow", handler.GetOrdersToFollowHandler(orderService))
	router.Put("/api/orders/{id}/preparing", handler.UpdateOrderPreparingHandler(orderService))
	router.Put("/api/orders/{id}/done", handler.UpdateOrderDoneHandler(orderService))
	router.Put("/api/orders/{id}/delivered", handler.UpdateOrderDeliveredHandler(orderService))
	router.Put("/api/orders/{id}/not-delivered", handler.UpdateOrderNotDeliveredandler(orderService))

	router.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:3210/swagger/doc.json"),
	))

	go http.ListenAndServe(":3211", doc.Handler())

	server := httpserver.New(router)
	server.Start()
}
