package repositories

import (
	"strconv"
	"time"

	"github.com/thiagoluis88git/tech1/internal/adapters/driven/external/model"
	"github.com/thiagoluis88git/tech1/internal/adapters/driven/external/remote"
	"github.com/thiagoluis88git/tech1/internal/core/domain"
	"github.com/thiagoluis88git/tech1/internal/core/ports"
)

type QRCodeGeneratorRepositoryImpl struct {
	ds remote.QRCodeGeneratorDataSource
}

func NewQRCodeGeneratorRepository(ds remote.QRCodeGeneratorDataSource) ports.QRCodeGeneratorRepository {
	return &QRCodeGeneratorRepositoryImpl{
		ds: ds,
	}
}

func (repo *QRCodeGeneratorRepositoryImpl) Generate(token string, form domain.Order, orderID int) (domain.QRCodeDataResponse, error) {
	items := make([]model.Item, 0)

	for _, value := range form.OrderProduct {
		items = append(items, model.Item{
			SkuNumber: strconv.Itoa(int(value.ProductID)),
		})
	}

	expirationDate := time.Now().Local().Add(time.Hour + time.Duration(10))

	input := model.QRCodeInput{
		TotalAmount:       form.TotalPrice,
		ExpirationDate:    expirationDate,
		ExternalReference: strconv.Itoa(orderID),
		Items:             items,
	}

	qrData, err := repo.ds.Generate(token, input)

	if err != nil {
		return domain.QRCodeDataResponse{}, err
	}

	return domain.QRCodeDataResponse{
		Data: qrData,
	}, nil
}
