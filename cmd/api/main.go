package main

import (
	"net/http"

	"github.com/thiagoluis88git/tech1/internal/adapters/driven/external"
	"github.com/thiagoluis88git/tech1/internal/adapters/driven/external/remote"
	extRepo "github.com/thiagoluis88git/tech1/internal/adapters/driven/external/repositories"
	"github.com/thiagoluis88git/tech1/internal/adapters/driven/repositories"
	"github.com/thiagoluis88git/tech1/internal/adapters/driver/handler"
	"github.com/thiagoluis88git/tech1/internal/adapters/driver/webhook"
	"github.com/thiagoluis88git/tech1/internal/core/usecases"
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
	environment.LoadEnvironmentVariables()

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

	httpClient := httpserver.NewHTTPClient()

	paymentRepo := repositories.NewPaymentRepository(db)
	paymentGateway := external.NewPaymentGateway()
	payOrderUseCase := usecases.NewPayOrderUseCase(paymentRepo, paymentGateway)
	getPaymentTypesUseCase := usecases.NewGetPaymentTypesUseCasee(paymentRepo)

	productRepo := repositories.NewProductRepository(db)
	validateProductCategoryUseCase := usecases.NewValidateProductCategoryUseCase()
	getCategoriesUseCase := usecases.NewGetCategoriesUseCase(productRepo)
	getProductsUseCase := usecases.NewGetProductsByCategoryUseCase(productRepo)
	getProductByIdUseCase := usecases.NewGetProductByIdUseCase(productRepo)
	deleteProductUseCase := usecases.NewDeleteProductUseCase(productRepo)
	updateProductUseCase := usecases.NewUpdateProductUseCase(productRepo)
	createProductUseCase := usecases.NewCreateProductUseCase(validateProductCategoryUseCase, productRepo)

	customerRepo := repositories.NewCustomerRepository(db)
	validateCPFUseCase := usecases.NewValidateCPFUseCase()
	createCustomerUseCase := usecases.NewCreateCustomerUseCase(validateCPFUseCase, customerRepo)
	updateCustomerUseCase := usecases.NewUpdateCustomerUseCase(validateCPFUseCase, customerRepo)
	getCustomerByIdUseCase := usecases.NewGetCustomerByIdUseCase(customerRepo)
	getCustomerByCPFUseCase := usecases.NewGetCustomerByCPFUseCase(validateCPFUseCase, customerRepo)

	orderRepo := repositories.NewOrderRespository(db)
	validateToPreare := usecases.NewValidateOrderToPrepareUseCase(orderRepo)
	validateToDone := usecases.NewValidateOrderToDoneUseCase(orderRepo)
	validateToDeliveredOrNot := usecases.NewValidateOrderToDeliveredOrNotUseCase(orderRepo)
	sortOrders := usecases.NewSortOrdersUseCase()
	createOrderUseCase := usecases.NewCreateOrderUseCase(
		orderRepo,
		customerRepo,
		validateToPreare,
		validateToDone,
		validateToDeliveredOrNot,
		sortOrders,
	)
	getOrderByIdUseCase := usecases.NewGetOrderByIdUseCase(orderRepo)
	getOrdersToPrepareUseCase := usecases.NewGetOrdersToPrepareUseCase(
		orderRepo,
		sortOrders,
	)
	getOrdersToFollowUseCase := usecases.NewGetOrdersToFollowUseCase(
		orderRepo,
		sortOrders,
	)
	getOrdersWaitingPaymentUseCase := usecases.NewGetOrdersWaitingPaymentUseCase(
		orderRepo,
		sortOrders,
	)
	updateToPreparingUseCase := usecases.NewUpdateToPreparingUseCase(
		orderRepo,
		validateToPreare,
	)
	updateToDoneUseCase := usecases.NewUpdateToDoneUseCase(
		orderRepo,
		validateToDone,
	)
	updateToDeliveredUseCase := usecases.NewUpdateToDeliveredUseCase(
		orderRepo,
		validateToDeliveredOrNot,
	)
	updateToNotDeliveredUseCase := usecases.NewUpdateToNotDeliveredUseCase(
		orderRepo,
		validateToDeliveredOrNot,
	)

	qrCodeRemoteDataSource := remote.NewMercadoLivreDataSource(httpClient)
	extQRCodeGeneratorRepository := extRepo.NewMercadoLivreRepository(qrCodeRemoteDataSource)
	generateQRCodePaymentUseCase := usecases.NewGenerateQRCodePaymentUseCase(
		extQRCodeGeneratorRepository,
		orderRepo,
		paymentRepo,
	)

	finishOrderForQRCodeUseCase := usecases.NewFinishOrderForQRCodeUseCase(
		extQRCodeGeneratorRepository,
		orderRepo,
		paymentRepo,
	)

	router.Get("/api/health", func(w http.ResponseWriter, r *http.Request) {
		httpserver.SendResponseSuccess(w, &responses.BusinessResponse{
			StatusCode: 200,
			Message:    "ok",
		})
	})

	router.Post("/api/qrcode/generate", handler.GenerateQRCodeHandler(generateQRCodePaymentUseCase))
	router.Post("/api/webhook/ml/payment", webhook.PostMercadoLivreWebhook(finishOrderForQRCodeUseCase))

	router.Post("/api/customers", handler.CreateCustomerHandler(createCustomerUseCase))
	router.Put("/api/customers/{id}", handler.UpdateCustomerHandler(updateCustomerUseCase))
	router.Get("/api/customers/{id}", handler.GetCustomerByIdHandler(getCustomerByIdUseCase))
	router.Post("/api/customers/login", handler.GetCustomerByCPFHandler(getCustomerByCPFUseCase))

	router.Post("/api/products", handler.CreateProductHandler(createProductUseCase))
	router.Delete("/api/products/{id}", handler.DeleteProductHandler(deleteProductUseCase))
	router.Put("/api/products/{id}", handler.UpdateProductHandler(updateProductUseCase))
	router.Get("/api/products/{id}", handler.GetProductsByIdHandler(getProductByIdUseCase))
	router.Get("/api/products/categories", handler.GetCategoriesHandler(getCategoriesUseCase))
	router.Get("/api/products/categories/{category}", handler.GetProductsByCategoryHandler(getProductsUseCase))

	router.Get("/api/payments/types", handler.GetPaymentTypeHandler(getPaymentTypesUseCase))
	router.Post("/api/payments", handler.CreatePaymentHandler(payOrderUseCase))

	router.Post("/api/orders", handler.CreateOrderHandler(createOrderUseCase))
	router.Get("/api/orders/{id}", handler.GetOrderByIdHandler(getOrderByIdUseCase))
	router.Get("/api/orders/to-prepare", handler.GetOrdersToPrepareHandler(getOrdersToPrepareUseCase))
	router.Get("/api/orders/follow", handler.GetOrdersToFollowHandler(getOrdersToFollowUseCase))
	router.Get("/api/orders/waiting-payment", handler.GetOrdersWaitingPaymentHandler(getOrdersWaitingPaymentUseCase))
	router.Put("/api/orders/{id}/preparing", handler.UpdateOrderPreparingHandler(updateToPreparingUseCase))
	router.Put("/api/orders/{id}/done", handler.UpdateOrderDoneHandler(updateToDoneUseCase))
	router.Put("/api/orders/{id}/delivered", handler.UpdateOrderDeliveredHandler(updateToDeliveredUseCase))
	router.Put("/api/orders/{id}/not-delivered", handler.UpdateOrderNotDeliveredandler(updateToNotDeliveredUseCase))

	router.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:3210/swagger/doc.json"),
	))

	go http.ListenAndServe(":3211", doc.Handler())

	server := httpserver.New(router)
	server.Start()
}
