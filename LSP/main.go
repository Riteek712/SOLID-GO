package main

import (
	"fmt"
)

// The Liskov Substitution Principle (LSP) is the "L" in SOLID principles and states that objects of a superclass should be replaceable with objects of a subclass without altering the correctness of the program. This principle essentially ensures that derived classes (subtypes) should behave consistently with the expectations set by their base classes (supertypes).

// Let’s demonstrate this in a Go example with a simple animal classification system. We’ll define a base interface for animals that can speak and then implement this interface in different animal types. Each animal class will provide its specific behavior while still honoring the contract of the base interface, allowing us to substitute any animal class where the base type is expected.

// Code Example: Liskov Substitution Principle
// Define a Speaker interface.
// Implement multiple types (animals) that adhere to the Speaker interface.
// Use a function that can accept any Speaker type and call the Speak method.

// Speaker defines the behavior for animals that can speak
type Speaker interface {
	Speak() string
}

// Dog struct implements Speaker interface
type Dog struct{}

func (d Dog) Speak() string {
	return "Woof!"
}

// Cat struct implements Speaker interface
type Cat struct{}

func (c Cat) Speak() string {
	return "Meow!"
}

// Parrot struct implements Speaker interface
type Parrot struct{}

func (p Parrot) Speak() string {
	return "Squawk!"
}

// DescribeAnimal takes a Speaker and calls its Speak method
func DescribeAnimal(speaker Speaker) {
	fmt.Println("The animal says:", speaker.Speak())
}

func main() {
	// Each of these can be passed to DescribeAnimal, honoring the LSP
	dog := Dog{}
	cat := Cat{}
	parrot := Parrot{}

	DescribeAnimal(dog)    // Outputs: The animal says: Woof!
	DescribeAnimal(cat)    // Outputs: The animal says: Meow!
	DescribeAnimal(parrot) // Outputs: The animal says: Squawk!
}
