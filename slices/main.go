// Slices in Go are a flexible and efficient way to represent arrays, and they are often used in place of arrays
// because of their dynamic size and added features.A slice is a reference to a portion of an array.
// It’s a data structure that describes a portion of an array by specifying the starting index and the length of the portion.
// Slice is a variable-length sequence that stores elements of a similar type, you are not allowed
// to store different type of elements in the same slice.
// It is allowed to store duplicate elements in the slice.
// The first index position in a slice is always 0 and the last one will be (length of slice – 1).
// Slices are dynamic, which means that their size can change as you add or
// remove elements. Go provides several built-in functions that allow you
// to modify slices, such as append, copy, and delete.
// a slice contain 3 components:-
// Pointer: The pointer is used to points to the first element of the array that is accessible
//
//	through the slice. Here, it is not necessary that the pointed element is the first element of the array.
//
// Length: The length is the total number of elements present in the array.
// Capacity: The capacity represents the maximum size upto which it can expand.
// In Go language, a slice can be created and initialized using the following ways:
// using slice literal:-The creation of slice literal is just like an array literal,
// but with one difference you are not allowed to specify the size of the slice in the square braces[]
// the right-hand side of this expression is the slice literal.
// var my_slice_1 = []string{"Geeks", "for", "Geeks"}
// using an array:-For creating a slice from the given array first you need
// to specify the lower and upper bound, which means slice can take elements
// from the array starting from the lower bound to the upper bound. It does not
// include the elements above from the upper bound.
// syntax:-array_name[low:high]
// Note: The default value of the lower bound is 0 and the default value of the
// upper bound is the total number of the elements present in the given array.
// using already existing slice:-
// For creating a slice from the given slice first you need to specify the
// lower and upper bound, which means slice can take elements from the given
// slice starting from the lower bound to the upper bound
// syntax:slice_name[low:high]
// usingmake():-This function takes three parameters, i.e, type, length, and capacity.

//		Here, capacity value is optional. It assigns an underlying array with a
//	 size that is equal to the given capacity and returns a slice which
//
// refers to the underlying array. Generally, make() function is used
//
//	to create an empty slice
//
// itrratiomoverslice
// important points in slice
// Zero value slice: In Go language, you are allowed to create a nil slice
// that does not contain any element in it. So the capacity and the length of this slice is 0.
package main

import "fmt"

