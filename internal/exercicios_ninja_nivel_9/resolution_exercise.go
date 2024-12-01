package exercicios_ninja_nivel_9

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var ws sync.WaitGroup
var mu sync.Mutex
var contador int

const numberOfGoroutines = 50000

// ResolucaoNaPraticaExercicio1 demonstrates the use of goroutines and WaitGroup in Go.
// It creates 10 goroutines, each printing a message to the console.
// The WaitGroup is used to wait for all goroutines to complete before the function returns.
// Note: The variable 'i' is captured by reference in the goroutine, which may lead to unexpected results.
func ResolucaoNaPraticaExercicio1() {
	var ws sync.WaitGroup
	ws.Add(10)

	for i := 0; i < 10; i++ {
		go func() {
			defer ws.Done()
			fmt.Println("Goroutine", i)
		}()
	}

	ws.Wait()
}

// ResolucaoNaPraticaExercicio2 creates an instance of the 'pessoa' struct with the name "fabiano" and age 39,
// and then calls the 'dizerAlgumaCoisa' function, passing a pointer to this instance.
// ResolucaoNaPraticaExercicio2 demonstrates the creation of a 'pessoa' struct instance
// and the usage of the 'dizerAlgumaCoisa' function. The 'pessoa' struct has fields 'nome'
// and 'idade'. The 'dizerAlgumaCoisa' function takes a pointer to a 'pessoa' instance
// and performs some operation on it. This function showcases how to work with structs
// and pointers in Go.
func ResolucaoNaPraticaExercicio2() {
	p := pessoa{
		nome:  "fabiano",
		idade: 39,
	}

	dizerAlgumaCoisa(&p)
}

// ResolucaoNaPraticaExercicio3 demonstrates a concurrency issue with the use of goroutines and a shared variable.
// The function initializes a WaitGroup to wait for 100 goroutines to complete their execution.
// Each goroutine increments a shared counter variable `contador`.
// However, due to the lack of synchronization mechanisms (like mutexes), the increments to `contador` are not atomic,
// leading to a race condition. As a result, the final value of `contador` may not be 100 as expected.
// The function prints the final value of `contador` after all goroutines have completed.
func ResolucaoNaPraticaExercicio3() {
	var ws sync.WaitGroup
	ws.Add(100)

	var contador int

	for i := 0; i < 100; i++ {
		go func() {
			defer ws.Done()
			valor := contador
			valor++
			contador = valor
		}()
	}

	ws.Wait()

	fmt.Println("Contador:", contador)
}

// ResolucaoNaPraticaExercicio4 demonstrates a concurrency example using goroutines, WaitGroup, and Mutex in Go.
// The function initializes a WaitGroup to wait for 100 goroutines to complete their execution.
// A shared counter variable is incremented by each goroutine in a thread-safe manner using a Mutex to avoid race conditions.
// The Mutex ensures that only one goroutine can access and modify the counter at a time.
// The function prints the value of the counter after each increment, but the print statement is outside the critical section,
// which may lead to inconsistent output due to concurrent access.
func ResolucaoNaPraticaExercicio4() {
	buildGoRoutine(numberOfGoroutines)
	ws.Wait()

	fmt.Println("All goroutines completed: ", contador)
}

// ResolucaoNaPraticaExercicio5 demonstrates the use of the atomic package to fix
// a race condition in a concurrent program. The function initializes a counter
// variable and a WaitGroup to synchronize 100 goroutines. Each goroutine increments
// the counter atomically using atomic.AddInt32, ensuring that the increment operation
// is thread-safe. The WaitGroup is used to wait for all goroutines to complete before
// printing the final value of the counter.
func ResolucaoNaPraticaExercicio5() {
	// Utilize atomic para consertar a condição de corrida do exercício #3.
	var contador int32
	var ws sync.WaitGroup

	ws.Add(100)

	for i := 0; i < 100; i++ {
		go func() {
			defer ws.Done()
			atomic.AddInt32(&contador, 1)
		}()
	}

	ws.Wait()

	fmt.Println("Contador:", contador)
}

// ResolucaoNaPraticaExercicio6 prints the operating system and architecture
// of the machine where the program is running. It uses the runtime package
// to retrieve the OS and architecture information and prints them using
// the fmt.Printf function.
func ResolucaoNaPraticaExercicio6() {
	fmt.Printf("Sistema Operacional: %v\nArquitetura: %v\n", runtime.GOOS, runtime.GOARCH)
}

// ResolucaoNaPraticaExercicio7 prints a URL to the console.
// The URL points to an article written by Fabiano Florentino
// about the Write interface in Go.
func ResolucaoNaPraticaExercicio7() {
	fmt.Println("https://dev.to/fabianoflorentino/a-interface-write-11c5")
}

// humanos is an interface that requires the implementation of the
// falar method. Any type that implements the falar method can be
// considered a humanos. The falar method is intended to define
// a way for the type to "speak" or communicate.
type humanos interface {
	falar()
}

// pessoa represents a person with a name and age.
type pessoa struct {
	nome  string
	idade int
}

// falar is a method for the pessoa type that prints a greeting message
// including the person's name to the standard output.
func (p *pessoa) falar() {
	fmt.Println("Olá, meu nome é", p.nome)
}

// dizerAlgumaCoisa takes an interface of type humanos and calls its falar method.
// This function demonstrates polymorphism by allowing any type that implements
// the humanos interface to be passed in and have its falar method invoked.
func dizerAlgumaCoisa(h humanos) {
	h.falar()
}

// buildGoRoutine launches a specified number of goroutines that increment a shared counter.
// Each goroutine increments the counter in a thread-safe manner using a mutex lock.
// The function takes an integer 'i' which determines the number of goroutines to be launched.
// It also adds 'i' to a WaitGroup to ensure all goroutines complete before the program exits.
//
// Parameters:
// - i: The number of goroutines to launch.
//
// Note: The function assumes that 'mu' is a mutex, 'contador' is a shared counter variable,
// and 'ws' is a WaitGroup that are defined elsewhere in the package.
func buildGoRoutine(i int) {
	ws.Add(i)

	for j := 0; j < i; j++ {
		go func() {
			defer mu.Unlock()
			v := contador
			v++
			contador = v

			mu.Lock()
			ws.Done()
		}()
	}
}
