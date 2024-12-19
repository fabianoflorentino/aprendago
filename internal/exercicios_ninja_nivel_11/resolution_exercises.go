// Package exercicios_ninja_nivel_11 contains solutions to various exercises
// designed to practice and enhance Go programming skills. The exercises cover
// different aspects of the language, including struct manipulation, JSON
// serialization, error handling, and more.
package exercicios_ninja_nivel_11

import (
	"encoding/json"
	"fmt"
	"log"
)

// person represents an individual with a first name, last name, and a list of sayings.
type person struct {
	First   string
	Last    string
	Sayings []string
}

// erroEspecial is a custom error type that wraps an existing error message.
// It is used to provide additional context or special handling for errors.
type erroEspecial struct {
	msg error
}

// sqrtError represents a custom error type that includes additional context
// about the location (latitude and longitude) where the error occurred,
// along with the original error.
type sqrtError struct {
	lat  string
	long string
	err  error
}

// ResolucaoNaPraticaExercicio1 serializes a 'person' struct into JSON format and prints it.
// If an error occurs during the marshaling process, the function will panic.
func ResolucaoNaPraticaExercicio1() {
	p1 := person{
		First:   "James",
		Last:    "Bond",
		Sayings: []string{"Shaken, not stirred", "Any last wishes?", "Never say never"},
	}

	bs, err := json.Marshal(p1)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bs))
}

// ResolucaoNaPraticaExercicio2 demonstrates the creation of a 'person' struct instance,
// converts it to JSON format using the toJSON function, and prints the resulting JSON string.
// If an error occurs during the conversion, it prints the error.
func ResolucaoNaPraticaExercicio2() {
	p1 := person{
		First:   "James",
		Last:    "Bond",
		Sayings: []string{"Shaken, not stirred", "Any last wishes?", "Never say never"},
	}

	bs, err := toJSON(p1)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(bs))
}

// ResolucaoNaPraticaExercicio3 demonstrates the creation of a custom error type
// "erroEspecial" that implements the built-in error interface. It also shows
// how to create a function that accepts an error type as a parameter and how
// to pass an instance of the custom error type to this function.
func ResolucaoNaPraticaExercicio3() {
	errEspecial := erroEspecial{
		msg: fmt.Errorf("erro especial"),
	}

	err := errEspecial.Error()

	fmt.Println(err)
}

func ResolucaoNaPraticaExercicio4() {
	_, err := sqrt(-10.23)
	if err != nil {
		log.Println(err)
	}
}

func ResolucaoNaPraticaExercicio5() {
	var content = `
- Nos capítulos seguintes, uma das coisas que veremos é testes.
- Para testar sua habilidade de se virar por conta própria... desafio:
  - Utilizando as seguintes fontes: https://godoc.org/testing & http://www.golang-book.com/books/intro/12
  - Tente descobrir por conta própria como funcionam testes em Go.
  - Pode usar tradutor automático, pode rodar código na sua máquina, pode procurar no Google. Vale tudo.
  - O exercício é: crie um teste simples de uma função ou método ou pedaço qualquer de código.
`
	resolution := "internal/exercicios_ninja_nivel_11/resolution_exercises_test.go"

	fmt.Println(content)
	fmt.Printf("Resolução: %v", resolution)
}

// toJSON converts a person struct to its JSON representation.
// It returns the JSON byte slice and an error if the conversion fails.
func toJSON(p person) ([]byte, error) {
	bs, err := json.Marshal(p)
	if err != nil {
		return []byte{}, fmt.Errorf("erro ao exibir as informações da pessoa")
	}
	return bs, nil
}

// Error implements the error interface for the erroEspecial type.
// It returns the error message contained within the erroEspecial instance.
func (e erroEspecial) Error() string {
	return e.msg.Error()
}

// Error implements the error interface for the sqrtError type.
// It returns a formatted string containing the latitude, longitude, and the original error message.
func (se sqrtError) Error() string {
	return fmt.Sprintf("math error: %v %v %v", se.lat, se.long, se.err)
}

// sqrt calculates the square root of a given float64 number.
// If the input number is negative, it returns an error.
func sqrt(f float64) (float64, error) {
	if f < 0 {
		// write your error code here
		erro := fmt.Errorf("número negativo: %v", f)
		return 0, sqrtError{"50.2289 N", "99.4656 W", erro}
	}
	return 42, nil
}
