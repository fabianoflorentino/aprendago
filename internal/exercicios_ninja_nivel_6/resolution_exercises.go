// Package exercicios_ninja_nivel_6 contains a series of functions and types
// designed to demonstrate various Go programming concepts through practical exercises.
// These exercises cover a range of topics including function usage, anonymous functions,
// methods, interfaces, and more. Each function is an example of a specific concept,
// providing a hands-on approach to learning Go.
package exercicios_ninja_nivel_6

import "fmt"

// figura is an interface that defines a geometric shape with a method area.
// The area method should be implemented by any type that represents a shape
// and should return the area of the shape as a float64.
type figura interface {
	area() float64
}

// quadrado represents a square with a specific side length.
// The lado field holds the length of the side of the square.
type quadrado struct {
	lado float64
}

// circulo represents a circle with a specific radius.
// The raio field holds the radius of the circle as a float64 value.
type circulo struct {
	raio float64
}

// pessoa represents a person with a first name, last name, and age.
// Fields:
// - nome: the first name of the person.
// - sobrenome: the last name of the person.
// - idade: the age of the person.
type pessoa struct {
	nome      string
	sobrenome string
	idade     int
}

// mostrarPessoa prints the details of a pessoa instance to the standard output.
// It displays the first name, last name, and age of the person in a formatted string.
func (p pessoa) mostrarPessoa() {
	fmt.Printf("Nome: %s %s\nIdade: %d\n", p.nome, p.sobrenome, p.idade)
}

// area calculates and returns the area of a square.
// It multiplies the length of one side (lado) by itself to get the area.
func (q quadrado) area() float64 {
	return q.lado * q.lado
}

// area calculates the area of a circle based on its radius.
// It uses the formula 2 * π * radius to compute the area.
// Returns the area as a float64.
func (c circulo) area() float64 {
	return 2 * 3.14 * c.raio
}

// ResolucaoNaPraticaExercicio1 demonstrates the usage of two functions: retornaInt and retornaIntString.
// It retrieves an integer from retornaInt and both an integer and a string from retornaIntString.
// The function then prints these values using fmt.Printf.
func ResolucaoNaPraticaExercicio1() {
	numero_int := retornaInt()
	numero, texto := retornaIntString()

	fmt.Printf("Inteiro: %v\nInteiro&Texto: %v e %v", numero_int, numero, texto)
}

// ResolucaoNaPraticaExercicio2 demonstrates the use of variadic functions and slice functions in Go.
// It creates two slices of integers and calculates their sums using two different functions: somaVariadico and somaSlice.
// The results are then printed to the console.
func ResolucaoNaPraticaExercicio2() {
	slice1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	result := somaVariadico(slice1...)
	fmt.Printf("Soma do variádico: %v", result)

	slice2 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	result2 := somaSlice(slice2)
	fmt.Printf("\nSoma do slice: %v", result2)
}

// ResolucaoNaPraticaExercicio3 demonstrates the use of the defer statement in Go.
// The defer statement defers the execution of a function until the surrounding function returns.
// In this example, the deferred function prints a message indicating that it was executed at the end of the function's context.
// The first Println statement executes immediately, while the deferred Println statement executes just before the function exits.
func ResolucaoNaPraticaExercicio3() {
	defer fmt.Println("Execução do defer ocorreu ao final do contexto.")

	fmt.Println("Execução do defer ocorre ao final do contexto ao qual ela pertence.")
}

// ResolucaoNaPraticaExercicio4 demonstrates the creation and usage of a 'pessoa' struct instance.
// It initializes a 'pessoa' with predefined values for 'nome', 'sobrenome', and 'idade' fields,
// and then calls the 'mostrarPessoa' method to display the details of the 'pessoa' instance.
func ResolucaoNaPraticaExercicio4() {

	p := pessoa{
		nome:      "Fabiano",
		sobrenome: "Florentino",
		idade:     39,
	}

	p.mostrarPessoa()
}

// ResolucaoNaPraticaExercicio5 calculates and prints the area of a square and a circle.
// It creates a square with a side length of 10 and a circle with a radius of 5.
// The function then uses the info function to calculate the area of each shape
// and prints the results using formatted output.
func ResolucaoNaPraticaExercicio5() {

	q := quadrado{
		lado: 10,
	}

	c := circulo{
		raio: 5,
	}

	fmt.Printf("\nÁrea do quadrado: %f", info(q))
	fmt.Printf("\nÁrea do círculo: %f", info(c))
}

