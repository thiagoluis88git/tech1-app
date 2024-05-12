package repositories

import (
	"context"
	"fmt"
	"time"

	"github.com/thiagoluis88git/tech1/internal/adapters/driven/entities"
	"github.com/thiagoluis88git/tech1/internal/core/domain"
	"github.com/thiagoluis88git/tech1/internal/core/ports"
	"github.com/thiagoluis88git/tech1/pkg/responses"

	"gorm.io/gorm"
)

type OrderRespository struct {
	db *gorm.DB
}

func NewOrderRespository(db *gorm.DB) ports.OrderRepository {
	return &OrderRespository{
		db: db,
	}
}

func (repository *OrderRespository) CreateOrder(ctx context.Context, order domain.Order) (domain.OrderResponse, error) {
	tx := repository.db.WithContext(ctx).Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return domain.OrderResponse{}, responses.GetDatabaseError(err)
	}

	orderEntity := &entities.Order{
		OrderStatus:  entities.OrderStatusCreated,
		TotalPrice:   order.TotalPrice,
		CustomerID:   order.CustomerID,
		PaymentID:    order.PaymentID,
		TicketNumber: order.TicketNumber,
	}

	err := tx.Create(orderEntity).Error

	if err != nil {
		tx.Rollback()
		return domain.OrderResponse{}, responses.GetDatabaseError(err)
	}

	orderProductsEntity := []*entities.OrderProduct{}

	for _, value := range order.OrderProduct {
		orderProductsEntity = append(orderProductsEntity, &entities.OrderProduct{
			ProductID: value.ProductID,
			OrderID:   orderEntity.ID,
		})
	}

	err = tx.Create(orderProductsEntity).Error

	if err != nil {
		tx.Rollback()
		return domain.OrderResponse{}, responses.GetDatabaseError(err)
	}

	err = tx.Commit().Error

	if err != nil {
		tx.Rollback()
		return domain.OrderResponse{}, responses.GetDatabaseError(err)
	}

	return domain.OrderResponse{
		OrderId:      orderEntity.ID,
		OrderDate:    orderEntity.CreatedAt,
		TicketNumber: orderEntity.TicketNumber,
	}, nil
}

func (repository *OrderRespository) GetOrderById(ctx context.Context, orderId uint) (domain.OrderResponse, error) {
	var orderEntity entities.Order
	err := repository.
		db.WithContext(ctx).
		Model(&entities.Order{}).
		Preload("OrderProduct").
		Preload("Customer").
		Where("id = ?", orderId).
		Find(&orderEntity).
		Limit(1).
		Error

	if err != nil {
		return domain.OrderResponse{}, responses.GetDatabaseError(err)
	}

	orderProduct := []domain.OrderProductResponse{}

	for _, value := range orderEntity.OrderProduct {
		orderProduct = append(orderProduct, domain.OrderProductResponse{
			ProductID:   orderId,
			ProductName: fmt.Sprintf("FALTA BUSCAR: %d", value.ID),
		})
	}

	var customerName *string

	if orderEntity.Customer != nil {
		customerName = &orderEntity.Customer.Name
	}

	return domain.OrderResponse{
		OrderId:      orderId,
		OrderDate:    orderEntity.CreatedAt,
		TicketNumber: orderEntity.TicketNumber,
		OrderStatus:  orderEntity.OrderStatus,
		OrderProduct: orderProduct,
		CustomerName: customerName,
	}, nil
}

func (repository *OrderRespository) GetOrdersToPrepare(ctx context.Context) ([]domain.OrderResponse, error) {
	var orderEntity []entities.Order
	err := repository.
		db.WithContext(ctx).
		Model(&entities.Order{}).
		Preload("OrderProduct").
		Preload("Customer").
		Where("order_status = ?", entities.OrderStatusCreated).
		Find(&orderEntity).
		Error

	if err != nil {
		return []domain.OrderResponse{}, responses.GetDatabaseError(err)
	}

	orders := []domain.OrderResponse{}

	for _, value := range orderEntity {
		orderProduct := []domain.OrderProductResponse{}

		for _, value := range value.OrderProduct {
			orderProduct = append(orderProduct, domain.OrderProductResponse{
				ProductID:   value.ProductID,
				ProductName: fmt.Sprintf("FALTA BUSCAR: %d", value.ID),
			})
		}

		var customerName *string

		if value.Customer != nil {
			customerName = &value.Customer.Name
		}

		orders = append(orders, domain.OrderResponse{
			OrderId:      value.ID,
			OrderDate:    value.CreatedAt,
			TicketNumber: value.TicketNumber,
			OrderStatus:  value.OrderStatus,
			OrderProduct: orderProduct,
			CustomerName: customerName,
		})
	}

	return orders, nil
}

