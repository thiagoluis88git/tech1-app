package domain

type QRCodeForm struct {
	Description       string     `json:"description"`
	ExpirationDate    string     `json:"expirationDate"`
	ExternalReference string     `json:"externalReference"`
	Items             []ItemForm `json:"items"`
	NotificationURL   string     `json:"notificationUrl"`
	Title             string     `json:"title"`
	TotalAmount       int        `json:"totalAmount"`
}

type ItemForm struct {
	SkuNumber   string `json:"skuNumber"`
	Category    string `json:"category"`
	Title       string `json:"title"`
	Description string `json:"description"`
	UnitPrice   int    `json:"unitPrice"`
	Quantity    int    `json:"quantity"`
	UnitMeasure string `json:"unitMeasure"`
	TotalAmount int    `json:"totalAmount"`
}

type QRCodeDataResponse struct {
	Data string `json:"data"`
}
