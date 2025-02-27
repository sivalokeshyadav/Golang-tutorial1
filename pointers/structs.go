// In Go, structs are used to create custom data types that group different
//  fields together. When working with structs, using pointers can be especially
//  beneficial for managing memory efficiently and for avoiding unnecessary
//  copying of data. A pointer to a struct allows you to directly reference
//  and modify the data in the original struct without making a copy.

// Why Use Pointers with Structs?
// Using pointers with structs is helpful when:

// You want to avoid copying large structs by passing references instead.
// You need to modify the original struct values from within a function.
// You want to improve performance by reducing memory usage.
// Declaring a Pointer to a Struct
// There are two common ways to create a pointer to a struct in Go:

// Using the & operator to get the memory address of an existing struct.
// Using the new function, which returns a pointer to a newly allocated struct.
// Using new(Person) creates a pointer to an empty Person struct (personPointer).
// We then set the fields of this struct through the pointer.

// Accessing Struct Fields Through a Pointer
// In Go, we don’t need to explicitly dereference the pointer (using *)
// to access the fields. Go automatically dereferences pointers when
// accessing fields, so personPointer.name works directly
// without (*personPointer).name
package main

import "fmt"

type Person struct {
	name string
	age  int
}

func struct_pointer() {
	//creatinganinstanceofthestruct
	p1 := Person{name: "A", age: 25}

	// Creating a pointer to the struct
	personPointer := &p1

	// Accessing fields using the pointer
	fmt.Println("Name:", personPointer.name) // Automatically dereferences
	fmt.Println("Age:", personPointer.age)

	// Modifying struct values using the pointer
	personPointer.age = 26
	fmt.Println("Updated struct:", p1)
	// Creating a pointer to a new instance of Person
	personPointer1 := new(Person)
	personPointer1.name = "B"
	personPointer1.age = 30

	fmt.Println("Struct created with new:", *personPointer1)
}
