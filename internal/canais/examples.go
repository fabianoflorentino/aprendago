// Package canais provides examples of using Go channels to synchronize and
// communicate between goroutines. It includes functions to demonstrate
// sending and receiving values through channels, as well as converging
// multiple channels into a single channel for processing.
package canais

import (
	"fmt"
	"sync"
)

// wg is a WaitGroup used to wait for a collection of goroutines to finish executing.
// It provides a way to synchronize the completion of multiple goroutines by using
// Add, Done, and Wait methods. The Add method increments the counter by the number
// of goroutines to wait for, the Done method decrements the counter when a goroutine
// completes, and the Wait method blocks until the counter becomes zero.
var wg sync.WaitGroup

// breakLine is a string variable used to represent a line break or separator in the output.
var breakLine string

// UsingConverge demonstrates the use of channels to converge values from two separate channels (odd and even) into a single channel (converge).
// It starts two goroutines: one to send values to the odd and even channels, and another to receive values from these channels and send them to the converge channel.
// The function then iterates over the values received from the converge channel, printing each value as either "Even Value" or "Odd Value".
// If the value is divisible by 5, a newline is added after the value for better readability.
func UsingConverge() {
	odd := make(chan int)
	even := make(chan int)
	converge := make(chan int)

	go sentToConverge(odd, even)
	go receivetoConverge(odd, even, converge)

	for value := range converge {

		if value%5 == 0 {
			breakLine = "\n"
		}

		if value%2 == 0 {
			fmt.Printf("Even Value: %v  %v", value, breakLine)
		} else {
			fmt.Printf("Odd Value: %v  %v", value, breakLine)
		}
	}
}

// sentToConverge sends numbers from 0 to 99 to either the odd or even channel.
// If the number is even, it is sent to the even channel; otherwise, it is sent to the odd channel.
// After sending all numbers, both channels are closed.
func sentToConverge(odd, even chan int) {
	for number := 0; number < 100; number++ {
		if number%2 == 0 {
			even <- number
		} else {
			odd <- number
		}
	}
	close(odd)
	close(even)
}

// receivetoConverge takes three channels as input: odd, even, and converge.
// It starts two goroutines to read values from the odd and even channels and send them to the converge channel.
// The function uses a WaitGroup to ensure both goroutines complete their execution before closing the converge channel.
// This function is useful for merging values from two separate channels into a single channel.
func receivetoConverge(odd, even, converge chan int) {
	wg.Add(1)
	go func() {
		for value := range odd {
			converge <- value
		}
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		for value := range even {
			converge <- value
		}
		wg.Done()
	}()

	wg.Wait()
	close(converge)
}
