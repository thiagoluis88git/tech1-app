package services

import (
	"context"
	"net/http"
	"strconv"
	"strings"
	"sync"

	"github.com/thiagoluis88git/tech1/internal/core/domain"
	"github.com/thiagoluis88git/tech1/internal/core/ports"
	"github.com/thiagoluis88git/tech1/pkg/responses"
)

type MercadoLivreService struct {
	repository        ports.MercadoLivreRepository
	orderRepository   ports.OrderRepository
	paymentRepository ports.PaymentRepository
}

func NewMercadoLivreService(
	repository ports.MercadoLivreRepository,
	orderRepository ports.OrderRepository,
	paymentRepository ports.PaymentRepository,
) *MercadoLivreService {
	return &MercadoLivreService{
		repository:        repository,
		orderRepository:   orderRepository,
		paymentRepository: paymentRepository,
	}
}

func (service *MercadoLivreService) GenerateQRCode(
	ctx context.Context,
	token string,
	qrOrder domain.QRCodeOrder,
	date int64,
	wg *sync.WaitGroup,
	ch chan bool,
) (domain.QRCodeDataResponse, error) {
	//Block this code below until this Channel be empty (by reading with <-ch)
	ch <- true

	qrOrder.TicketNumber = service.orderRepository.GetNextTicketNumber(ctx, date)

	payment := domain.Payment{
		TotalPrice:  qrOrder.TotalPrice,
		PaymentType: "QR Code",
	}

	paymentResponse, err := service.paymentRepository.CreatePaymentOrder(ctx, payment)

	if err != nil {
		return domain.QRCodeDataResponse{}, responses.GetResponseError(err, "QRCodeGeneratorService")
	}

	qrOrder.PaymentID = paymentResponse.PaymentId

	order := domain.Order{
		TotalPrice:   qrOrder.TotalPrice,
		CustomerID:   qrOrder.CustomerID,
		OrderProduct: []domain.OrderProduct(qrOrder.OrderProduct),
		TicketNumber: qrOrder.TicketNumber,
		PaymentID:    qrOrder.PaymentID,
	}

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
	if form.Topic != "merchant_order" {
		return &responses.NetworkError{
			Code: http.StatusNotAcceptable,
		}
	}

	mercadoLivrePayment, err := service.repository.GetMercadoLivrePaymentData(ctx, token, form.Resource)

	if err != nil {
		return err
	}

	if mercadoLivrePayment.OrderStatus == "paid" {
		ids := strings.Split(mercadoLivrePayment.ExternalReference, "|")

		orderID, _ := strconv.Atoi(ids[0])
		paymentID, _ := strconv.Atoi(ids[1])

		service.paymentRepository.FinishPaymentWithSuccess(ctx, uint(paymentID))
		service.orderRepository.FinishOrderWithPayment(ctx, uint(orderID), uint(paymentID))
	}

	return nil
}
