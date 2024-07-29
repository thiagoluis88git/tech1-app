package dto

type ExternalPaymentEvent struct {
	Resource string `json:"resource"`
	Topic    string `json:"topic"`
}
