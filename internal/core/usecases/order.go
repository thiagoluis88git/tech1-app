package usecases

import (
	"context"
	"sync"

	"github.com/thiagoluis88git/tech1/internal/core/domain"
	"github.com/thiagoluis88git/tech1/internal/core/ports"
	"github.com/thiagoluis88git/tech1/pkg/responses"
)

type CreateOrderUseCase struct {
	orderRepo                ports.OrderRepository
	customerRepo             ports.CustomerRepository
	validateToPrepare        *ValidateOrderToPrepareUseCase
	validateToDone           *ValidateOrderToDoneUseCase
	validateToDeliveredOrNot *ValidateOrderToDeliveredOrNotUseCase
	sortOrderUseCase         *SortOrdersUseCase
}

type UpdateToPreparingUseCase struct {
	orderRepo         ports.OrderRepository
	validateToPrepare *ValidateOrderToPrepareUseCase
}

type UpdateToDoneUseCase struct {
	orderRepo      ports.OrderRepository
	validateToDone *ValidateOrderToDoneUseCase
}

type UpdateToDeliveredUseCase struct {
	orderRepo                ports.OrderRepository
	validateToDeliveredOrNot *ValidateOrderToDeliveredOrNotUseCase
}

type UpdateToNotDeliveredUseCase struct {
	orderRepo                ports.OrderRepository
	validateToDeliveredOrNot *ValidateOrderToDeliveredOrNotUseCase
}

type GetOrderByIdUseCase struct {
	orderRepo ports.OrderRepository
}

type GetOrdersToPrepareUseCase struct {
	orderRepo        ports.OrderRepository
	sortOrderUseCase *SortOrdersUseCase
}

type GetGetOrdersToFollowUseCase struct {
	orderRepo        ports.OrderRepository
	sortOrderUseCase *SortOrdersUseCase
}

func NewCreateOrderUseCase(
	orderRepo ports.OrderRepository,
	customerRepo ports.CustomerRepository,
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
	orderRepo ports.OrderRepository,
) *GetOrderByIdUseCase {
	return &GetOrderByIdUseCase{
		orderRepo: orderRepo,
	}
}

func NewGetOrdersToPrepareUseCase(
	orderRepo ports.OrderRepository,
	sortOrderUseCase *SortOrdersUseCase,
) *GetOrdersToPrepareUseCase {
	return &GetOrdersToPrepareUseCase{
		orderRepo:        orderRepo,
		sortOrderUseCase: sortOrderUseCase,
	}
}

func NewGetOrdersToFollowUseCase(
	orderRepo ports.OrderRepository,
	sortOrderUseCase *SortOrdersUseCase,
) *GetGetOrdersToFollowUseCase {
	return &GetGetOrdersToFollowUseCase{
		orderRepo:        orderRepo,
		sortOrderUseCase: sortOrderUseCase,
	}
}

func NewUpdateToPreparingUseCase(
	orderRepo ports.OrderRepository,
	validateToPrepare *ValidateOrderToPrepareUseCase,
) *UpdateToPreparingUseCase {
	return &UpdateToPreparingUseCase{
		orderRepo:         orderRepo,
		validateToPrepare: validateToPrepare,
	}
}

func NewUpdateToDoneUseCase(
	orderRepo ports.OrderRepository,
	validateToDone *ValidateOrderToDoneUseCase,
) *UpdateToDoneUseCase {
	return &UpdateToDoneUseCase{
		orderRepo:      orderRepo,
		validateToDone: validateToDone,
	}
}

func NewUpdateToDeliveredUseCase(
	orderRepo ports.OrderRepository,
	validateToDeliveredOrNot *ValidateOrderToDeliveredOrNotUseCase,
) *UpdateToDeliveredUseCase {
	return &UpdateToDeliveredUseCase{
		orderRepo:                orderRepo,
		validateToDeliveredOrNot: validateToDeliveredOrNot,
	}
}

func NewUpdateToNotDeliveredUseCase(
	orderRepo ports.OrderRepository,
	validateToDeliveredOrNot *ValidateOrderToDeliveredOrNotUseCase,
) *UpdateToNotDeliveredUseCase {
	return &UpdateToNotDeliveredUseCase{
		orderRepo:                orderRepo,
		validateToDeliveredOrNot: validateToDeliveredOrNot,
	}
}

func (usecase *CreateOrderUseCase) Execute(ctx context.Context, order domain.Order, date int64, wg *sync.WaitGroup, ch chan bool) (domain.OrderResponse, error) {
	//Block this code below until this Channel be empty (by reading with <-ch)
	ch <- true

	order.TicketNumber = usecase.GenerateTicket(ctx, date)

	response, err := usecase.orderRepo.CreateOrder(ctx, order)

	if err != nil {
		return domain.OrderResponse{}, responses.GetResponseError(err, "OrderService -> CreateOrder")
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

func (usecase *GetOrderByIdUseCase) Execute(ctx context.Context, orderId uint) (domain.OrderResponse, error) {
	response, err := usecase.orderRepo.GetOrderById(ctx, orderId)

	if err != nil {
		return domain.OrderResponse{}, responses.GetResponseError(err, "OrderService -> GetOrderById")
	}

	return response, nil
}

func (usecase *GetOrdersToPrepareUseCase) Execute(ctx context.Context) ([]domain.OrderResponse, error) {
	response, err := usecase.orderRepo.GetOrdersToPrepare(ctx)

	if err != nil {
		return []domain.OrderResponse{}, responses.GetResponseError(err, "OrderService -> GetOrdersToPrepare")
	}

	usecase.sortOrderUseCase.Execute(response)

	return response, nil
}

func (usecase *GetGetOrdersToFollowUseCase) Execute(ctx context.Context) ([]domain.OrderResponse, error) {
	response, err := usecase.orderRepo.GetOrdersToFollow(ctx)

	if err != nil {
		return []domain.OrderResponse{}, responses.GetResponseError(err, "OrderService -> GetOrdersStatus")
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
