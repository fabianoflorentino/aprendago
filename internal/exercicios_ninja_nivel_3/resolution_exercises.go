// Package exercicios_ninja_nivel_3 contains a series of functions that solve various exercises
// demonstrating basic Go programming concepts such as loops, conditionals, and switch statements.
package exercicios_ninja_nivel_3

import (
	"fmt"
	"time"
)

// ResolucaoNaPraticaExercicio1 prints the numbers from 1 to 60, each followed by a space.
// It uses a for loop to iterate through the numbers and fmt.Printf to print each number.
func ResolucaoNaPraticaExercicio1() {
	for i := 1; i <= 60; i++ {
		fmt.Printf("%d ", i)
	}
}

// ResolucaoNaPraticaExercicio2 prints the ASCII values and characters of uppercase letters (A-Z).
// For each uppercase letter, it prints the ASCII value followed by the character three times.
// The outer loop iterates through the ASCII values of uppercase letters (65 to 90).
// The inner loop prints each character three times in the format " %#U '%c'".
func ResolucaoNaPraticaExercicio2() {
	for upperLetter := 65; upperLetter <= 90; upperLetter++ {
		fmt.Printf("%d: ", upperLetter)

		for threeTimes := 0; threeTimes < 3; threeTimes++ {
			fmt.Printf(" %#U '%c'", threeTimes, upperLetter)
		}
		fmt.Printf("\n")
	}
}

// ResolucaoNaPraticaExercicio3 prints all the years from the user's birth year up to the current year,
// and then prints the user's age. The birth year is hardcoded as 1985, and the current year is obtained
// using the time.Now().Year() function. The function iterates through each year from the birth year to
// the current year, printing each year followed by a comma, except for the current year which is printed
// without a comma. Finally, it calculates and prints the user's age based on the difference between the
// current year and the birth year.
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

// ResolucaoNaPraticaExercicio4 prints each year from the given birth year (1985) up to the current year.
// It uses an infinite loop that breaks when the birth year is greater than or equal to the current year.
// For each year, it prints the year followed by a comma, except for the last year which is printed without a comma.
// Finally, it prints the calculated age based on the current year and the birth year.
func ResolucaoNaPraticaExercicio4() {
	var birdthYear int = 1985
	currentYear := time.Now().Year()

	for {
		if birdthYear >= currentYear {
			fmt.Printf("%d ", birdthYear)
			// nessário para quebrar o loop
			break
		} else {
			breakLineby10(birdthYear)

			fmt.Printf("%d, ", birdthYear)
			birdthYear++
		}
	}

	fmt.Printf("\n\nSua idade: %v", time.Now().Year()-1985)
}

var breakLineby10 = func(i int) error {
	if i%10 == 0 {
		fmt.Printf("\n")
	}

	return nil
}

// ResolucaoNaPraticaExercicio5 prints the remainder of the division by 4
// for all numbers between 10 and 100. It iterates through each number in
// this range, calculates the remainder when the number is divided by 4,
// and prints the result in a formatted string. The output is a comma-separated
// list of numbers and their remainders, except for the last number which is
// followed by a space instead of a comma.
func ResolucaoNaPraticaExercicio5() {
	fmt.Printf("O resto da divisão por 4 de todos os números entre 10 e 100\n\n")
	for i := 10; i <= 100; i++ {
		if i%10 == 0 && i != 10 {
			fmt.Printf("\n")
		}
		if i == 100 {
			fmt.Printf("%d: %d ", i, i%4)
		} else {
			fmt.Printf("%d: %d, ", i, i%4)
		}
	}
}

// ResolucaoNaPraticaExercicio6 prints a message asking if 10 is greater than 5,
// and then prints "Sim" (which means "Yes" in Portuguese) if the condition is true.
func ResolucaoNaPraticaExercicio6() {
	fmt.Printf("10 é maior que 5? ")
	if 10 > 5 {
		fmt.Print("Sim")
	}
}

// ResolucaoNaPraticaExercicio7 prints whether 10 is greater than 5.
// It first prints the question "10 é maior que 5? ".
// Then, it evaluates the condition if 10 is greater than 5.
// If true, it prints "Sim".
// If 5 is greater than 10, it prints "Não".
// Otherwise, it prints "Não sei".
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

// ResolucaoNaPraticaExercicio8 demonstrates a switch statement with boolean conditions.
// The function initializes a boolean variable `statement` to true and uses a switch statement
// to print a message based on the value of `statement`. Since `statement` is true, the first
// case will match, and "Não deve ser impresso" will be printed. The second case will not be
// executed because `statement` is not false.
func ResolucaoNaPraticaExercicio8() {
	statement := true

	switch {
	case statement:
		fmt.Println("Não deve ser impresso")
	case !statement:
		fmt.Println("Deve ser impresso")
	}
}

// ResolucaoNaPraticaExercicio9 demonstrates a simple switch statement in Go.
// It initializes a variable `esporteFavorito` with the value "futebol americano"
// and uses a switch statement to print a message based on the value of `esporteFavorito`.
// If `esporteFavorito` is "futebol", it prints "Seu esporte favorito é futebol".
// If `esporteFavorito` is "futebol americano", it prints "Seu esporte favorito é futebol americano".
// If `esporteFavorito` is "basquete", it prints "Seu esporte favorito é basquete".
// For any other value of `esporteFavorito`, it prints "Seu esporte favorito não é futebol, futebol americano ou basquete".
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

// ResolucaoNaPraticaExercicio10 demonstrates the evaluation of basic boolean expressions in Go.
// It prints a multi-line string containing the boolean expressions, followed by the actual evaluation of these expressions.
// The expressions include:
// - true && true
// - true && false
// - true || true
// - true || false
// - !true
// The function helps in understanding how logical AND (&&), OR (||), and NOT (!) operators work in Go.
func ResolucaoNaPraticaExercicio10() {
	anotacao := `
	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(true || true)
	fmt.Println(true || false)
	fmt.Println(!true)
	`
	fmt.Printf("%v\n", anotacao)

	fmt.Println(true)
	fmt.Println(true && false)
	fmt.Println(true)
	fmt.Println(true || false)
	fmt.Println(!true)
}
