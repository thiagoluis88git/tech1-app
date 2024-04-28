package repositories

import (
	"context"
	"thiagoluis88git/tech1/internal/adapters/driven/entities"
	"thiagoluis88git/tech1/internal/core/domain"
	"thiagoluis88git/tech1/internal/core/ports"
	"thiagoluis88git/tech1/pkg/responses"
	"time"

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
		PaymentID:   order.PaymentID,
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

	err = tx.Commit().Error

	if err != nil {
		tx.Rollback()
		return domain.OrderResponse{}, responses.GetDatabaseError(err)
	}

	return domain.OrderResponse{
		OrderId:   orderEntity.ID,
		OrderDate: orderEntity.CreatedAt,
		TickerId:  orderEntity.TickerID,
	}, nil
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
