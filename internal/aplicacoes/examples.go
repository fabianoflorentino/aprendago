package aplicacoes

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
)

// pessoa represents an individual with a name, lastname, and age.
// It is used to store personal information about a person.
type pessoa struct {
	Name     string
	Lastname string
	Age      int
}

// extractInfo represents a structure to hold personal information.
// It includes the following fields:
// - Name: The first name of the individual.
// - Lastname: The last name of the individual.
// - Age: The age of the individual.
type extractInfo struct {
	Name     string `json:"Name"`
	Lastname string `json:"Lastname"`
	Age      int    `json:"Age"`
}

// carro represents a car with its model, year, and power.
// Modelo: The model of the car as a string.
// Ano: The manufacturing year of the car as an integer.
// Potencia: The power of the car's engine as an integer.
type carro struct {
	Modelo   string
	Ano      int
	Potencia int
}

// UsingJsonMarshal demonstrates how to convert a struct to a JSON string using the json.Marshal function.
// It creates an instance of the pessoa struct, populates it with sample data, and then marshals it into a JSON byte slice.
// If the marshaling process encounters an error, the function will panic.
// Finally, it writes the resulting JSON to the standard output.
func UsingJsonMarshal() {
	p := pessoa{
		Name:     "Fabiano",
		Lastname: "Florentino",
		Age:      40,
	}

	pessoaToJSON, err := json.Marshal(p)
	if err != nil {
		panic(err)
	}

	os.Stdout.Write(pessoaToJSON)
}

// UsingJsonUnmarshal demonstrates how to unmarshal a JSON-encoded byte slice into a Go struct.
// It defines a JSON string representing a person with fields Name, Lastname, and Age.
// The JSON string is unmarshaled into an instance of the extractInfo struct.
// If the unmarshaling process encounters an error, the function will panic.
// Finally, it prints the extracted information (Name, Lastname, and Age) to the console.
func UsingJsonUnmarshal() {
	pessoaJSON := []byte(`{"Name":"Fabiano","Lastname":"Florentino","Age":40}`)

	var p extractInfo
	err := json.Unmarshal(pessoaJSON, &p)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Name: %s\nLastname: %s\nAge: %d\n", p.Name, p.Lastname, p.Age)
}

// UsingJsonEncoder demonstrates how to encode a struct into JSON format
// and write it to the standard output using the json.NewEncoder method.
// It creates an instance of the 'pessoa' struct with sample data and
// encodes it to JSON, printing the result to the console.
func UsingJsonEncoder() {
	p := pessoa{
		Name:     "Fabiano",
		Lastname: "Florentino",
		Age:      40,
	}

	json.NewEncoder(os.Stdout).Encode(p)
}

// UsingPackageSort demonstrates the usage of the sort package to sort slices of strings and integers.
// It initializes two slices, one with strings and another with integers, and prints them unsorted.
// Then, it sorts the slices using custom sorting functions (usingSortStrings and usingSortInts) and prints the sorted slices.
func UsingPackageSort() {
	example_of_strings := []string{"Fabiano", "Florentino", "Aprendendo", "Go"}
	example_of_ints := []int{0, 9, 5, 3, 1, 7, 8, 2, 6, 4}

	fmt.Printf("Unsorted Strings: %v\n", example_of_strings)
	fmt.Printf("Unsorted Ints:    %v\n", example_of_ints)

	usingSortSrtrings(example_of_strings)
	fmt.Printf("\nSorted Strings: %v\n", example_of_strings)

	usingSortInts(example_of_ints)
	fmt.Printf("Sorted Ints:    %v\n", example_of_ints)
}

// usingSortStrings sorts a slice of strings in ascending order.
// It takes a slice of strings as input and modifies it in place.
// The function uses the sort.Strings function from the sort package
// to perform the sorting.
func usingSortSrtrings(example []string) {
	sort.Strings(example)
}

// usingSortInts sorts a slice of integers in ascending order.
// It takes a slice of integers as input and modifies it in place
// to arrange the elements in increasing order using the sort.Ints function.
func usingSortInts(example []int) {
	sort.Ints(example)
}

// UsingCustomSort demonstrates how to sort a slice of 'carro' structs
// using custom sorting criteria. It sorts the cars by model, year, and power.
// The function first prints the unsorted list of cars, then sorts and prints
// the list by model, year, and power respectively.
func UsingCustomSort() {
	carros := []carro{
		{"Fusca", 1978, 50},
		{"Brasilia", 1975, 60},
		{"Chevette", 1980, 70},
		{"Corcel", 1979, 80},
	}

	fmt.Printf("Unsorted Cars: %v\n", carros)

	sort.Sort(ByModel(carros))
	fmt.Printf("\nSorted by Model: %v\n", carros)

	sort.Sort(ByYear(carros))
	fmt.Printf("Sorted by Year: %v\n", carros)

	sort.Sort(ByPower(carros))
	fmt.Printf("Sorted by Power: %v\n", carros)
}

// ByModel is a type alias for a slice of carro structs.
// It is used to implement custom sorting based on the model of the car.
type ByModel []carro

// ByYear is a type alias for a slice of carro structs.
// It is used to sort a collection of carro instances by their year of manufacture.
type ByYear []carro

// ByPower is a type alias for a slice of carro structs.
// It is used to implement custom sorting based on the power of the cars.
type ByPower []carro

// Len returns the number of elements in the ByModel collection.
// It implements the Len method of the sort.Interface.
func (name ByModel) Len() int { return len(name) }

// Less compares the Modelo field of two elements in the ByModel slice.
// It returns true if the Modelo of the element at index_i is less than
// the Modelo of the element at index_j, and false otherwise.
// This function is used to sort a slice of elements by their Modelo field.
func (name ByModel) Less(index_i, index_j int) bool {
	return name[index_i].Modelo < name[index_j].Modelo
}

// Swap exchanges the elements at the specified indices in the ByModel slice.
// It takes two integer parameters, index_i and index_j, which represent the positions
// of the elements to be swapped.
func (name ByModel) Swap(index_i, index_j int) {
	name[index_i], name[index_j] = name[index_j], name[index_i]
}

// Len returns the number of elements in the ByYear collection.
// It implements the sort.Interface Len method.
func (year ByYear) Len() int { return len(year) }

// Less compares the Ano field of two elements in the ByYear slice.
// It returns true if the Ano of the element at index_i is less than
// the Ano of the element at index_j, and false otherwise.
func (year ByYear) Less(index_i, index_j int) bool {
	return year[index_i].Ano < year[index_j].Ano
}

// Swap exchanges the elements at the specified indices in the ByYear slice.
// It takes two integer parameters, index_i and index_j, which represent the
// positions of the elements to be swapped.
func (year ByYear) Swap(index_i, index_j int) {
	year[index_i], year[index_j] = year[index_j], year[index_i]
}

// Len returns the number of elements in the ByPower slice.
// It implements the Len method of the sort.Interface.
func (power ByPower) Len() int { return len(power) }

// Less compares the Potencia field of two elements in the ByPower slice.
// It returns true if the Potencia of the element at index_i is less than
// the Potencia of the element at index_j, and false otherwise.
func (power ByPower) Less(index_i, index_j int) bool {
	return power[index_i].Potencia < power[index_j].Potencia
}

// Swap exchanges the elements at the specified indices in the ByPower slice.
// It takes two integer indices, index_i and index_j, and swaps the elements
// at these positions within the slice.
func (power ByPower) Swap(index_i, index_j int) {
	power[index_i], power[index_j] = power[index_j], power[index_i]
}
