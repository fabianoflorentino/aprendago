package exercicios_ninja_nivel_10

import "fmt"

// ResolucaoNaPraticaExercicio1 demonstrates a simple use of a goroutine and a channel in Go.
// It creates an unbuffered channel of type int and starts a goroutine that sends the value 42 into the channel.
// The main function then receives the value from the channel and prints it using fmt.Printf.
func ResolucaoNaPraticaExercicio1() {
	channel := make(chan int)
	go func() {
		channel <- 42
	}()

	fmt.Printf("Value: %v\n", <-channel)
}

// ResolucaoNaPraticaExercicio2 demonstrates the use of unidirectional channels in Go.
// It creates a bidirectional channel and converts it to a send-only channel.
// A goroutine is used to send a value (42) to the channel.
// The main function then receives the value from the channel and prints it.
// Finally, it prints the type of the send-only channel.
func ResolucaoNaPraticaExercicio2() {
	channel := make(chan int)
	channel_send := chan<- int(channel)

	go func() {
		channel_send <- 42
	}()

	fmt.Printf("Value: %v\n", <-channel)
	fmt.Printf("--------------------\n")
	fmt.Printf("Type: %T\n", channel_send)
}

// ResolucaoNaPraticaExercicio3 demonstrates the process of generating a channel,
// receiving values from it, and then printing a message before exiting.
// It first calls generateChannels() to create a channel, then passes this channel
// to receiveChannel() to handle the received values. Finally, it prints a message
// indicating that the program is about to exit.
func ResolucaoNaPraticaExercicio3() {
	channel := generateExercicio3()
	receiveExercicio3(channel)

	fmt.Printf("About to exit...\n")
}

func ResolucaoNaPraticaExercicio4() {
	slct := make(chan int)
	channel := generateExercicio4(slct)

	go receiveExercicio4(channel, slct)

	fmt.Println("About to exit...")
}

func ResolucaoNaPraticaExercicio5() {
	channel := make(chan int)

	go func() { channel <- 42 }()

	value, ok := <-channel
	fmt.Printf("Value: %v\n", value)
	fmt.Printf("Channel open: %v\n", ok)
	close(channel)

	value, ok = <-channel
	fmt.Printf("Value: %v\n", value)
	fmt.Printf("Channel open: %v\n", ok)
}

// ResolucaoNaPraticaExercicio6 demonstrates the use of goroutines and channels in Go.
// It creates a channel to communicate between a goroutine and the main function.
// The goroutine sends integers from 0 to 100 to the channel and then closes the channel.
// The main function receives and prints these integers from the channel until it is closed.
func ResolucaoNaPraticaExercicio6() {
	channel := make(chan int)

	go func() {
		for idx := 0; idx < 101; idx++ {
			channel <- idx
		}
		close(channel)
	}()

	for value := range channel {
		fmt.Printf("Channel: %v\n", value)
	}
}

// ResolucaoNaPraticaExercicio7 demonstrates the use of goroutines and channels in Go.
// It creates a channel to communicate between goroutines and the main function.
// The function starts 10 goroutines, each sending the numbers 0 to 9 to the channel.
// The main function then receives and prints 100 numbers from the channel.
func ResolucaoNaPraticaExercicio7() {
	channel := make(chan int)

	for idx := 0; idx < 10; idx++ {
		go func() {
			for number := 0; number < 10; number++ {
				channel <- number
			}
		}()
	}

	for number := 0; number < 100; number++ {
		fmt.Println(number, "\t", <-channel)
	}
}

// generateChannels creates and returns a read-only channel of integers.
// It launches 100 goroutines, each sending its index value to the channel.
// The channel is not closed, so the caller must handle the channel closure if needed.
func generateExercicio3() <-chan int {
	channel := make(chan int)

	go func() {
		for idx := 0; idx < 100; idx++ {
			channel <- idx
		}
		close(channel)
	}()

	return channel
}

// receiveChannel reads values from a read-only integer channel and prints each value to the standard output.
// It continuously listens to the channel until it is closed, printing each received value with a formatted message.
func receiveExercicio3(channel <-chan int) {
	for value := range channel {
		fmt.Printf("Value: %v\n", value)
	}
}

// generateExercicio4 creates and returns a read-only channel that emits integers from 0 to 99.
// It also takes a send-only channel as an argument, which it uses to signal completion by sending a value of 0.
// The function launches a goroutine that iterates from 0 to 99, sending each value to the returned channel.
// After sending all values, the channel is closed and a signal is sent to the provided channel to indicate completion.
func generateExercicio4(slct chan<- int) <-chan int {
	channel := make(chan int)

	go func() {
		for idx := 0; idx < 100; idx++ {
			channel <- idx
		}
		close(channel)
		slct <- 0
	}()

	return channel
}

// receiveExercicio4 listens to two channels and prints the received values.
// It runs an infinite loop where it uses a select statement to wait for values
// from either the 'channel' or 'slct' channels. When a value is received from
// either channel, it prints the value to the standard output.
//
// Parameters:
// - channel: a read-only channel of integers from which values are received.
// - slct: a read-only channel of integers from which values are received.
func receiveExercicio4(channel <-chan int, slct chan int) {
	for {
		select {
		case value := <-channel:
			fmt.Printf("Value: %v\n", value)
		case value := <-slct:
			fmt.Printf("Value: %v\n", value)
		}
	}
}
