package services

import (
	"context"
	"net/http"
	"sync"

	"github.com/thiagoluis88git/tech1/internal/core/domain"
	"github.com/thiagoluis88git/tech1/internal/core/ports"
	"github.com/thiagoluis88git/tech1/pkg/responses"
)

type MercadoLivreService struct {
	repository      ports.MercadoLivreRepository
	orderRepository ports.OrderRepository
}

func NewMercadoLivreService(
	repository ports.MercadoLivreRepository,
	orderRepository ports.OrderRepository,
) *MercadoLivreService {
	return &MercadoLivreService{
		repository:      repository,
		orderRepository: orderRepository,
	}
}

func (service *MercadoLivreService) GenerateQRCode(
	ctx context.Context,
	token string,
	order domain.Order,
	date int64,
	wg *sync.WaitGroup,
	ch chan bool,
) (domain.QRCodeDataResponse, error) {
	//Block this code below until this Channel be empty (by reading with <-ch)
	ch <- true

	order.TicketNumber = service.orderRepository.GetNextTicketNumber(ctx, date)

	orderResponse, err := service.orderRepository.CreatePayingOrder(ctx, order)

	if err != nil {
		return domain.QRCodeDataResponse{}, responses.GetResponseError(err, "QRCodeGeneratorService")
	}

	qrCode, err := service.repository.Generate(ctx, token, order, int(orderResponse.OrderId))

	if err != nil {
		errDelete := service.orderRepository.DeleteOrder(ctx, orderResponse.OrderId)

		if errDelete != nil {
			return domain.QRCodeDataResponse{}, responses.GetResponseError(errDelete, "QRCodeGeneratorService")
		}

		return domain.QRCodeDataResponse{}, responses.GetResponseError(err, "QRCodeGeneratorService")
	}

	// Release the channel to others process be able to start a new order creation
	<-ch
	wg.Done()

	return qrCode, nil
}

func (service *MercadoLivreService) FinishOrder(ctx context.Context, token string, form domain.MercadoLivrePaymentForm) error {
	if form.Topic != "payment" {
		return &responses.NetworkError{
			Code: http.StatusNotAcceptable,
		}
	}

	err := service.repository.GetMercadoLivrePaymentData(ctx, token, form.Resource)

	if err != nil {
		return err
	}

	return nil
}
