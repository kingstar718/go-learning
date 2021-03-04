package main

import "fmt"

/**
Channels are the pipes that connect concurrent goroutines
*/

func main() {

	// Create a new channel with make(chan val-type)
	messages := make(chan string)

	// Send a value into a channel using the channel <- syntax
	go func() { messages <- "ping" }()

	// The <-channel syntax receives a values from the channel
	msg := <-messages
	fmt.Println(msg)
}
