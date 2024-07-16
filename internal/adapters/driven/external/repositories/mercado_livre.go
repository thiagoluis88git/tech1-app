package repositories

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/thiagoluis88git/tech1/internal/adapters/driven/external/model"
	"github.com/thiagoluis88git/tech1/internal/adapters/driven/external/remote"
	"github.com/thiagoluis88git/tech1/internal/core/domain"
	"github.com/thiagoluis88git/tech1/internal/core/ports"
	"github.com/thiagoluis88git/tech1/pkg/environment"
)

type MercadoLivreRepositoryImpl struct {
	ds      remote.MercadoLivreDataSource
	webHook string
}

func NewMercadoLivreRepository(ds remote.MercadoLivreDataSource) ports.QRCodePaymentRepository {
	return &MercadoLivreRepositoryImpl{
		ds:      ds,
		webHook: environment.GetWebhookMercadoLivrePaymentURL(),
	}
}

func (repo *MercadoLivreRepositoryImpl) Generate(ctx context.Context, token string, form domain.Order, orderID int) (domain.QRCodeDataResponse, error) {
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
		return domain.QRCodeDataResponse{}, err
	}

	return domain.QRCodeDataResponse{
		Data: qrData,
	}, nil
}

func (repo *MercadoLivreRepositoryImpl) GetQRCodePaymentData(ctx context.Context, token string, endpoint string) (domain.MercadoLivrePayment, error) {
	response, err := repo.ds.GetPaymentData(ctx, token, endpoint)

	if err != nil {
		return domain.MercadoLivrePayment{}, err
	}

	mercadoLivrePayment := domain.MercadoLivrePayment(response)

	return mercadoLivrePayment, nil
}
