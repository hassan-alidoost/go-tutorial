package services

import (
	"github.com/hassan-alidoost/go-tutorial/chapter7-interface/ops-di/domains"
	"github.com/hassan-alidoost/go-tutorial/chapter7-interface/ops-di/repositories"
)

type OrderService struct {
	productRepo repositories.ProductRepository
	orderRepo   repositories.OrderRepository
	payment     PaymentGateway
	inventory   InventoryService
	notifier    NotificationService
	logger      Logger
}

func NewOrderService(
	productRepo repositories.ProductRepository,
	orderRepo repositories.OrderRepository,
	payment PaymentGateway,
	inventory InventoryService,
	notifier NotificationService,
	logger Logger,
) *OrderService {
	return &OrderService{
		productRepo: productRepo,
		orderRepo:   orderRepo,
		payment:     payment,
		inventory:   inventory,
		notifier:    notifier,
		logger:      logger,
	}
}

func (s *OrderService) CreateOrder(customerID uint, items []domains.OrderItem) (*domains.Order, error) {
	s.logger.Info("Creating order for customer ", customerID)

	order := &domains.Order{
		ID:         1,
		CustomerID: customerID,
		Items:      items,
		Status:     domains.OrderPending,
	}
	order.CalTotal()

	for _, item := range items {
		if err := s.inventory.Reserve(item.ProductID, item.Quantity); err != nil {
			s.logger.Error("Failed to reserve the product", err.Error())
			return nil, err
		}
	}

	trxID, err := s.payment.Authorize(order.Total)
	if err != nil {
		s.logger.Error("Failed to authorize payment", err.Error())
		for _, item := range items {
			_ = s.inventory.Release(item.ProductID, item.Quantity)
		}
		return nil, err
	}
	s.logger.Info("Authorized payment successfully", trxID)

	if err := s.orderRepo.Save(*order); err != nil {
		s.logger.Error("Failed to save order", err.Error())
		_ = s.payment.Refund(order.Total)
		for _, item := range items {
			_ = s.inventory.Release(item.ProductID, item.Quantity)
		}
		return nil, err
	}

	s.notifier.Notify(customerID, "Your order is confirmed")
	err = s.orderRepo.UpdateStatus(order.ID, domains.OrderConfirmed)
	if err != nil {
		s.logger.Error("Failed to update order status to confirmed", err.Error())
	}

	return order, nil
}

func (s *OrderService) CancelOrder(orderID domains.OrderID) error {
	s.logger.Info("Canceling order ", orderID)

	order, err := s.orderRepo.Find(orderID)
	if err != nil {
		s.logger.Error("Failed to find order", err.Error())
		return err
	}

	err = s.orderRepo.UpdateStatus(order.ID, domains.OrderCancelled)
	if err != nil {
		s.logger.Error("Failed to update order status", err.Error())
		return err
	}

	for _, item := range order.Items {
		err := s.inventory.Release(item.ProductID, item.Quantity)
		if err != nil {
			s.logger.Error("Failed to release order items", err.Error())
			return err
		}
	}

	err = s.payment.Refund(order.Total)
	if err != nil {
		s.logger.Error("Failed to refund amount to customer", err.Error())
		return err
	}

	return nil
}

func (s *OrderService) GetOrder(orderID domains.OrderID) (*domains.Order, error) {
	order, err := s.orderRepo.Find(orderID)
	if err != nil {
		s.logger.Error("Failed to find order", err.Error())
		return nil, err
	}
	return order, nil
}
