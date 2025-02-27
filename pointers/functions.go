//	Create a pointer and simply pass it to the Function
//
// In the below program we are taking a function ptf which have integer type pointer
//
//	parameter which instructs the function to accept only the pointer type argument.
//
// Basically, this function changed the value of the variable x.
// At starting x contains the value 100. But after the function call,
// value changed to 748 as shown in the output.

package main

import "fmt"

func ptf(a *int) {

	// dereferencing
	*a = 748
}
func pointer_func() {
	var x = 100
	fmt.Printf("The value of x before function call is:%d\n", x)
	//taking a pointervariable
	//andassigningtheaddress
	//ofxtoit
	var pa *int = &x
	//callingthefunctionby
	//passingpointofunction
	ptf(pa)
	fmt.Printf("The value of x after function call is:%d\n", x)

}
