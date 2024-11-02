package main

import "fmt"

// This principle states that high-level modules should not depend on low-level modules, but rather both should depend on abstractions. This helps to reduce the coupling between components and make the code more flexible and maintainable.

// Notifier interface represents the abstraction for sending notifications
type Notifier interface {
	Send(message string) error
}

// EmailNotifier is a low-level module that implements the Notifier interface
type EmailNotifier struct{}

func (e EmailNotifier) Send(message string) error {
	fmt.Println("Sending email with message:", message)
	return nil
}

// SMSNotifier is another low-level module that implements the Notifier interface
type SMSNotifier struct{}

func (s SMSNotifier) Send(message string) error {
	fmt.Println("Sending SMS with message:", message)
	return nil
}

// NotificationService is a high-level module that depends on the Notifier abstraction
type NotificationService struct {
	notifier Notifier
}

// NewNotificationService creates a new NotificationService with the specified notifier
func NewNotificationService(notifier Notifier) *NotificationService {
	return &NotificationService{notifier: notifier}
}

// Notify sends a notification using the provided Notifier
func (n *NotificationService) Notify(message string) {
	err := n.notifier.Send(message)
	if err != nil {
		fmt.Println("Failed to send notification:", err)
	} else {
		fmt.Println("Notification sent successfully!")
	}
}

func main() {
	// High-level module depends on abstractions (Notifier), not concrete implementations (EmailNotifier, SMSNotifier)
	emailNotifier := EmailNotifier{}
	smsNotifier := SMSNotifier{}

	// Create NotificationService with EmailNotifier
	notificationServiceEmail := NewNotificationService(emailNotifier)
	notificationServiceEmail.Notify("Hello via Email!")

	// Switch to SMSNotifier by simply changing the dependency injection
	notificationServiceSMS := NewNotificationService(smsNotifier)
	notificationServiceSMS.Notify("Hello via SMS!")
}
