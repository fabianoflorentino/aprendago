// Package canais provides examples of using Go channels to synchronize and
// communicate between goroutines. It includes functions to demonstrate
// sending and receiving values through channels, as well as converging
// multiple channels into a single channel for processing.
package canais

import (
	"fmt"
	"sync"
	"time"
)

// wg is a WaitGroup used to wait for a collection of goroutines to finish executing.
// It provides a way to synchronize the completion of multiple goroutines by using
// Add, Done, and Wait methods. The Add method increments the counter by the number
// of goroutines to wait for, the Done method decrements the counter when a goroutine
// completes, and the Wait method blocks until the counter becomes zero.
var wg sync.WaitGroup

// UsingConverge demonstrates the use of channels to converge values from two separate channels (odd and even) into a single channel (converge).
// It starts two goroutines: one to send values to the odd and even channels, and another to receive values from these channels and send them to the converge channel.
// The function then iterates over the values received from the converge channel, printing each value as either "Even Value" or "Odd Value".
// If the value is divisible by 5, a newline is added after the value for better readability.
func UsingConverge() {
	odd := make(chan int)
	even := make(chan int)
	converge := make(chan int)

	go sentToConverge(odd, even)
	go receiveToConverge(odd, even, converge)

	for value := range converge {

		if value%2 == 0 {
			fmt.Printf("Even Value: %v  ", value)
		} else {

			fmt.Printf("Odd Value: %v  ", value)
		}
	}
}

// UsingConvergeString demonstrates how to converge two string channels into a single channel.
// It creates two channels, chan1 and chan2, by calling workWithChanString with "A" and "B" respectively.
// These channels are then passed to receiveChanStringToConverge, which merges them into a new channel, newChan.
// The function iterates over newChan, printing each value received from the merged channels.
func UsingConvergeString() {
	chan1 := workWithChanString("A")
	chan2 := workWithChanString("B")

	newChan := receiveChanStringToConverge(chan1, chan2)

	fmt.Printf("\nUse Ctrl+C to interrupt the program!!\n\n")
	for value := range newChan {
		fmt.Printf("%v\t", value)
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

// receiveToConverge takes three channels as input: odd, even, and converge.
// It starts two goroutines to read values from the odd and even channels and send them to the converge channel.
// The function uses a WaitGroup to ensure both goroutines complete their execution before closing the converge channel.
// This function is useful for merging values from two separate channels into a single channel.
func receiveToConverge(odd, even, converge chan int) {
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

// workWithChanString takes a string value and returns a channel of strings.
// It starts a goroutine that continuously sends formatted strings to the channel.
// Each string is the input value concatenated with an incrementing index, separated by a space.
// The goroutine sleeps for 2 seconds between each send operation.
func workWithChanString(value string) chan string {
	stringChan := make(chan string)

	go func(value string, chann chan string) {
		for idx := 0; ; idx++ {
			chann <- fmt.Sprintf("%s %d", value, idx)
			time.Sleep(time.Second * 2)
		}
	}(value, stringChan)

	return stringChan
}

// receiveChanStringToConverge converges two input channels (chan1 and chan2) into a single output channel (newChan).
// It launches two goroutines, each of which continuously reads from one of the input channels and sends the received
// value to the output channel. This allows values from both input channels to be received on the single output channel.
//   - The line `newChan <- <-chan2` reads a value from chan2 and sends it to newChan.
func receiveChanStringToConverge(chan1, chan2 chan string) chan string {
	newChan := make(chan string)

	go func() {
		for {
			newChan <- <-chan1
		}
	}()
	go func() {
		for {
			newChan <- <-chan2
		}
	}()

	return newChan
}
