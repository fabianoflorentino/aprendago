package exercicios_ninja_nivel_4

import (
	"fmt"
	"strings"

	"github.com/fabianoflorentino/aprendago/pkg/format"
)

func ResolucaoNaPraticaExercicio1() {
	array := [5]int{1, 2, 3, 4, 5}
	var resolucao string

	for _, value := range array {
		resolucao += fmt.Sprintf("%v, ", value)
	}

	resolucao = strings.Trim(resolucao, ", ")
	format.FormatResolucaoExercicios(resolucao)
}

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

func ResolucaoNaPraticaExercicio4() {
	slice := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	append52 := append(slice, 52)
	append53to55 := append(append52, 53, 54, 55)

	sliceY := []int{56, 57, 58, 59, 60}
	appendSliceY := append(append53to55, sliceY...)

	resolucao := fmt.Sprintf("append52: %v\nappend53to55: %v\nappendSliceY: %v", append52, append53to55, appendSliceY)

	format.FormatResolucaoExercicios(resolucao)
}

func ResolucaoNaPraticaExercicio5() {
	slice := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	novoSlice := append(slice[:3], slice[6:]...)

	resolucao := fmt.Sprintf("%v", novoSlice)

	format.FormatResolucaoExercicios(resolucao)
}

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

func ResolucaoNaPraticaExercicio7() {
	var resolucao string

	pessoas := [][]string{
		{"Fábio", "Florentino", "Programar"},
		{"Fulano", "de Tal", "Jogar bola"},
		{"Ciclano", "da Silva", "Assistir filmes"},
	}

	for _, pessoa := range pessoas {
		resolucao += fmt.Sprintf("Nome: %v, Sobrenome: %v, Hobby favorito: %v\n", pessoa[0], pessoa[1], pessoa[2])
	}

	format.FormatResolucaoExercicios(resolucao)
}

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

func ResolucaoNaPraticaExercicio9() {
	var resolucao string

	pessoas := map[string][]string{
		"florentino_fabio": {"Programar", "Jogar bola", "Assistir filmes"},
		"de_tal_fulano":    {"Programar", "Jogar bola", "Assistir filmes"},
		"da_silva_ciclano": {"Programar 2", "Jogar bola 2", "Assistir filmes 2"},
		"de_tal_ciclano":   {"Programar 3", "Jogar bola 3", "Assistir filmes 3"},
	}

	for sobrenome_nome, hobbies := range pessoas {
		resolucao += fmt.Sprintf("Sobrenome_Nome: %v\n", sobrenome_nome)
		for idx, hobby := range hobbies {
			resolucao += fmt.Sprintf("  Hobby %v: %v\n", idx+1, hobby)
		}
	}

	format.FormatResolucaoExercicios(resolucao)
}

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
