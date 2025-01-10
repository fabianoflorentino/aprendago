// Package exercicios_ninja_nivel_4 contains solutions for various exercises
// at level 4 of the Ninja Go programming course. Each function within this
// package demonstrates different Go programming concepts and techniques,
// such as working with arrays, slices, maps, and string manipulation.
package exercicios_ninja_nivel_4

import (
	"fmt"
	"strings"

	"github.com/fabianoflorentino/aprendago/pkg/format"
)

// ResolucaoNaPraticaExercicio1 demonstrates a simple example of iterating over an array of integers,
// converting each integer to a string, and concatenating them into a single comma-separated string.
// The resulting string is then formatted and processed by the FormatResolucaoExercicios function.
// The function uses fmt.Sprintf to format each integer and strings.Trim to remove the trailing comma and space.
func ResolucaoNaPraticaExercicio1() {
	array := [5]int{1, 2, 3, 4, 5}
	var resolucao string

	for _, value := range array {
		resolucao += fmt.Sprintf("%v, ", value)
	}

	resolucao = strings.Trim(resolucao, ", ")
	format.FormatResolucaoExercicios(resolucao)
}

// ResolucaoNaPraticaExercicio2 processes an integer slice, converts each integer to a string,
// and joins them into a single string with a comma and space separator, except for the last element.
// It then formats and prints the resulting string and its type using a custom formatting function.
func ResolucaoNaPraticaExercicio2() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var newSlice []string

	for _, value := range slice {
		if value != len(slice) {
			newSlice = append(newSlice, fmt.Sprintf("%v, ", value))
		} else {
			newSlice = append(newSlice, fmt.Sprintf("%v", value))
		}
	}

	joinSlice := strings.Join(newSlice, "")
	resolucao := fmt.Sprintf("\nSlice: %v\nTipo: %T", joinSlice, newSlice)

	format.FormatResolucaoExercicios(resolucao)
}

// ResolucaoNaPraticaExercicio3 processes a slice of integers and converts it into a formatted string.
// It iterates over the integer slice, converts each integer to a string, and appends it to a new slice of strings.
// If the integer is not the last element in the slice, it appends a comma and a space after the integer.
// After processing the slice, it joins the first four elements of the new string slice into a single string.
// Then, it trims any trailing commas and spaces from the joined string.
// Finally, it formats the result into a specific string format and passes it to the FormatResolucaoExercicios function for further handling.
func ResolucaoNaPraticaExercicio3() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var newSlice []string

	for _, value := range slice {
		if value != len(slice) {
			newSlice = append(newSlice, fmt.Sprintf("%v, ", value))
		} else {
			newSlice = append(newSlice, fmt.Sprintf("%v", value))
		}
	}
	joinSlice := strings.Join(newSlice[0:4], "")
	trimSlice := strings.Trim(joinSlice, ", ")

	resolucao := fmt.Sprintf("\nSlice: %v\nTipo: %T", trimSlice, newSlice)

	format.FormatResolucaoExercicios(resolucao)
}

// ResolucaoNaPraticaExercicio4 demonstrates the use of the append function in Go.
// It starts with a slice of integers and performs a series of append operations:
// 1. Appends the integer 52 to the original slice.
// 2. Appends the integers 53, 54, and 55 to the result of the first append.
// 3. Appends another slice of integers (56, 57, 58, 59, 60) to the result of the second append.
// The final result of each append operation is formatted into a string and passed to the
// FormatResolucaoExercicios function for further processing or display.
func ResolucaoNaPraticaExercicio4() {
	slice := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	append52 := append(slice, 52)
	append53to55 := append(append52, 53, 54, 55)

	sliceY := []int{56, 57, 58, 59, 60}
	appendSliceY := append(append53to55, sliceY...)

	resolucao := fmt.Sprintf("append52: %v\nappend53to55: %v\nappendSliceY: %v", append52, append53to55, appendSliceY)

	format.FormatResolucaoExercicios(resolucao)
}

// ResolucaoNaPraticaExercicio5 demonstrates how to manipulate slices in Go.
// It starts with a slice of integers from 42 to 51. Then, it creates a new slice
// by appending the first three elements of the original slice with the elements
// starting from the seventh element to the end of the original slice.
// Finally, it formats the new slice into a string and passes it to the
// FormatResolucaoExercicios function for further processing.
func ResolucaoNaPraticaExercicio5() {
	slice := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	novoSlice := append(slice[:3], slice[6:]...)

	resolucao := fmt.Sprintf("%v", novoSlice)

	format.FormatResolucaoExercicios(resolucao)
}

// ResolucaoNaPraticaExercicio6 demonstrates the creation and manipulation of a slice in Go.
// It initializes a slice of strings with a length of 26 to hold the names of Brazilian states.
// The `copy` function is used to populate the slice with the names of the states.
// The slice is then converted to a single string with the states separated by commas using `strings.Join`.
// Finally, the function formats a string containing the length, capacity, and the list of states,
// and passes this formatted string to the `FormatResolucaoExercicios` function for further processing or display.
func ResolucaoNaPraticaExercicio6() {
	estados := make([]string, 26)

	copy(estados, []string{
		"Acre", "Alagoas", "Amapá", "Amazonas", "Bahia", "Ceará", "Espírito Santo", "Goiás", "Maranhão",
		"Mato Grosso", "Mato Grosso do Sul", "Minas Gerais", "Pará", "Paraíba", "Paraná", "Pernambuco", "Piauí",
		"Rio de Janeiro", "Rio Grande do Norte", "Rio Grande do Sul", "Rondônia", "Roraima", "Santa Catarina",
		"São Paulo", "Sergipe", "Tocantins",
	})

	editEstados := strings.Join(estados, ", ")

	resolucao := fmt.Sprintf("len: %v cap: %v\nEstados: %v", len(estados), cap(estados), editEstados)

	format.FormatResolucaoExercicios(resolucao)
}

