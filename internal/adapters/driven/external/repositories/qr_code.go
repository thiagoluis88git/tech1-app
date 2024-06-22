package repositories

import (
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

func (repo *QRCodeGeneratorRepositoryImpl) Generate(token string, form domain.QRCodeForm) (domain.QRCodeDataResponse, error) {
	items := make([]model.Item, 0)

	for _, value := range form.Items {
		items = append(items, model.Item(value))
	}

	input := model.QRCodeInput{
		Description:       form.Description,
		ExpirationDate:    form.ExpirationDate,
		ExternalReference: form.ExternalReference,
		NotificationURL:   form.NotificationURL,
		Title:             form.Title,
		TotalAmount:       form.TotalAmount,
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
