package main

import (
	"fmt"
	"time"
)

func display(str string) {
	for i := 0; i < 3; i++ {
		fmt.Println(str)
	}

}
func task1(ch chan string) {
	time.Sleep(2 * time.Second)
	ch <- "Task1 completed"
}
func task2(ch chan string) {
	time.Sleep(4 * time.Second)
	ch <- "Task2 completed"
}
func portal1(channel1 chan string) {
	time.Sleep(3 * time.Second)
	channel1 <- "Welcome to channel 1"
}
func portal2(channel2 chan string) {
	time.Sleep(9 * time.Second)
	channel2 <- "Welcome to channel 2"
}
func Aname() {
	arr1 := [4]string{"Rohit", "Suman", "Aman", "Ria"}
	for t1 := 0; t1 <= 3; t1++ {
		time.Sleep(150 * time.Millisecond)
		fmt.Printf("%s\n", arr1[t1])
	}
}
func Aid() {
	arr2 := [4]int{300, 301, 302, 303}
	for t2 := 0; t2 <= 3; t2++ {
		time.Sleep(500 * time.Millisecond)
		fmt.Printf("%d\n", arr2[t2])
	}
}
func main() {
	// Goroutines in Go let functions run concurrently,
	// using less memory than traditional threads.
	// Every Go program starts with a main Goroutine,
	// and if it exits, all others stop.

	// go display("Hello, Goroutine!") // Runs concurrently// Using `go` to run as a Goroutine
	// display("Hello, Main!")
	//beforeusingtime.sleepoutputis
	//Hello, Main!
	//Hello, Main!
	// Hello, Main!
	//time.Sleep(time.Second) // Allow time for the goroutine to finish
	//beforeusingtime.sleepoutputis
	// Hello, Main!
	// Hello, Main!
	// Hello, Main!
	// Hello, Goroutine!
	// Hello, Goroutine!
	// Hello, Goroutine!
	//becausethe main goroutine may exit before the goroutine has a chance to complete, so "Hello, Goroutine!" may not appear at all in the output
	//To create a Goroutine, prefix the function or method call with the go keyword.

	//Syntax:-
	//func functionName(){
	// statements
	//}
	// Using `go` to run as a Goroutine
	//go functionName()
	//anonymous goroutines
	//gofunc(parameters){
	//statement or functionlogic
	// }(arguments)
	//eg:-
	// go func(str string) {
	// 	for i := 0; i < 3; i++ {
	// 		fmt.Println(str)
	// 		time.Sleep(500 * time.Millisecond)
	// 	}
	// }("Hello from Anonymous Goroutine!")
	// time.Sleep(2 * time.Second) //AllowGoroutinetofinish
	// fmt.Println("Main function complete")

	//In Go, the select statement allows you to wait on multiple channel operations, such as sending or receiving values. Similar to a switch statement, select enables you to proceed with whichever channel case is ready, making it ideal for handling asynchronous tasks efficiently.
	//Consider a scenario where two tasks complete at different times. We’ll use select to receive data from whichever task finishes first.
	// ch1 := make(chan string)
	// ch2 := make(chan string)
	// go task1(ch1)
	// go task2(ch2)
	// select {
	// case msg1 := <-ch1:
	// 	fmt.Println(msg1)
	// case msg2 := <-ch2:
	// 	fmt.Println(msg2)
	// }
	//Note:-After 2 seconds, Task 1 completed will be printed because task1 completes before task2. If task2 finished first, then Task 2 completed would be printed.
	//keypoints
	//select waits until at least one channel operation is ready.
	//If multiple cases are ready, one is chosen at random.
	//The default case executes if no other case is ready, avoiding a block.
	//selectstatementvariations
	//defaultcase
	// ch := make(chan string)
	// select {
	// case msg := <-ch:
	// 	fmt.Println(msg)
	// default:
	// 	fmt.Println("No Channels are ready")
	// }
	// //multiplecasesarereadyatsametime
	// R1 := make(chan string)
	// R2 := make(chan string)
	// //callingfunction1andfunction2ingoroutine
	// go portal1(R1)
	// go portal2(R2)
	// select {
	// case op1 := <-R1:
	// 	fmt.Println(op1)
	// case op2 := <-R2:
	// 	fmt.Println(op2)
	// }
	//In this scenario, select randomly picks one of the two cases if both are ready at the same time, meaning you may see "Task 1 completed" or "Task 2 completed" depending on the selection.
	// ch1 := make(chan string)
	// ch2 := make(chan string)
	// select {
	// case msg1 := <-ch1:
	// 	fmt.Println(msg1)
	// case msg2 := <-ch2:
	// 	fmt.Println(msg2)
	// default:
	// 	fmt.Println("No tasks are ready yet")
	// }
	//Since neither ch1 nor ch2 is ready, the default case executes, outputting "No tasks are ready yet".
	//multiplegoroutines
	// fmt.Println("!....Main Go-routine Start ...!")
	// //calling goroutine1
	// go Aname()
	// //calling goroutine2
	// go Aid()

	// time.Sleep(3500 * time.Millisecond)
	// fmt.Println("\n!...Main Go-routine End...!")
	Channels()
	fmt.Println("start Main method")
	// Creating a channel
	ch := make(chan int)

	go myfunc(ch)
	ch <- 23
	fmt.Println("End Main method")
	c := make(chan string)
	go myfunc1(c)
	// When the value of ok is
	// set to true means the
	// channel is open and it
	// can send or receive data
	// When the value of ok is set to
	// false means the channel is closed
	for {
		res, ok := <-c
		if ok == false {
			fmt.Println("Channel Close", ok)
			break
		}
		fmt.Println("Channel Open", res, ok)
	}
	unidirectionalFlow()
	//bidirectionaltounidirectionalflow
	bitouni := make(chan string)
	// Here, sending() function convert
	// the bidirectional channel
	// into send only channel
	go sending(bitouni)

	// Here, the channel is sent
	// only inside the goroutine
	// outside the goroutine the
	// channel is bidirectional
	// So, it print GeeksforGeeks
	fmt.Println(<-bitouni)
}