// ResolucaoNaPraticaExercicio7 prints the details of a list of people.
// It iterates over a slice of slices, where each inner slice contains
// the first name, last name, and favorite hobby of a person. The function
// formats these details into a string and passes it to the FormatResolucaoExercicios
// function for further processing or display.
func ResolucaoNaPraticaExercicio7() {
	var resolucao string

	pessoas := [][]string{
		{"Fabiano", "Florentino", "Programar"},
		{"Fulano", "de Tal", "Jogar bola"},
		{"Ciclano", "da Silva", "Assistir filmes"},
	}

	for _, pessoa := range pessoas {
		resolucao += fmt.Sprintf("Nome: %v, Sobrenome: %v, Hobby favorito: %v\n", pessoa[0], pessoa[1], pessoa[2])
	}

	format.FormatResolucaoExercicios(resolucao)
}

// ResolucaoNaPraticaExercicio8 generates a formatted string that lists people's names and their hobbies,
// then passes this string to the FormatResolucaoExercicios function for further processing or display.
// It creates a map where the keys are people's names and the values are slices of their hobbies.
// The function iterates over the map, constructing a formatted string that includes each person's name
// and their hobbies, and then calls FormatResolucaoExercicios with the resulting string.
func ResolucaoNaPraticaExercicio8() {
	var resolucao string

	pessoas := map[string][]string{
		"florentino_fabio": {"Programar", "Jogar bola", "Assistir filmes"},
		"de_tal_fulano":    {"Programar", "Jogar bola", "Assistir filmes"},
		"da_silva_ciclano": {"Programar 2", "Jogar bola 2", "Assistir filmes 2"},
	}

	for sobrenome_nome, hobbies := range pessoas {
		resolucao += fmt.Sprintf("Sobrenome_Nome: %v\n", sobrenome_nome)
		for idx, hobby := range hobbies {
			resolucao += fmt.Sprintf("  Hobby %v: %v\n", idx+1, hobby)
		}
	}

	format.FormatResolucaoExercicios(resolucao)
}

// ResolucaoNaPraticaExercicio9 is a function that creates a map of people and their hobbies,
// iterates over the map to format the data into a string, and then formats the string using
// the FormatResolucaoExercicios function from the format package. The map keys are strings
// representing the names of people, and the values are slices of strings representing their hobbies.
// The formatted string includes each person's name and their hobbies, with each hobby listed on a new line.
func ResolucaoNaPraticaExercicio9() {
	var resolucao string

	pessoas := map[string][]string{
		"florentino_fabiano": {"Programar", "Jogar bola", "Assistir filmes"},
		"de_tal_fulano":      {"Programar", "Jogar bola", "Assistir filmes"},
		"da_silva_ciclano":   {"Programar 2", "Jogar bola 2", "Assistir filmes 2"},
		"de_tal_ciclano":     {"Programar 3", "Jogar bola 3", "Assistir filmes 3"},
	}

	for sobrenome_nome, hobbies := range pessoas {
		resolucao += fmt.Sprintf("Sobrenome_Nome: %v\n", sobrenome_nome)
		for idx, hobby := range hobbies {
			resolucao += fmt.Sprintf("  Hobby %v: %v\n", idx+1, hobby)
		}
	}

	format.FormatResolucaoExercicios(resolucao)
}

// ResolucaoNaPraticaExercicio10 demonstrates the creation and manipulation of a map in Go.
// It initializes a map where the keys are strings representing names and the values are slices of strings representing hobbies.
// The function then removes an entry from the map and iterates over the map using a range loop to construct a formatted string.
// This string includes each key (name) and its associated hobbies, which are then passed to a formatting function for display.
func ResolucaoNaPraticaExercicio10() {
	var resolucao string

	pessoas := map[string][]string{
		"florentino_fabio": {"Programar", "Jogar bola", "Assistir filmes"},
		"de_tal_fulano":    {"Programar", "Jogar bola", "Assistir filmes"},
		"da_silva_ciclano": {"Programar 2", "Jogar bola 2", "Assistir filmes 2"},
		"de_tal_ciclano":   {"Programar 3", "Jogar bola 3", "Assistir filmes 3"},
	}

	// Removendo uma entrada do map
	delete(pessoas, "da_silva_ciclano")

	// Exibindo o map inteiro utilizando range loop
	// O range loop retorna a chave e o valor do map em cada iteração do loop
	// A chave é o sobrenome_nome e o valor é o slice de hobbies
	for sobrenome_nome, hobbies := range pessoas {
		resolucao += fmt.Sprintf("Sobrenome_Nome: %v\n", sobrenome_nome)
		for idx, hobby := range hobbies {
			resolucao += fmt.Sprintf("  Hobby %v: %v\n", idx+1, hobby)
		}
	}

	format.FormatResolucaoExercicios(resolucao)
}