func (repository *OrderRespository) GetOrdersStatus(ctx context.Context) ([]domain.OrderResponse, error) {
	var orderEntity []entities.Order
	err := repository.
		db.WithContext(ctx).
		Model(&entities.Order{}).
		Preload("OrderProduct").
		Preload("Customer").
		Where("order_status in ?", entities.OrderStatusPreparing, entities.OrderStatusDone).
		Find(&orderEntity).
		Error

	if err != nil {
		return []domain.OrderResponse{}, responses.GetDatabaseError(err)
	}

	orders := []domain.OrderResponse{}

	for _, value := range orderEntity {
		orderProduct := []domain.OrderProductResponse{}

		for _, value := range value.OrderProduct {
			orderProduct = append(orderProduct, domain.OrderProductResponse{
				ProductID:   value.ProductID,
				ProductName: fmt.Sprintf("FALTA BUSCAR: %d", value.ID),
			})
		}

		var customerName *string

		if value.Customer != nil {
			customerName = &value.Customer.Name
		}

		orders = append(orders, domain.OrderResponse{
			OrderId:      value.ID,
			OrderDate:    value.CreatedAt,
			TicketNumber: value.TicketNumber,
			OrderStatus:  value.OrderStatus,
			OrderProduct: orderProduct,
			CustomerName: customerName,
		})
	}

	return orders, nil
}

func (repository *OrderRespository) UpdateToPreparing(ctx context.Context, orderId uint) error {
	err := repository.db.WithContext(ctx).
		Model(&entities.Order{}).
		Where("id = ?", orderId).
		Update("order_status", entities.OrderStatusPreparing).
		Update("preparing_at", time.Now()).
		Error

	if err != nil {
		return responses.GetDatabaseError(err)
	}

	return nil
}

func (repository *OrderRespository) UpdateToDone(ctx context.Context, orderId uint) error {
	err := repository.db.WithContext(ctx).
		Model(&entities.Order{}).
		Where("id = ?", orderId).
		Update("order_status", entities.OrderStatusDone).
		Update("done_at", time.Now()).
		Error

	if err != nil {
		return responses.GetDatabaseError(err)
	}

	return nil
}

func (repository *OrderRespository) UpdateToDelivered(ctx context.Context, orderId uint) error {
	err := repository.db.WithContext(ctx).
		Model(&entities.Order{}).
		Where("id = ?", orderId).
		Update("order_status", entities.OrderStatusDelivered).
		Update("delivered_at", time.Now()).
		Error

	if err != nil {
		return responses.GetDatabaseError(err)
	}

	return nil
}

func (repository *OrderRespository) UpdateToNotDelivered(ctx context.Context, orderId uint) error {
	err := repository.db.WithContext(ctx).
		Model(&entities.Order{}).
		Where("id = ?", orderId).
		Update("order_status", entities.OrderStatusNotDelivered).
		Update("not_delivered_at", time.Now()).
		Error

	if err != nil {
		return responses.GetDatabaseError(err)
	}

	return nil
}

func (repository *OrderRespository) GetNextTicketNumber(ctx context.Context, date int64) int {
	var orderTicketNumber entities.OrderTicketNumber
	err := repository.db.WithContext(ctx).
		Model(&entities.OrderTicketNumber{}).
		Where("date = ?", date).
		Find(&orderTicketNumber).
		Limit(1).
		Error

	if err != nil || orderTicketNumber.Date == 0 {
		return repository.createNewTicketForDate(ctx, date)
	}

	newTicketNumber := orderTicketNumber.TicketNumber + 1

	return repository.updateTicketForDate(ctx, date, newTicketNumber)
}

func (repository *OrderRespository) createNewTicketForDate(ctx context.Context, date int64) int {
	orderTicketNumber := entities.OrderTicketNumber{
		Date:         date,
		TicketNumber: 1,
	}

	errCreate := repository.db.WithContext(ctx).Create(&orderTicketNumber).Error

	if errCreate != nil {
		return 999
	}

	return 1
}

func (repository *OrderRespository) updateTicketForDate(ctx context.Context, date int64, newTicketNumber int) int {
	err := repository.db.WithContext(ctx).
		Model(&entities.OrderTicketNumber{}).
		Where("date = ?", date).
		Update("ticket_number", newTicketNumber).
		Error

	if err != nil {
		return 999
	}

	return newTicketNumber
}
