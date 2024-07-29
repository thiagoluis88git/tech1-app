package repositories

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/thiagoluis88git/tech1/internal/core/domain/dto"
	"github.com/thiagoluis88git/tech1/internal/core/domain/repository"
	"github.com/thiagoluis88git/tech1/internal/integrations/model"
	"github.com/thiagoluis88git/tech1/internal/integrations/remote"
	"github.com/thiagoluis88git/tech1/pkg/environment"
)

type MercadoLivreRepositoryImpl struct {
	ds      remote.MercadoLivreDataSource
	webHook string
}

func NewMercadoLivreRepository(ds remote.MercadoLivreDataSource) repository.QRCodePaymentRepository {
	return &MercadoLivreRepositoryImpl{
		ds:      ds,
		webHook: environment.GetWebhookMercadoLivrePaymentURL(),
	}
}

func (repo *MercadoLivreRepositoryImpl) Generate(ctx context.Context, token string, form dto.Order, orderID int) (dto.QRCodeDataResponse, error) {
	items := make([]model.Item, 0)

	totalAmount := 0

	for _, value := range form.OrderProduct {
		productId := strconv.Itoa(int(value.ProductID))
		totalAmount += int(value.ProductPrice)

		items = append(items, model.Item{
			Description: fmt.Sprintf("FastFood Pagamento - Produto: %v", productId),
			SkuNumber:   productId,
			Title:       fmt.Sprintf("FastFood Pagamento - Produto: %v", productId),
			UnitMeasure: "unit",
			Quantity:    1,
			UnitPrice:   int(value.ProductPrice),
			TotalAmount: int(value.ProductPrice),
		})
	}

	expirationDate := time.Now().Local().Add(time.Hour * 12)

	input := model.QRCodeInput{
		Description:       fmt.Sprintf("Order: %v", orderID),
		TotalAmount:       totalAmount,
		ExpirationDate:    expirationDate.Format("2006-01-02T15:04:05.999Z07:00"),
		ExternalReference: fmt.Sprintf("%v|%v", strconv.Itoa(orderID), strconv.Itoa(int(form.PaymentID))),
		Items:             items,
		Title:             fmt.Sprintf("FastFood Pagamento - Nr: %v", form.TicketNumber),
		NotificationUrl:   repo.webHook,
	}

	qrData, err := repo.ds.Generate(ctx, token, input)

	if err != nil {
		return dto.QRCodeDataResponse{}, err
	}

	return dto.QRCodeDataResponse{
		Data: qrData,
	}, nil
}

func (repo *MercadoLivreRepositoryImpl) GetQRCodePaymentData(ctx context.Context, token string, endpoint string) (dto.ExternalPaymentInformation, error) {
	response, err := repo.ds.GetPaymentData(ctx, token, endpoint)

	if err != nil {
		return dto.ExternalPaymentInformation{}, err
	}

	mercadoLivrePayment := dto.ExternalPaymentInformation(response)

	return mercadoLivrePayment, nil
}
