// Package exercicios_ninja_nivel_7 contains solutions for practical exercises
// at level 7 of the Ninja Go course. It includes functions demonstrating
// basic Go concepts such as pointers, structs, and function calls.
package exercicios_ninja_nivel_7

import "fmt"

// pessoa represents a person with a first name, last name, and age.
// Fields:
// - nome: the first name of the person
// - sobrenome: the last name of the person
// - idade: the age of the person
type pessoa struct {
	nome      string
	sobrenome string
	idade     int
}

// ResolucaoNaPraticaExercicio1 prints the memory address of a string variable.
// It initializes a string variable 'valor' with the value "valor" and then
// prints the memory address of this variable using the built-in println function.
func ResolucaoNaPraticaExercicio1() {
	valor := "valor"
	println(&valor)
}

// ResolucaoNaPraticaExercicio2 demonstrates the creation and modification of a 'pessoa' struct instance.
// It initializes a 'pessoa' struct with predefined values, prints the original values,
// then calls the 'mudeMe' function to modify the struct via a pointer, and finally prints the modified values.
func ResolucaoNaPraticaExercicio2() {
	p := pessoa{
		nome:      "Fulano",
		sobrenome: "de Tal",
		idade:     25,
	}

	fmt.Printf("Original -> Nome: %s %s, Idade: %d\n", p.nome, p.sobrenome, p.idade)

	mudeMe(&p)

	fmt.Printf("Alterado -> Nome: %s %s, Idade: %d\n", p.nome, p.sobrenome, p.idade)
}

// mudeMe modifies the fields of the given pessoa struct.
// It sets the nome field to "João", the sobrenome field to "Silva",
// and the idade field to 30.
//
// Parameters:
//
//	p (*pessoa): A pointer to a pessoa struct that will be modified.
func mudeMe(p *pessoa) {
	p.nome = "João"
	p.sobrenome = "Silva"
	p.idade = 30
}
