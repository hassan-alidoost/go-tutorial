package services

type NotificationService interface {
	Notify(customerID uint, msg string)
}
