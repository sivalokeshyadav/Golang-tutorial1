// channel:- it is a medium through whic a goroutine communicates with another goroutine and this communications is lock-free
// or a channel is a technique which allows to let one goroutine to send data to another goroutine.By defauly channel is bidirectional,means the goroutines can send or receive data through the same channel
// in go,a channel is created using chan keyword and it can only transfer data of the same type,different types of the data are not allowed to transport from the same channel
// syntax:- var Channel_name chan Type
// using make syntax:-channel_name:=make(chan Type)
package main // Use the same package name as main.go

import "fmt"

func Channels() {
	// Creating a channel
	// Using var keyword
	var mychannel chan int
	fmt.Println("Value of the channel:", mychannel)
	fmt.Printf("Type of the channel: %T", mychannel)

	// Creating a channel using make() function
	mychannel1 := make(chan int)
	fmt.Println("\nValue of the channel1: ", mychannel1)
	fmt.Printf("Type of the channel1: %T \n", mychannel1)
}

// In Go language, channel work with two principal operations one is sending and another one is receiving,
// both the operations collectively known as communication. And the direction of <- operator indicates whether the data is received or send.
// In the channel, the send and receive operation block until another side is not ready by default. It allows goroutine to
// synchronize with each other without explicit locks or condition variables.

// Send operation: The send operation is used to send data from one goroutine to another goroutine with the help of a channel.
// Values like int, float64, and bool can safe and easy to send through a channel because they are copied so there is no risk
//  of accidental concurrent access of the same value. Similarly,
// strings are also safe to transfer because they are immutable. But for sending pointers or reference like a slice, map, etc.
// through a channel are not safe because the value of pointers or reference may change by sending goroutine or by the receiving
// goroutine at the same time and the result is unpredicted. So, when you use pointers or references in the channel you must make
// sure that they can only access by the one goroutine at a time.
// Mychannel <- element
// The above statement indicates that the data(element) send to the channel(Mychannel) with the help of a <- operator.

// Receive operation: The receive operation is used to receive the data sent by the send operator.
// element := <-Mychannel
// The above statement indicates that the element receives data from the channel(Mychannel). If the result of the received statement is not
//  going to use is also a valid statement. You can also write a receive statement as:

// <-Mychannel
func myfunc(ch chan int) {

	fmt.Println(234 + <-ch)
}

//closing a channel:You can also close a channel with the help of close() function.
// This is an in-built function and sets a flag which indicates that no more value will send to this channel.
//syntax:-close()
func myfunc1(mychnl chan string) {
	for v := 0; v < 4; v++ {
		mychnl <- "GeeksforGeeks"
	}
	close(mychnl)
}

//Important points:
//Blocking Send and Receive: In the channel when the data sent to a channel
// the control is blocked in that send statement until other goroutine
//  reads from that channel. Similarly, when a channel receives data
//  from the goroutine the read statement block until another goroutine statement
//Zero Value Channel: The zero value of the channel is nil.
//For loop in Channel: A for loop can iterate over the sequential values sent on the channel until it closed.
//Length of the Channel: In channel, you can find the length of the channel
//  using len() function. Here, the length indicates the number of value
// queued in the channel buffer.
//Capacity of the Channel: In channel, you can find the capacity of the channel using cap() function.
//  Here, the capacity indicates the size of the buffer.
//Select and case statement in Channel: In go language, select statement
//  is just like a switch statement without any input parameter.
//  This select statement is used in the channel to perform a single
//  operation out of multiple operations provided by the case block.
//unidirectional channel in golang
//Use of Unidirectional Channel: The unidirectional channel is used to provide the type-safety
//  of the program so, that the program gives less error. Or you can also
//  use a unidirectional channel when you want to create a channel that can
//  only send or receive data.
func unidirectionalFlow() {
	//onlyforreceiving
	mychanl1 := make(<-chan string)
	//onlyforsending
	mychanl2 := make(chan<- string)
	//displaythetypesofchannels
	fmt.Printf("%T", mychanl1)
	fmt.Printf("\n%T\n", mychanl2)
}

//converting bidirectional to unidirectional
func sending(s chan<- string) {
	s <- "GeeksforGeeks"
}
