package repositories

import (
	"context"
	"thiagoluis88git/tech1/internal/adapters/driven/entities"
	"thiagoluis88git/tech1/internal/core/domain"
	"thiagoluis88git/tech1/internal/core/ports"
	"thiagoluis88git/tech1/pkg/responses"

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
		OrderStatus: entities.OrderStatusCreated,
		TotalPrice:  order.TotalPrice,
		CustomerID:  order.CustomerID,
		TickerID:    order.TickerId,
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

	var customerName *string
	if orderEntity.Customer != nil {
		customerName = &orderEntity.Customer.Name
	}

	return domain.OrderResponse{
		OrderDate:    orderEntity.CreatedAt,
		TickerId:     orderEntity.TickerID,
		CustomerName: customerName,
	}, nil
}
