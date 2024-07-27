package usecases

import (
	"context"
	"sync"

	"github.com/thiagoluis88git/tech1/internal/core/domain/dto"
	"github.com/thiagoluis88git/tech1/internal/core/domain/repository"
	"github.com/thiagoluis88git/tech1/pkg/responses"
)

type CreateOrderUseCase struct {
	orderRepo                repository.OrderRepository
	customerRepo             repository.CustomerRepository
	validateToPrepare        *ValidateOrderToPrepareUseCase
	validateToDone           *ValidateOrderToDoneUseCase
	validateToDeliveredOrNot *ValidateOrderToDeliveredOrNotUseCase
	sortOrderUseCase         *SortOrdersUseCase
}

type UpdateToPreparingUseCase struct {
	orderRepo         repository.OrderRepository
	validateToPrepare *ValidateOrderToPrepareUseCase
}

type UpdateToDoneUseCase struct {
	orderRepo      repository.OrderRepository
	validateToDone *ValidateOrderToDoneUseCase
}

type UpdateToDeliveredUseCase struct {
	orderRepo                repository.OrderRepository
	validateToDeliveredOrNot *ValidateOrderToDeliveredOrNotUseCase
}

type UpdateToNotDeliveredUseCase struct {
	orderRepo                repository.OrderRepository
	validateToDeliveredOrNot *ValidateOrderToDeliveredOrNotUseCase
}

type GetOrderByIdUseCase struct {
	orderRepo repository.OrderRepository
}

type GetOrdersToPrepareUseCase struct {
	orderRepo        repository.OrderRepository
	sortOrderUseCase *SortOrdersUseCase
}

type GetOrdersToFollowUseCase struct {
	orderRepo        repository.OrderRepository
	sortOrderUseCase *SortOrdersUseCase
}

type GetOrdersWaitingPaymentUseCase struct {
	orderRepo        repository.OrderRepository
	sortOrderUseCase *SortOrdersUseCase
}

func NewCreateOrderUseCase(
	orderRepo repository.OrderRepository,
	customerRepo repository.CustomerRepository,
	validateToPrepate *ValidateOrderToPrepareUseCase,
	validateToDone *ValidateOrderToDoneUseCase,
	validateToDeliveredOrNot *ValidateOrderToDeliveredOrNotUseCase,
	sortOrderUseCase *SortOrdersUseCase,
) *CreateOrderUseCase {
	return &CreateOrderUseCase{
		orderRepo:                orderRepo,
		customerRepo:             customerRepo,
		validateToPrepare:        validateToPrepate,
		validateToDone:           validateToDone,
		validateToDeliveredOrNot: validateToDeliveredOrNot,
		sortOrderUseCase:         sortOrderUseCase,
	}
}

func NewGetOrderByIdUseCase(
	orderRepo repository.OrderRepository,
) *GetOrderByIdUseCase {
	return &GetOrderByIdUseCase{
		orderRepo: orderRepo,
	}
}

func NewGetOrdersToPrepareUseCase(
	orderRepo repository.OrderRepository,
	sortOrderUseCase *SortOrdersUseCase,
) *GetOrdersToPrepareUseCase {
	return &GetOrdersToPrepareUseCase{
		orderRepo:        orderRepo,
		sortOrderUseCase: sortOrderUseCase,
	}
}

func NewGetOrdersToFollowUseCase(
	orderRepo repository.OrderRepository,
	sortOrderUseCase *SortOrdersUseCase,
) *GetOrdersToFollowUseCase {
	return &GetOrdersToFollowUseCase{
		orderRepo:        orderRepo,
		sortOrderUseCase: sortOrderUseCase,
	}
}

func NewGetOrdersWaitingPaymentUseCase(
	orderRepo repository.OrderRepository,
	sortOrderUseCase *SortOrdersUseCase,
) *GetOrdersWaitingPaymentUseCase {
	return &GetOrdersWaitingPaymentUseCase{
		orderRepo:        orderRepo,
		sortOrderUseCase: sortOrderUseCase,
	}
}

func NewUpdateToPreparingUseCase(
	orderRepo repository.OrderRepository,
	validateToPrepare *ValidateOrderToPrepareUseCase,
) *UpdateToPreparingUseCase {
	return &UpdateToPreparingUseCase{
		orderRepo:         orderRepo,
		validateToPrepare: validateToPrepare,
	}
}

func NewUpdateToDoneUseCase(
	orderRepo repository.OrderRepository,
	validateToDone *ValidateOrderToDoneUseCase,
) *UpdateToDoneUseCase {
	return &UpdateToDoneUseCase{
		orderRepo:      orderRepo,
		validateToDone: validateToDone,
	}
}

func NewUpdateToDeliveredUseCase(
	orderRepo repository.OrderRepository,
	validateToDeliveredOrNot *ValidateOrderToDeliveredOrNotUseCase,
) *UpdateToDeliveredUseCase {
	return &UpdateToDeliveredUseCase{
		orderRepo:                orderRepo,
		validateToDeliveredOrNot: validateToDeliveredOrNot,
	}
}

