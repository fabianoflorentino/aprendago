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

func ResolucaoNaPraticaExercicio6() {
	fmt.Printf("Sistema Operacional: %v\nArquitetura: %v\n", runtime.GOOS, runtime.GOARCH)
}

func ResolucaoNaPraticaExercicio7() {
	fmt.Println("https://dev.to/fabianoflorentino/a-interface-write-11c5")
}

type humanos interface {
	falar()
}

type pessoa struct {
	nome  string
	idade int
}

func (p *pessoa) falar() {
	fmt.Println("Olá, meu nome é", p.nome)
}

func dizerAlgumaCoisa(h humanos) {
	h.falar()
}

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
