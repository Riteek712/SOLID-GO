package main

import (
	"fmt"
)

// To demonstrate the Open-Closed Principle (OCP) in a practical Go example, let’s consider a scenario where we need to calculate discounts for different types of customers. The Open-Closed Principle states that software entities (classes, modules, functions) should be open for extension but closed for modification.

// In our example, we will create a base interface for discounts and specific implementations for different types of customers. This way, we can add new discount types without changing the existing code, adhering to the OCP.

// Code Example: Open-Closed Principle
// Define an interface for Discounts.
// Implement concrete discount types for various customer categories.
// Add a function to calculate discounts without modifying it for new customer types.

// DiscountCalculator defines the behavior for calculating discounts
type DiscountCalculator interface {
	Calculate(price float64) float64
}

// RegularCustomerDiscount applies no discount for regular customers
type RegularCustomerDiscount struct{}

func (r RegularCustomerDiscount) Calculate(price float64) float64 {
	return price // No discount applied
}

// LoyalCustomerDiscount applies a 10% discount for loyal customers
type LoyalCustomerDiscount struct{}

func (l LoyalCustomerDiscount) Calculate(price float64) float64 {
	return price * 0.90 // 10% discount
}

// VIPCustomerDiscount applies a 20% discount for VIP customers
type VIPCustomerDiscount struct{}

func (v VIPCustomerDiscount) Calculate(price float64) float64 {
	return price * 0.80 // 20% discount
}

// NewCustomerDiscount applies a 5% discount for new customers
type NewCustomerDiscount struct{}

func (n NewCustomerDiscount) Calculate(price float64) float64 {
	return price * 0.95 // 5% discount
}

// CalculateFinalPrice calculates the final price after applying the discount
func CalculateFinalPrice(price float64, discountCalculator DiscountCalculator) float64 {
	return discountCalculator.Calculate(price)
}

func main() {
	// Test data
	price := 100.0

	// Calculate prices for different customer types
	regularDiscount := RegularCustomerDiscount{}
	loyalDiscount := LoyalCustomerDiscount{}
	vipDiscount := VIPCustomerDiscount{}
	newCustomerDiscount := NewCustomerDiscount{}

	fmt.Printf("Regular customer price: $%.2f\n", CalculateFinalPrice(price, regularDiscount))
	fmt.Printf("Loyal customer price: $%.2f\n", CalculateFinalPrice(price, loyalDiscount))
	fmt.Printf("VIP customer price: $%.2f\n", CalculateFinalPrice(price, vipDiscount))
	fmt.Printf("New customer price: $%.2f\n", CalculateFinalPrice(price, newCustomerDiscount))
}