func main() {
	array := [5]int{1, 2, 3, 4, 5}
	slice := array[1:4]
	fmt.Println("Array: ", array)
	fmt.Println("Slice: ", slice)
	appendSlice := append(slice, 6, 7, 8)
	fmt.Println("Append Slice:", appendSlice)
	arr1 := [7]string{"This", "is", "the", "tutorial",
		"of", "Go", "language"}
	fmt.Println("Array:", arr1)
	// Creating a slice
	myslice := arr1[1:6]

	// Display slice
	fmt.Println("Slice:", myslice)

	// Display length of the slice
	fmt.Printf("Length of the slice: %d", len(myslice))

	// Display the capacity of the slice
	fmt.Printf("\nCapacity of the slice: %d", cap(myslice))
	//slice using string literal
	var my_slice_1 = []string{"Geeks", "for", "Geeks"}
	fmt.Println("--------------------------------")
	fmt.Println("\nMy Slice 1:", my_slice_1)

	// Creating a slice
	//using shorthand declaration
	fmt.Println("--------------------------------")
	my_slice_2 := []int{12, 45, 67, 56, 43, 34, 45}
	fmt.Println("My Slice 2:", my_slice_2)
	//creating slice from array
	arr32 := [4]string{"Geeks", "for", "Geeks", "GFG"}

	var my_slice_111 = arr32[1:2]
	my_slice_211 := arr32[0:]
	my_slice_311 := arr32[:2]
	my_slice_411 := arr32[:]
	fmt.Println("--------------------------------")
	// Display the result
	fmt.Println("My Array: ", arr32)
	fmt.Println("My Slice 111: ", my_slice_111)
	fmt.Println("My Slice 211: ", my_slice_211)
	fmt.Println("My Slice 311: ", my_slice_311)
	fmt.Println("My Slice 411: ", my_slice_411)

	//creating slice from existing slice
	oRignAl_slice := []int{90, 60, 40, 50,
		34, 49, 30}

	// Creating slices from the given slice
	var my_slice_12 = oRignAl_slice[1:5]
	my_slice_22 := oRignAl_slice[0:]
	my_slice_32 := oRignAl_slice[:6]
	my_slice_42 := oRignAl_slice[:]
	my_slice_52 := my_slice_32[2:4]
	fmt.Println("--------------------------------")
	// Display the result
	fmt.Println("Original Slice:", oRignAl_slice)
	fmt.Println("New Slice 12:", my_slice_12)
	fmt.Println("New Slice 22:", my_slice_22)
	fmt.Println("New Slice 32:", my_slice_32)
	fmt.Println("New Slice 42:", my_slice_42)
	fmt.Println("New Slice 52:", my_slice_52)
	//using make function
	var sl123 = make([]int, 4, 7)
	fmt.Printf("Slice1=%v,\nlength=%d,\ncapacity=%d\n", sl123, len(sl123), cap(sl123))
	fmt.Println("--------------------------------")
	var sl111 = make([]int, 7)
	fmt.Printf("Slice 2 = %v, \nlength = %d, \ncapacity = %d\n",
		sl111, len(sl111), cap(sl111))
	fmt.Println("--------------------------------")
	myslice1111 := []string{"This", "is", "the", "tutorial",
		"of", "Go", "language"}

	// Iterate using for loop
	for e := 0; e < len(myslice1111); e++ {
		fmt.Println(myslice1111[e])
	}
	fmt.Println("--------------------------------")
	//using range loop
	for index, ele := range myslice1111 {
		fmt.Printf("Index = %d and element = %s\n", index+3, ele)
	}
	fmt.Println("--------------------------------")
	//using range without index
	for _, ele := range myslice1111 {
		fmt.Printf("Element = %s\n", ele)
	}
	fmt.Println("--------------------------------")
	//zeroslice
	var zeroslice []string
	fmt.Printf("Length = %d\n", len(zeroslice))
	fmt.Printf("Capacity = %d\n ", cap(zeroslice))
	fmt.Println("--------------------------------")
	//modifying slice
	arr := [6]int{55, 66, 77, 88, 99, 22}
	slc := arr[0:4]

	// Before modifying

	fmt.Println("Original_Array: ", arr)
	fmt.Println("Original_Slice: ", slc)

	// After modification
	slc[0] = 100
	slc[1] = 1000
	slc[2] = 1000

	fmt.Println("\nNew_Array: ", arr)
	fmt.Println("New_Slice: ", slc)
	fmt.Println("--------------------------------")
	//comparision
	comps1 := []int{12, 34, 56}
	var comps2 []int

	// If you try to run this commented
	// code compiler will give an error
	/*s3 := []int{23, 45, 66}
	fmt.Println(s1 == s3)
	*/
	// Checking if the given slice is nil or not
	fmt.Println(comps1 == nil)
	fmt.Println(comps2 == nil)
	fmt.Println("--------------------------------")
	mds1 := [][]int{{12, 34},
		{56, 47},
		{29, 40},
		{46, 78},
	}

	// Accessing multi-dimensional slice
	fmt.Println("Slice 1 : ", mds1)

	// Creating multi-dimensional slice
	mds2 := [][]string{
		[]string{"Geeks", "for"},
		[]string{"Geeks", "GFG"},
		[]string{"gfg", "geek"},
	}

	// Accessing multi-dimensional slice
	fmt.Println("Slice 2 : ", mds2)
}
