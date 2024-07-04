package usecases

import (
	"context"
	"sync"

	"github.com/thiagoluis88git/tech1/internal/core/domain"
	"github.com/thiagoluis88git/tech1/internal/core/ports"
	"github.com/thiagoluis88git/tech1/pkg/responses"
)

type OrderService struct {
	orderRepo                ports.OrderRepository
	customerRepo             ports.CustomerRepository
	validateToPrepare        *ValidateOrderToPrepareUseCase
	validateToDone           *ValidateOrderToDoneUseCase
	validateToDeliveredOrNot *ValidateOrderToDeliveredOrNotUseCase
	sortOrderUseCase         *SortOrdersUseCase
}

func NewOrderService(
	orderRepo ports.OrderRepository,
	customerRepo ports.CustomerRepository,
	validateToPrepate *ValidateOrderToPrepareUseCase,
	validateToDone *ValidateOrderToDoneUseCase,
	validateToDeliveredOrNot *ValidateOrderToDeliveredOrNotUseCase,
	sortOrderUseCase *SortOrdersUseCase,
) *OrderService {
	return &OrderService{
		orderRepo:                orderRepo,
		customerRepo:             customerRepo,
		validateToPrepare:        validateToPrepate,
		validateToDone:           validateToDone,
		validateToDeliveredOrNot: validateToDeliveredOrNot,
		sortOrderUseCase:         sortOrderUseCase,
	}
}

func (service *OrderService) CreateOrder(ctx context.Context, order domain.Order, date int64, wg *sync.WaitGroup, ch chan bool) (domain.OrderResponse, error) {
	//Block this code below until this Channel be empty (by reading with <-ch)
	ch <- true

	order.TicketNumber = service.GenerateTicket(ctx, date)

	response, err := service.orderRepo.CreateOrder(ctx, order)

	if err != nil {
		return domain.OrderResponse{}, responses.GetResponseError(err, "OrderService -> CreateOrder")
	}

	if order.CustomerID != nil {
		customer, err := service.customerRepo.GetCustomerById(ctx, *order.CustomerID)
		if err == nil {
			response.CustomerName = &customer.Name
		}
	}

	// Release the channel to others process be able to start a new order creation
	<-ch
	wg.Done()

	return response, nil
}

func (service *OrderService) GenerateTicket(ctx context.Context, date int64) int {
	return service.orderRepo.GetNextTicketNumber(ctx, date)
}

func (service *OrderService) GetOrderById(ctx context.Context, orderId uint) (domain.OrderResponse, error) {
	response, err := service.orderRepo.GetOrderById(ctx, orderId)

	if err != nil {
		return domain.OrderResponse{}, responses.GetResponseError(err, "OrderService -> GetOrderById")
	}

	return response, nil
}

func (service *OrderService) GetOrdersToPrepare(ctx context.Context) ([]domain.OrderResponse, error) {
	response, err := service.orderRepo.GetOrdersToPrepare(ctx)

	if err != nil {
		return []domain.OrderResponse{}, responses.GetResponseError(err, "OrderService -> GetOrdersToPrepare")
	}

	service.sortOrderUseCase.Execute(response)

	return response, nil
}

func (service *OrderService) GetOrdersToFollow(ctx context.Context) ([]domain.OrderResponse, error) {
	response, err := service.orderRepo.GetOrdersToFollow(ctx)

	if err != nil {
		return []domain.OrderResponse{}, responses.GetResponseError(err, "OrderService -> GetOrdersStatus")
	}

	service.sortOrderUseCase.Execute(response)

	return response, nil
}

func (service *OrderService) UpdateToPreparing(ctx context.Context, orderId uint) error {
	err := service.validateToPrepare.Execute(ctx, orderId)

	if err != nil {
		return responses.GetResponseError(err, "OrderService -> UpdateToPreparing")
	}

	err = service.orderRepo.UpdateToPreparing(ctx, orderId)

	if err != nil {
		return responses.GetResponseError(err, "OrderService -> UpdateToPreparing")
	}

	return nil
}

func (service *OrderService) UpdateToDone(ctx context.Context, orderId uint) error {
	err := service.validateToDone.Execute(ctx, orderId)

	if err != nil {
		return responses.GetResponseError(err, "OrderService -> UpdateToDone")
	}

	err = service.orderRepo.UpdateToDone(ctx, orderId)

	if err != nil {
		return responses.GetResponseError(err, "OrderService -> UpdateToDone")
	}

	return nil
}

func (service *OrderService) UpdateToDelivered(ctx context.Context, orderId uint) error {
	err := service.validateToDeliveredOrNot.Execute(ctx, orderId)

	if err != nil {
		return responses.GetResponseError(err, "OrderService -> UpdateToDelivered")
	}

	err = service.orderRepo.UpdateToDelivered(ctx, orderId)

	if err != nil {
		return responses.GetResponseError(err, "OrderService -> UpdateToDelivered")
	}

	return nil
}

func (service *OrderService) UpdateToNotDelivered(ctx context.Context, orderId uint) error {
	err := service.validateToDeliveredOrNot.Execute(ctx, orderId)

	if err != nil {
		return responses.GetResponseError(err, "OrderService -> UpdateToNotDelivered")
	}

	err = service.orderRepo.UpdateToNotDelivered(ctx, orderId)

	if err != nil {
		return responses.GetResponseError(err, "OrderService -> UpdateToNotDelivered")
	}

	return nil
}
