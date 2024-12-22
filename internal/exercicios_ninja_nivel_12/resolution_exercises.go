package exercicios_ninja_nivel_12

import "fmt"

// ResolucaoNaPraticaExercicio1 calculates and prints the age of a dog in dog years
// based on a given human age. The age is hardcoded as 3 years in this example.
func ResolucaoNaPraticaExercicio1() {
	idade := 3
	idadeCachorro := idadeDeCachorro(idade)

	fmt.Printf("A idade do cachorro em anos humanos é: %d", idadeCachorro)
}

func ResolucaoNaPraticaExercicio2() {
	fmt.Printf("Não há resolução para este exercício.")
}

func ResolucaoNaPraticaExercicio3() {
	fmt.Printf("Não há resolução para este exercício.")
}

// idadeDeCachorro converts a dog's age from human years to dog years.
// It takes an integer representing the dog's age in human years and returns
// the equivalent age in dog years, calculated by multiplying the input by 7.
func idadeDeCachorro(idade int) int {
	return idade * 7
}
