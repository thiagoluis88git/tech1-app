package model

type MercadoLivrePaymentResponse struct {
	ID                int64  `json:"id"`
	Status            string `json:"status"`
	ExternalReference string `json:"external_reference"`
	PreferenceID      string `json:"preference_id"`
	Marketplace       string `json:"marketplace"`
	NotificationURL   string `json:"notification_url"`
	DateCreated       string `json:"date_created"`
	LastUpdated       string `json:"last_updated"`
	OrderStatus       string `json:"order_status"`
	ClientID          string `json:"client_id"`
}
