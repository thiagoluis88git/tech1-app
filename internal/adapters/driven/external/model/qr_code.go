package model

import (
	"bytes"
	"encoding/json"
	"time"
)

type QRCodeData struct {
	QRData string `json:"qr_data"`
}

type QRCodeInput struct {
	ExpirationDate    time.Time `json:"expiration_date"`
	ExternalReference string    `json:"external_reference"`
	Items             []Item    `json:"items"`
	TotalAmount       int       `json:"total_amount"`
}

type Item struct {
	SkuNumber string `json:"sku_number"`
}

func (input *QRCodeInput) GetJSONBody() (*bytes.Buffer, error) {
	jsonValue, err := json.Marshal(input)

	if err != nil {
		return nil, err
	}

	return bytes.NewBuffer(jsonValue), nil
}
