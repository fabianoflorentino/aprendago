// Package funcoes provides utility functions for performing various operations
// such as summing numbers with specific conditions. This package includes
// functions that demonstrate the use of higher-order functions and variadic
// parameters in Go.
package funcoes

import (
	"fmt"
)

// ResolucaoDesafioCallback calculates the sum of odd numbers from a given slice of integers
// using the somaImpares function and prints the result along with an explanation.
// The somaImpares function takes a function and a variadic number of integers, and returns
// the sum of the odd numbers. The provided function sums the numbers if they are odd.
func ResolucaoDesafioCallback() {
	result := somaImapres(soma, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}...)

	explicacao := `
A função somaImpares() recebe uma função e um slice de inteiros e retorna a soma dos números pares.
A função passada recebe um número variatico de inteiros e retorna um inteiro.
Os números recebidos pela função são variaticos e serão passados para a função que soma os números somente se forem ímpares.
    `

	fmt.Printf("Soma dos números ímpares é: %d\n", result)
	fmt.Println(explicacao)
}

// soma takes a variable number of integer arguments and returns their sum.
// It iterates over the provided integers, adding each one to a cumulative total,
// which is then returned as the result.
func soma(numeros ...int) int {
	somaNum := 0

	for _, num := range numeros {
		somaNum += num
	}

	return somaNum
}

// somaImapres takes a function `f` and a variadic number of integers `numeros`.
// It sums up the results of applying `f` to each odd number in `numeros`.
// The function `f` should accept a variadic number of integers and return an integer.
// The function returns the total sum of the results.
func somaImapres(f func(...int) int, numeros ...int) int {
	soma := 0

	for _, num := range numeros {
		if num%2 != 0 {
			soma += f(num)
		}
	}
	return soma
}
