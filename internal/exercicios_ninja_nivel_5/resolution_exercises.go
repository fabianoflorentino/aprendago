// Package exercicios_ninja_nivel_5 contains solutions for practical exercises
// in level 5 of the Ninja Go course. The exercises cover various Go language
// features such as structs, slices, maps, and anonymous structs.
package exercicios_ninja_nivel_5

import "fmt"

// ResolucaoNaPraticaExercicio1 demonstrates the creation and usage of a custom struct type in Go.
// The function defines a 'pessoa' struct with fields for first name, last name, and favorite ice cream flavors.
// It then creates two instances of 'pessoa', each with different values for these fields.
// These instances are stored in a slice of 'pessoa'.
// The function iterates over the slice and prints out the first name, last name, and favorite ice cream flavors of each person.
func ResolucaoNaPraticaExercicio1() {
	type pessoa struct {
		nome               string
		sobrenome          string
		sabores_de_sorvete []string
	}

	p1 := pessoa{
		nome:               "Fulano",
		sobrenome:          "de Tal",
		sabores_de_sorvete: []string{"Chocolate", "Morango", "Baunilha"},
	}

	p2 := pessoa{
		nome:               "Ciclano",
		sobrenome:          "da Silva",
		sabores_de_sorvete: []string{"Pistache", "Creme", "Coco"},
	}

	pessoas := []pessoa{p1, p2}

	for _, p := range pessoas {
		fmt.Println(p.nome, p.sobrenome)
		for _, s := range p.sabores_de_sorvete {
			fmt.Printf("\t%v\n", s)
		}
	}
}

// ResolucaoNaPraticaExercicio2 demonstrates the creation and usage of a custom struct type in Go.
// It defines a struct type `pessoa` with fields for first name, last name, and favorite ice cream flavors.
// Two instances of `pessoa` are created and stored in a map, with the last name as the key.
// The function then iterates over the map and prints each person's last name followed by their favorite ice cream flavors.
func ResolucaoNaPraticaExercicio2() {
	type pessoa struct {
		nome               string
		sobrenome          string
		sabores_de_sorvete []string
	}

	p1 := pessoa{
		nome:               "Fulano",
		sobrenome:          "de Tal",
		sabores_de_sorvete: []string{"Chocolate", "Morango", "Baunilha"},
	}

	p2 := pessoa{
		nome:               "Ciclano",
		sobrenome:          "da Silva",
		sabores_de_sorvete: []string{"Pistache", "Creme", "Coco"},
	}

	pessoas := map[string]pessoa{
		p1.sobrenome: p1,
		p2.sobrenome: p2,
	}

	for k, p := range pessoas {
		println(k)
		for _, s := range p.sabores_de_sorvete {
			println("\t", s)
		}
	}
}

// ResolucaoNaPraticaExercicio3 demonstrates the use of embedded structs in Go.
// It defines a base struct `veiculo` with fields `portas` (doors) and `cor` (color).
// Two other structs, `caminhonete` (pickup truck) and `sedan`, embed the `veiculo` struct
// and add their own fields: `tracaoNasQuatro` (four-wheel drive) for `caminhonete` and
// `modeloLuxo` (luxury model) for `sedan`.
// The function creates instances of `caminhonete` and `sedan`, initializes them with values,
// and prints their details using the `fmt.Printf` function. It also demonstrates accessing
// the embedded struct fields directly.
func ResolucaoNaPraticaExercicio3() {
	type veiculo struct {
		portas int
		cor    string
	}

	type caminhonete struct {
		veiculo
		tracaoNasQuatro bool
	}

	type sedan struct {
		veiculo
		modeloLuxo bool
	}

	cnt := caminhonete{
		veiculo: veiculo{
			portas: 4,
			cor:    "Preto",
		},

		tracaoNasQuatro: true,
	}

	sdn := sedan{
		veiculo: veiculo{
			portas: 2,
			cor:    "Branco",
		},

		modeloLuxo: true,
	}

	fmt.Printf("Caminhonete: %+v \n", cnt)
	fmt.Printf("Sedan: %+v \n", sdn)

	fmt.Println("Portas da caminhonete:", cnt.portas)
	fmt.Println("Cor do sedan:", sdn.cor)
}

// ResolucaoNaPraticaExercicio4 demonstrates the creation and initialization of an anonymous struct
// representing a person with fields for name, age, friends, and favorite foods. The function then
// prints the struct to the console. The struct includes:
// - Nome: a string representing the person's name
// - Idade: an integer representing the person's age
// - Amigos: a map where the keys are friends' names (strings) and the values are their ages (integers)
// - ComidasFavoritas: a slice of strings representing the person's favorite foods
// The function also contains commented-out code for converting the struct to JSON and printing it.
func ResolucaoNaPraticaExercicio4() {
	pessoa := struct {
		Nome             string
		Idade            int
		Amigos           map[string]int
		ComidasFavoritas []string
	}{
		Nome:  "Fabiano",
		Idade: 39,
		Amigos: map[string]int{
			"Ale":   46,
			"Lucas": 38,
		},
		ComidasFavoritas: []string{
			"Pizza",
			"Lasanha",
			"Hamburguer",
		},
	}

	// mais para frente
	// jsonData, _ := json.Marshal(pessoa)
	// fmt.Println(string(jsonData))

	fmt.Println(pessoa)
}