func NewUpdateToNotDeliveredUseCase(
	orderRepo repository.OrderRepository,
	validateToDeliveredOrNot *ValidateOrderToDeliveredOrNotUseCase,
) *UpdateToNotDeliveredUseCase {
	return &UpdateToNotDeliveredUseCase{
		orderRepo:                orderRepo,
		validateToDeliveredOrNot: validateToDeliveredOrNot,
	}
}

func (usecase *CreateOrderUseCase) Execute(ctx context.Context, order dto.Order, date int64, wg *sync.WaitGroup, ch chan bool) (dto.OrderResponse, error) {
	//Block this code below until this Channel be empty (by reading with <-ch)
	ch <- true

	order.TicketNumber = usecase.GenerateTicket(ctx, date)

	response, err := usecase.orderRepo.CreateOrder(ctx, order)

	if err != nil {
		return dto.OrderResponse{}, responses.GetResponseError(err, "OrderService -> CreateOrder")
	}

	if order.CustomerID != nil {
		customer, err := usecase.customerRepo.GetCustomerById(ctx, *order.CustomerID)
		if err == nil {
			response.CustomerName = &customer.Name
		}
	}

	// Release the channel to others process be able to start a new order creation
	<-ch
	wg.Done()

	return response, nil
}

func (usecase *CreateOrderUseCase) GenerateTicket(ctx context.Context, date int64) int {
	return usecase.orderRepo.GetNextTicketNumber(ctx, date)
}

func (usecase *GetOrderByIdUseCase) Execute(ctx context.Context, orderId uint) (dto.OrderResponse, error) {
	response, err := usecase.orderRepo.GetOrderById(ctx, orderId)

	if err != nil {
		return dto.OrderResponse{}, responses.GetResponseError(err, "OrderService -> GetOrderById")
	}

	return response, nil
}

func (usecase *GetOrdersToPrepareUseCase) Execute(ctx context.Context) ([]dto.OrderResponse, error) {
	response, err := usecase.orderRepo.GetOrdersToPrepare(ctx)

	if err != nil {
		return []dto.OrderResponse{}, responses.GetResponseError(err, "OrderService -> GetOrdersToPrepare")
	}

	usecase.sortOrderUseCase.Execute(response)

	return response, nil
}

func (usecase *GetOrdersToFollowUseCase) Execute(ctx context.Context) ([]dto.OrderResponse, error) {
	response, err := usecase.orderRepo.GetOrdersToFollow(ctx)

	if err != nil {
		return []dto.OrderResponse{}, responses.GetResponseError(err, "OrderService -> GetOrdersToFollow")
	}

	usecase.sortOrderUseCase.Execute(response)

	return response, nil
}

func (usecase *GetOrdersWaitingPaymentUseCase) Execute(ctx context.Context) ([]dto.OrderResponse, error) {
	response, err := usecase.orderRepo.GetOrdersWaitingPayment(ctx)

	if err != nil {
		return []dto.OrderResponse{}, responses.GetResponseError(err, "OrderService -> GetOrdersWaitingPayment")
	}

	usecase.sortOrderUseCase.Execute(response)

	return response, nil
}

func (usecase *UpdateToPreparingUseCase) Execute(ctx context.Context, orderId uint) error {
	err := usecase.validateToPrepare.Execute(ctx, orderId)

	if err != nil {
		return responses.GetResponseError(err, "OrderService -> UpdateToPreparing")
	}

	err = usecase.orderRepo.UpdateToPreparing(ctx, orderId)

	if err != nil {
		return responses.GetResponseError(err, "OrderService -> UpdateToPreparing")
	}

	return nil
}

func (usecase *UpdateToDoneUseCase) Execute(ctx context.Context, orderId uint) error {
	err := usecase.validateToDone.Execute(ctx, orderId)

	if err != nil {
		return responses.GetResponseError(err, "OrderService -> UpdateToDone")
	}

	err = usecase.orderRepo.UpdateToDone(ctx, orderId)

	if err != nil {
		return responses.GetResponseError(err, "OrderService -> UpdateToDone")
	}

	return nil
}

func (usecase *UpdateToDeliveredUseCase) Execute(ctx context.Context, orderId uint) error {
	err := usecase.validateToDeliveredOrNot.Execute(ctx, orderId)

	if err != nil {
		return responses.GetResponseError(err, "OrderService -> UpdateToDelivered")
	}

	err = usecase.orderRepo.UpdateToDelivered(ctx, orderId)

	if err != nil {
		return responses.GetResponseError(err, "OrderService -> UpdateToDelivered")
	}

	return nil
}

func (usecase *UpdateToNotDeliveredUseCase) Execute(ctx context.Context, orderId uint) error {
	err := usecase.validateToDeliveredOrNot.Execute(ctx, orderId)

	if err != nil {
		return responses.GetResponseError(err, "OrderService -> UpdateToNotDelivered")
	}

	err = usecase.orderRepo.UpdateToNotDelivered(ctx, orderId)

	if err != nil {
		return responses.GetResponseError(err, "OrderService -> UpdateToNotDelivered")
	}

	return nil
}
