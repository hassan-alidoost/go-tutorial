package main

import (
	"errors"
	"fmt"
	"time"
)

type Timestamped struct {
	CreatedAt time.Time
	SentAt    time.Time
}

func (t *Timestamped) SetSentAt() {
	t.SentAt = time.Now()
}

type Retryable struct {
	Attempts    uint
	MaxAttempts uint
}

func (r *Retryable) IncrementAttempt() {
	r.Attempts++
}

type PriorityType uint

const (
	High PriorityType = iota
	Medium
	Low
)

type Prioritized struct {
	Priority PriorityType
}

type NotificationSender interface {
	Send(message string, recipient string) error
	GetType() string
	GetMaxAttempt() uint
	GetAttempt() uint
	IncrementAttempt()
}

type EmailNotification struct {
	Timestamped
	Retryable
	Prioritized
	Email string
}

func (e *EmailNotification) Send(message string, recipient string) error {
	if e.GetAttempt() < 3 {
		return errors.New("failed to send email")
	}
	fmt.Printf("Sending email to %v, email: %v ...\n", recipient, e.Email)
	fmt.Printf("Message: %v\n", message)
	e.SetSentAt()
	e.IncrementAttempt()
	return nil
}

func (e *EmailNotification) GetType() string {
	return "EmailNotification"
}
func (e *EmailNotification) GetMaxAttempt() uint {
	return e.MaxAttempts
}
func (e *EmailNotification) GetAttempt() uint {
	return e.Attempts
}
func (e *EmailNotification) IncrementAttempt() {
	e.Attempts++
}

type SMSNotification struct {
	Timestamped
	Retryable
	Prioritized
	PhoneNumber string
}

func (s *SMSNotification) Send(message string, recipient string) error {
	fmt.Printf("Sending sms to %v ...\n, PhoneNumber: %v", recipient, s.PhoneNumber)
	fmt.Printf("Message: %v\n", message)
	s.SetSentAt()
	s.IncrementAttempt()
	return nil
}

func (s *SMSNotification) GetType() string {
	return "SMSNotification"
}
func (s *SMSNotification) GetMaxAttempt() uint {
	return s.MaxAttempts
}
func (s *SMSNotification) GetAttempt() uint {
	return s.MaxAttempts
}
func (s *SMSNotification) IncrementAttempt() {
	s.Attempts++
}

type PushNotification struct {
	Timestamped
	Retryable
	Prioritized
	DeviceToken string
}

func (p *PushNotification) Send(message string, recipient string) error {
	fmt.Printf("Sending push notification to %v, DeviceToken: %v ...\n", recipient, p.DeviceToken)
	fmt.Printf("Message: %v\n", message)
	p.SetSentAt()
	p.IncrementAttempt()
	return nil
}

func (p *PushNotification) GetType() string {
	return "PushNotification"
}
func (p *PushNotification) GetMaxAttempt() uint {
	return p.MaxAttempts
}
func (p *PushNotification) GetAttempt() uint {
	return p.MaxAttempts
}
func (p *PushNotification) IncrementAttempt() {
	p.Attempts++
}

type NotificationService struct {
	sender NotificationSender
}

func NewNotificationService(sender NotificationSender) NotificationService {
	return NotificationService{sender: sender}
}

func (ns *NotificationService) SendWithRetry(message string, recipient string) error {
	for {
		err := ns.sender.Send(message, recipient)

		if err == nil {
			fmt.Println("Notification sent successfully!")
			return nil
		}

		ns.sender.IncrementAttempt()
		fmt.Println("Retry attempt: ", ns.sender.GetAttempt())

		if ns.sender.GetAttempt() >= ns.sender.GetMaxAttempt() {
			return errors.New("max attempt reached")
		}

		time.Sleep(time.Duration(ns.sender.GetAttempt()) * time.Second)
	}
}

func main() {

	email := &EmailNotification{
		Email: "user@example.com",
		Timestamped: Timestamped{
			CreatedAt: time.Now(),
		},
		Retryable: Retryable{
			MaxAttempts: 4,
		},
		Prioritized: Prioritized{
			Priority: High,
		},
	}

	service := NewNotificationService(email)

	err := service.SendWithRetry(
		"Welcome to our platform!",
		email.Email,
	)

	if err != nil {
		fmt.Println("Error:", err)
	}
}
