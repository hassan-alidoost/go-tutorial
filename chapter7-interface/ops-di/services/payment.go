package services

type PaymentGateway interface {
	Authorize(amount float64) (string, error)
	Refund(amount float64) error
}
