package model

import (
	"bytes"
	"encoding/json"
)

type QRCodeData struct {
	QRData string `json:"qr_data"`
}

type QRCodeInput struct {
	Description       string `json:"description"`
	ExpirationDate    string `json:"expiration_date"`
	ExternalReference string `json:"external_reference"`
	Items             []Item `json:"items"`
	NotificationURL   string `json:"notification_url"`
	Title             string `json:"title"`
	TotalAmount       int    `json:"total_amount"`
}

type Item struct {
	SkuNumber   string `json:"sku_number"`
	Category    string `json:"category"`
	Title       string `json:"title"`
	Description string `json:"description"`
	UnitPrice   int    `json:"unit_price"`
	Quantity    int    `json:"quantity"`
	UnitMeasure string `json:"unit_measure"`
	TotalAmount int    `json:"total_amount"`
}

func (input *QRCodeInput) GetJSONBody() (*bytes.Buffer, error) {
	jsonValue, err := json.Marshal(input)

	if err != nil {
		return nil, err
	}

	return bytes.NewBuffer(jsonValue), nil
}
