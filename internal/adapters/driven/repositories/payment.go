package repositories

import (
	"context"

	"github.com/thiagoluis88git/tech1/internal/adapters/driven/model"
	"github.com/thiagoluis88git/tech1/internal/core/domain"
	"github.com/thiagoluis88git/tech1/internal/core/ports"
	"github.com/thiagoluis88git/tech1/pkg/responses"

	"gorm.io/gorm"
)

type PaymentRepository struct {
	db *gorm.DB
}

func NewPaymentRepository(db *gorm.DB) ports.PaymentRepository {
	return &PaymentRepository{
		db: db,
	}
}

func (repository *PaymentRepository) GetPaymentTypes() []string {
	return []string{
		model.PaymentQRCodeType,
		model.PaymentCreditType,
	}
}

func (repository *PaymentRepository) CreatePaymentOrder(ctx context.Context, payment domain.Payment) (domain.PaymentResponse, error) {
	tx := repository.db.WithContext(ctx).Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return domain.PaymentResponse{}, responses.GetDatabaseError(err)
	}

	paymentEntity := model.Payment{
		CustomerID:    payment.CustomerID,
		TotalPrice:    payment.TotalPrice,
		PaymentType:   payment.PaymentType,
		PaymentStatus: model.PaymentPayingStatus,
	}

	err := tx.Create(&paymentEntity).Error

	if err != nil {
		tx.Rollback()
		return domain.PaymentResponse{}, responses.GetDatabaseError(err)
	}

	err = tx.Commit().Error

	if err != nil {
		tx.Rollback()
		return domain.PaymentResponse{}, responses.GetDatabaseError(err)
	}

	return domain.PaymentResponse{
		PaymentId: paymentEntity.ID,
	}, nil
}

func (repository *PaymentRepository) FinishPaymentWithError(ctx context.Context, paymentId uint) error {
	err := repository.db.WithContext(ctx).Model(&model.Payment{}).Where("id = ?", paymentId).Update("payment_status", model.PaymentPayedStatus).Error

	if err != nil {
		return responses.GetDatabaseError(err)
	}

	return nil
}

func (repository *PaymentRepository) FinishPaymentWithSuccess(ctx context.Context, paymentId uint) error {
	err := repository.db.WithContext(ctx).Model(&model.Payment{}).Where("id = ?", paymentId).Update("payment_status", model.PaymentErrorStatus).Error

	if err != nil {
		return responses.GetDatabaseError(err)
	}

	return nil
}