// ResolucaoNaPraticaExercicio6 demonstrates the use of an anonymous function in Go.
// An anonymous function is a function without a name that can be defined and executed inline.
// They are useful for quick tasks and can be passed as arguments to other functions.
// In this example, the anonymous function prints a message to the console when executed.
func ResolucaoNaPraticaExercicio6() {

	// An anonymous function is a function without a name that can be defined and executed inline.
	// They are useful for quick tasks and can be passed as arguments to other functions.
	func() {
		fmt.Println("Anonymous function executed.")
	}()
}

// ResolucaoNaPraticaExercicio7 demonstrates how to assign a function to a variable
// and then call that function. It prints a message to the console when executed.
func ResolucaoNaPraticaExercicio7() {

	// Função atribuída a uma variável.
	f := func() {
		fmt.Println("Função atribuída a uma variável.")
	}

	f()
}

// ResolucaoNaPraticaExercicio8 calls a function that returns another function,
// and then executes the returned function.
func ResolucaoNaPraticaExercicio8() {
	f := retornaFuncao()
	f()
}

// ResolucaoNaPraticaExercicio9 demonstrates how to pass a function as an argument to another function.
// It defines an anonymous function that prints a message and then passes this function to executaFuncao.
func ResolucaoNaPraticaExercicio9() {
	f := func() {
		fmt.Println("Função passada como argumento.")
	}

	executaFuncao(f)
}

// ResolucaoNaPraticaExercicio10 demonstrates the use of an anonymous function
// to print the value of a variable from its enclosing scope. The function
// initializes a variable `x` with the value 10, and then defines and immediately
// invokes an anonymous function that prints the value of `x` to the console.
func ResolucaoNaPraticaExercicio10() {
	x := 10

	func() {
		fmt.Println(x)
	}()
}

// ResolucaoNaPraticaExercicio11 prints a multi-line string containing instructions
// for a practical exercise. The exercise involves choosing a favorite piece of code
// related to functions, recording a video explaining the topic, uploading the video
// to YouTube, and sharing the video on Twitter while tagging a specific user.
func ResolucaoNaPraticaExercicio11() {
	var content string = `
- Uma das melhores maneiras de aprender é ensinando.
- Para este exercício escolha o seu código favorito dentre os que vimos estudando funções. Pode ser das aulas ou da seção de exercícios. Então:
		- Faça download e instale isso aqui: https://obsproject.com/
		- Grave um vídeo onde *você* ensina o tópico em questão
		- Faça upload do vídeo no YouTube
		- Compartilhe o vídeo no Twitter e me marque no tweet (@ellenkorbes)
	`

	fmt.Println(content)
}

// retornaInt returns an integer value of 42.
// This function does not take any parameters and always returns the same value.
func retornaInt() int {
	return 42
}

// retornaIntString returns an integer and a string.
// The integer returned is 42, and the string returned is "Olá, mundo!".
func retornaIntString() (int, string) {
	return 42, "Olá, mundo!"
}

// somaVariadico takes a variadic number of integer arguments and returns their sum.
// It initializes a variable 'soma' to 0, then iterates over each integer in the variadic
// parameter 'x', adding each integer to 'soma'. Finally, it returns the total sum.
func somaVariadico(x ...int) int {
	soma := 0

	for _, inteiro := range x {
		soma += inteiro
	}
	return soma
}

// somaSlice takes a slice of integers and returns the sum of all the integers in the slice.
// It iterates over each integer in the slice, adding each integer to a running total (soma),
// which is then returned as the final result.
func somaSlice(slice []int) int {
	soma := 0

	for _, inteiro := range slice {
		soma += inteiro
	}
	return soma
}

// info takes a figura interface as an argument and returns the area of the figura.
// The figura interface must have an area method that returns a float64.
func info(f figura) float64 {
	return f.area()
}

// retornaFuncao returns a function that, when called, prints "Função retornada." to the standard output.
// This demonstrates how to return a function from another function in Go.
func retornaFuncao() func() {
	return func() {
		fmt.Println("Função retornada.")
	}
}

// executaFuncao takes a function as an argument and executes it.
// The function f is expected to have no parameters and no return value.
// This allows for the execution of any function that matches this signature.
func executaFuncao(f func()) {
	f()
}
