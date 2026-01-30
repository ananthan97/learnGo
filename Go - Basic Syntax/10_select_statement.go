package basicSyntax

import "fmt"

/*
	Select Statement in Go
	The select statement is used to wait on multiple channel operations.
	It allows a goroutine to wait on multiple communication operations, proceeding with the one that is ready first.
	The syntax of a select statement is as follows:
	select {
		case communication1:
			// code to be executed if communication1 is ready
		case communication2:
			// code to be executed if communication2 is ready
		...
		default:
			// code to be executed if none of the communications are ready
	}
	Each case in a select statement must be a communication operation (send or receive) on a channel.
	If multiple cases are ready, one of them is chosen at random to proceed.
	The default case is optional and is executed if none of the other cases are ready.
	Example:
*/
func selectStatement() {
	ch1 := make(chan string) // Creating channels, channel of type string, channel is a way to communicate between goroutines
	ch2 := make(chan string) // Creating channels, goroutines are lightweight threads managed by the Go runtime
	go func() {              // Anonymous goroutine to send message to channel 1
		ch1 <- "Message from channel 1" // Sending message to channel 1, <- is the send operator
	}() // Invoking the anonymous function as a goroutine
	go func() { // Anonymous goroutine to send message to channel 2
		ch2 <- "Message from channel 2" // Sending message to channel 2
	}() // Invoking the anonymous function as a goroutine
	select {
	case msg1 := <-ch1: // Receiving message from channel 1, storing it in msg1, <- is the receive operator, ch1: receiving from channel 1, : is used at end of ch to indicate it's a case
		fmt.Println(msg1)
	case msg2 := <-ch2:
		fmt.Println(msg2)
	}
	// Example with default case
	select {
	case msg1 := <-ch1:
		fmt.Println(msg1)
	case msg2 := <-ch2:
		fmt.Println(msg2)
	default:
		fmt.Println("No messages received")
	}
}
