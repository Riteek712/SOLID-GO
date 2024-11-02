package main

import (
	"fmt"
)

// The Interface Segregation Principle (ISP), part of the SOLID principles, states that a client should not be forced to depend on interfaces it does not use. Essentially, instead of one large, general-purpose interface, it’s better to have multiple, smaller interfaces that are specific to client needs. This way, implementing structs don’t end up with methods they don’t need, leading to more modular and maintainable code.

// Let’s look at a CRUD example in Go. Imagine we have a system where some entities only need read capabilities, while others need full CRUD (Create, Read, Update, Delete) operations. We can define smaller interfaces tailored to each use case.

// Code Example: Interface Segregation Principle
// Define specific interfaces for each operation.
// Compose a complete CRUD interface from the smaller, focused interfaces.
// Implement only the required interfaces in each struct.

// Define smaller, segregated interfaces for each operation
type Reader interface {
	Read(id int) string
}

type Writer interface {
	Write(data string) int
}

type Updater interface {
	Update(id int, data string) bool
}

type Deleter interface {
	Delete(id int) bool
}

// CRUD interface composed of smaller interfaces
type CRUD interface {
	Reader
	Writer
	Updater
	Deleter
}

// LogStore only implements Read and Write
type LogStore struct{}

func (l LogStore) Read(id int) string {
	return fmt.Sprintf("Reading log entry %d", id)
}

func (l LogStore) Write(data string) int {
	fmt.Printf("Writing log entry: %s\n", data)
	return 1 // returning an example ID
}

// DatabaseStore implements full CRUD
type DatabaseStore struct{}

func (db DatabaseStore) Read(id int) string {
	return fmt.Sprintf("Reading database record %d", id)
}

func (db DatabaseStore) Write(data string) int {
	fmt.Printf("Writing database record: %s\n", data)
	return 1 // returning an example ID
}

func (db DatabaseStore) Update(id int, data string) bool {
	fmt.Printf("Updating database record %d with %s\n", id, data)
	return true
}

func (db DatabaseStore) Delete(id int) bool {
	fmt.Printf("Deleting database record %d\n", id)
	return true
}

// ReadData only requires the Reader interface
func ReadData(r Reader, id int) {
	fmt.Println(r.Read(id))
}

// WriteData only requires the Writer interface
func WriteData(w Writer, data string) {
	w.Write(data)
}

// ManageData requires the full CRUD interface
func ManageData(crud CRUD, id int, data string) {
	fmt.Println(crud.Read(id))
	crud.Write(data)
	crud.Update(id, data)
	crud.Delete(id)
}

func main() {
	logStore := LogStore{}
	dbStore := DatabaseStore{}

	// Using LogStore with only Read and Write
	ReadData(logStore, 1)      // Works because LogStore implements Reader
	WriteData(logStore, "log") // Works because LogStore implements Writer

	// Using DatabaseStore with full CRUD
	ManageData(dbStore, 1, "new data") // Works because DatabaseStore implements CRUD
}
