package model

import (
	"bytes"
	"encoding/json"
)

type QRCodeData struct {
	QRData string `json:"qr_data"`
}

type QRCodeInput struct {
	ExpirationDate    string `json:"expiration_date"`
	ExternalReference string `json:"external_reference"`
	Description       string `json:"description"`
	Title             string `json:"title"`
	NotificationUrl   string `json:"notification_url"`
	Items             []Item `json:"items"`
	TotalAmount       int    `json:"total_amount"`
}

type Item struct {
	Description string `json:"description"`
	SkuNumber   string `json:"sku_number"`
	Title       string `json:"title"`
	UnitMeasure string `json:"unit_measure"`
	Quantity    int    `json:"quantity"`
	UnitPrice   int    `json:"unit_price"`
	TotalAmount int    `json:"total_amount"`
}

func (input *QRCodeInput) GetJSONBody() (*bytes.Buffer, error) {
	jsonValue, err := json.Marshal(input)

	if err != nil {
		return nil, err
	}

	return bytes.NewBuffer(jsonValue), nil
}
