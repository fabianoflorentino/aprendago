package aplicacoes

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
)

// pessoa is a struct that represents a person.
type pessoa struct {
	Name     string
	Lastname string
	Age      int
}

// extractInfo is a struct that represents a json object with fields from a json file.
type extractInfo struct {
	Name     string `json:"Name"`
	Lastname string `json:"Lastname"`
	Age      int    `json:"Age"`
}

// carro is a struct that represents a car.
type carro struct {
	Modelo   string
	Ano      int
	Potencia int
}

// UsingJsonMarshal is a function that demonstrates how to use the json package to marshal a struct.
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

// UsingJsonUnmarshal is a function that demonstrates how to use the json package to unmarshal a struct.
func UsingJsonUnmarshal() {
	pessoaJSON := []byte(`{"Name":"Fabiano","Lastname":"Florentino","Age":40}`)

	var p extractInfo
	err := json.Unmarshal(pessoaJSON, &p)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Name: %s\nLastname: %s\nAge: %d\n", p.Name, p.Lastname, p.Age)
}

// UsingJsonEncoder is a function that demonstrates how to use the json package to encode a struct.
func UsingJsonEncoder() {
	p := pessoa{
		Name:     "Fabiano",
		Lastname: "Florentino",
		Age:      40,
	}

	json.NewEncoder(os.Stdout).Encode(p)
}

// UsingPackageSort is a function that demonstrates how to use the sort package to sort a slice of strings and ints.
// The sort package is a package that provides primitives for sorting slices and user-defined collections.
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

// usingSortSrtrings is a function that demonstrates how to use the sort package to sort a slice of strings.
func usingSortSrtrings(example []string) {
	sort.Strings(example)
}

// UsingSortInts is a function that demonstrates how to use the sort package to sort a slice of ints.
func usingSortInts(example []int) {
	sort.Ints(example)
}

// UsingCustomSort is a function that demonstrates how to use the sort package to sort a slice of structs.
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

type ByModel []carro
type ByYear []carro
type ByPower []carro

// ByModel implements the sort.Interface for []carro based on the Modelo field.
func (name ByModel) Len() int { return len(name) }
func (name ByModel) Less(index_i, index_j int) bool {
	return name[index_i].Modelo < name[index_j].Modelo
}
func (name ByModel) Swap(index_i, index_j int) {
	name[index_i], name[index_j] = name[index_j], name[index_i]
}

// ByYear implements the sort.Interface for []carro based on the Ano field.
func (year ByYear) Len() int { return len(year) }
func (year ByYear) Less(index_i, index_j int) bool {
	return year[index_i].Ano < year[index_j].Ano
}
func (year ByYear) Swap(index_i, index_j int) {
	year[index_i], year[index_j] = year[index_j], year[index_i]
}

// ByPower implements the sort.Interface for []carro based on the Potencia field.
func (power ByPower) Len() int { return len(power) }
func (power ByPower) Less(index_i, index_j int) bool {
	return power[index_i].Potencia < power[index_j].Potencia
}
func (power ByPower) Swap(index_i, index_j int) {
	power[index_i], power[index_j] = power[index_j], power[index_i]
}
