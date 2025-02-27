package main

//Pointers in Go programming language or Golang is a variable that is used
//  to store the memory address of another variable. Pointers in Golang is
// also termed as the special variables. The variables are used to store
// some data at a particular memory address in the system. The memory
// address is always found in hexadecimal format(starting with 0x like 0xFFAAF etc.).
//what is the need for the pointers?
//to remember the memorylocation of variable is overhead,so we use pointers,now varabiles accessed with their memeorylocation
//Declaration and initialization
// * Operator also termed as the dereferencing operator used to declare pointer variable and access the value stored in the address.
// & operator termed as address operator used to returns the address of a variable or to access the address of a variable to a pointer.
//Declaration a pointer:-var pointer_name *Data_Type
//var s *string
//initilizationofpointer
//var a = 45
//initializationofpointerswithmenoryaddressofvariable
//var s *int = &a
//importantpoints
//1.The default value or zero-value of a pointer is always nil. Or you can say that an uninitialized pointer will always have a nil value.
//2.Declaration and initialization of the pointers can be done into a single line
//3. If you are specifying the data type along with the pointer declaration then the pointer will be able to handle the memory address of
// that specified data type variable. For example, if you taking a pointer
//  of string type then the address of the variable that you will give to
//  a pointer will be only of string data type variable, not any other type.
//4. To overcome the above mention problem you can use the Type Inference
// concept of the var keyword. There is no need to specify the data type
// during the declaration. The type of a pointer variable can also be
// determined by the compiler like a normal variable. Here we will not use
// the * operator. It will internally determine by the compiler as we are
// initializing the variable with the address of another variable.
func main() {

	// storing the hexadecimal
	// values in variables
	// x := 0xFF
	// y := 0x9C

	// // Displaying the values
	// fmt.Printf("Type of variable x is %T\n", x)
	// fmt.Printf("Value of x in hexadecimal is %X\n", x)
	// fmt.Printf("Value of x in decimal is %v\n", x)

	// fmt.Printf("Type of variable y is %T\n", y)
	// fmt.Printf("Value of y in hexadecimal is %X\n", y)
	// fmt.Printf("Value of y in decimal is %v\n", y)
	// //takinganormalvariable

	// var px int = 5748
	// //declarationofpointer
	// var p *int
	// //initilazationofpointer
	// p = &px
	// // displaying the result
	// fmt.Println("Value stored in x = ", px)
	// fmt.Println("Address of x = ", &px)
	// fmt.Println("Value stored in variable p = ", p)

	// // using var keyword
	// // we are not defining
	// // any type with variable
	// var py = 458

	// // taking a pointer variable using
	// // var keyword without specifying
	// // the type
	// var pp = &py

	// fmt.Println("Value stored in y = ", py)
	// fmt.Println("Address of y = ", &py)
	// fmt.Println("Value stored in pointer variable p = ", pp)

	// // using := operator to declare
	// // and initialize the variable
	// pz := 458

	// // taking a pointer variable using
	// // := by assigning it with the
	// // address of variable y
	// ppp := &pz

	// fmt.Println("Value stored in y = ", pz)
	// fmt.Println("Address of y = ", &pz)
	// fmt.Println("Value stored in pointer variable p = ", ppp)

	// // this is dereferencing a pointer
	// // using * operator before a pointer
	// // variable to access the value stored
	// // at the variable at which it is pointing

	// fmt.Println("Value stored in y(*p) = ", *p)
	// // changing the value of y by assigning
	// // the new value to the pointer
	// *p = 500

	// fmt.Println("Value stored in y(*p) after Changing = ", y)
	//passing pointer in function
	//pointer_func()
	struct_pointer()
}
