package exercicios_ninja_nivel_3

import (
	"fmt"
	"time"
)

func ResolucaoNaPraticaExercicio1() {
	for i := 1; i <= 60; i++ {
		fmt.Printf("%d ", i)
	}
}

func ResolucaoNaPraticaExercicio2() {
	for upperLetter := 65; upperLetter <= 90; upperLetter++ {
		fmt.Printf("%d: ", upperLetter)

		for threeTimes := 0; threeTimes < 3; threeTimes++ {
			fmt.Printf(" %#U '%c'", threeTimes, upperLetter)
		}
		fmt.Printf("\n")
	}
}

func ResolucaoNaPraticaExercicio3() {
	var birdthYear int = 1985
	currentYear := time.Now().Year()

	// fmt.Printf("\nDigite o ano de seu nascimento: ")
	// fmt.Scan(&birdthYear)

	for year := birdthYear; year <= currentYear; year++ {
		if year == currentYear {
			fmt.Printf("%d\n", year)
		} else {
			fmt.Printf("%d, ", year)
		}
	}

	fmt.Printf("Sua idade: %v", currentYear-birdthYear)
}

// ResolucaoNaPraticaExercicio4 - Exercício #4
func ResolucaoNaPraticaExercicio4() {
	var birdthYear int = 1985
	currentYear := time.Now().Year()

	// Nesse formato de loop, é importante quebrar o loop com o break
	for {
		if birdthYear >= currentYear {
			fmt.Printf("%d ", birdthYear)
			// nessário para quebrar o loop
			break
		} else {
			fmt.Printf("%d, ", birdthYear)
			birdthYear++
		}
	}

	fmt.Printf("\nSua idade: %v", time.Now().Year()-1985)
}

func ResolucaoNaPraticaExercicio5() {
	fmt.Printf("O resto da divisão por 4 de todos os números entre 10 e 100\n\n")
	for i := 10; i <= 100; i++ {
		if i == 100 {
			fmt.Printf("%d: %d ", i, i%4)
		} else {
			fmt.Printf("%d: %d, ", i, i%4)
		}
	}
}

func ResolucaoNaPraticaExercicio6() {
	fmt.Printf("10 é maior que 5? ")
	if 10 > 5 {
		fmt.Print("Sim")
	}
}

func ResolucaoNaPraticaExercicio7() {
	fmt.Printf("10 é maior que 5? ")

	if 10 > 5 {
		fmt.Print("Sim")
	} else if 5 > 10 {
		fmt.Print("Não")
	} else {
		fmt.Print("Não sei")
	}
}

func ResolucaoNaPraticaExercicio8() {
	statement := true

	switch {
	case statement:
		fmt.Println("Não deve ser impresso")
	case !statement:
		fmt.Println("Deve ser impresso")
	}
}

func ResolucaoNaPraticaExercicio9() {
	var esporteFavorito string = "futebol americano"

	switch esporteFavorito {
	case "futebol":
		fmt.Println("Seu esporte favorito é futebol")
	case "futebol americano":
		fmt.Println("Seu esporte favorito é futebol americano")
	case "basquete":
		fmt.Println("Seu esporte favorito é basquete")
	default:
		fmt.Println("Seu esporte favorito não é futebol, futebol americano ou basquete")
	}
}

func ResolucaoNaPraticaExercicio10() {
	anotacao := `
fmt.Println(true && true)
fmt.Println(true && false)
fmt.Println(true || true)
fmt.Println(true || false)
fmt.Println(!true)
`
	fmt.Printf("%v\n", anotacao)

	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(true || true)
	fmt.Println(true || false)
	fmt.Println(!true)

}
