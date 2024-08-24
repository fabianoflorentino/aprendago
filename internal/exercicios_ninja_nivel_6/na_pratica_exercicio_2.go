package exercicios_ninja_nivel_6

import (
	"fmt"

	"github.com/fabianoflorentino/aprendago/pkg/format"
)

func NaPraticaExercicio2() {
	topic := format.OutlineContent{
		Title: "Na Prática - Exercício 2",
		Content: `
- Crie uma função que receba um parâmetro variádico do tipo int e retorne a soma de todos os ints recebidos;
- Passe um valor do tipo slice of int como argumento para a função;
- Crie outra função, esta deve receber um valor do tipo slice of int e retornar a soma de todos os elementos da slice;
- Passe um valor do tipo slice of int como argumento para a função.
- Solução: https://play.golang.org/p/Tgv3wwxKV-
    `,
	}

	format.FormatOutlineTopic(topic)
}

func ResolucaoNaPraticaExercicio2() {
	slice1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	result := SomaVariadico(slice1...)
	fmt.Printf("Soma do variádico: %v", result)

	slice2 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	result2 := SomaSlice(slice2)
	fmt.Printf("\nSoma do slice: %v", result2)
}

func SomaVariadico(x ...int) int {
	soma := 0

	for _, inteiro := range x {
		soma += inteiro
	}
	return soma
}

func SomaSlice(slice []int) int {
	soma := 0

	for _, inteiro := range slice {
		soma += inteiro
	}
	return soma
}
