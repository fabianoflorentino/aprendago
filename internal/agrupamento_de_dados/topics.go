package agrupamento_de_dados

import (
	"fmt"

	"github.com/fabianoflorentino/aprendago/pkg/format"
)

const rootDir = "internal/agrupamento_de_dados"

func AgrupamentoDeDados() {
	fmt.Printf("\n\n08 - Agrupamento de Dados\n")

	executeSection("Array")
	executeSection("Slice: literal composta")
	executeSection("Slice: for range")
	executeSection("Slice: fatiando ou deletando de uma fatia")
	executeSection("Slice: anexando a uma slice")
	executeSection("Slice: make")
	executeSection("Slice: multi dimensional")
	executeSection("Slice: a surpresa do array subjacente")
	executeSection("Maps: introdução")
	executeSection("Maps: range e deletando")
}

func MenuAgrupamentoDeDados([]string) []format.MenuOptions {
	return []format.MenuOptions{
		{Options: "--array", ExecFunc: func() { executeSection("Array") }},
		{Options: "--slice-literal-composta", ExecFunc: func() { executeSection("Slice: literal composta") }},
		{Options: "--slice-for-range", ExecFunc: func() { executeSection("Slice: for range") }},
		{Options: "--slice-fatiando-ou-deletando-de-uma-fatia", ExecFunc: func() { executeSection("Slice: fatiando ou deletando de uma fatia") }},
		{Options: "--slice-anexando-a-uma-slice", ExecFunc: func() { executeSection("Slice: anexando a uma slice") }},
		{Options: "--slice-make", ExecFunc: func() { executeSection("Slice: make") }},
		{Options: "--slice-multi-dimensional", ExecFunc: func() { executeSection("Slice: multi dimensional") }},
		{Options: "--slice-a-surpresa-do-array-subjacente", ExecFunc: func() { executeSection("Slice: a surpresa do array subjacente") }},
		{Options: "--maps-introducao", ExecFunc: func() { executeSection("Maps: introdução") }},
		{Options: "--maps-range-e-deletando", ExecFunc: func() { executeSection("Maps: range e deletando") }},
	}
}

func HelpMeAgrupamentoDeDados() {
	hlp := []format.HelpMe{
		{Flag: "--array", Description: "Apresenta o tópico Array.", Width: 0},
		{Flag: "--slice-literal-composta", Description: "Apresenta o tópico Slice Literal Composta.", Width: 0},
		{Flag: "--slice-for-range", Description: "Apresenta o tópico Slice: for range.", Width: 0},
		{Flag: "--slice-fatiando-ou-deletando-de-uma-fatia", Description: "Apresenta o tópico Slice: fatiando ou deletando de uma fatia.", Width: 0},
		{Flag: "--slice-fatiando-ou-deletando-de-uma-fatia --resolucao", Description: "Apresenta a resolução do tópico Slice: fatiando ou deletando de uma fatia.", Width: 0},
		{Flag: "--slice-anexando-a-uma-slice", Description: "Apresenta o tópico Slice: anexando a uma slice.", Width: 0},
		{Flag: "--slice-make", Description: "Apresenta o tópico Slice: Make.", Width: 0},
		{Flag: "--slice-multi-dimensional", Description: "Apresenta o tópico Slice: Multi Dimensional.", Width: 0},
		{Flag: "--slice-a-surpresa-do-array-subjacente", Description: "Apresenta o tópico Slice: a surpresa do array subjacente.", Width: 0},
		{Flag: "--maps-introducao", Description: "Apresenta o tópico Maps: introdução.", Width: 0},
		{Flag: "--maps-range-e-deletando", Description: "Apresenta o tópico Maps: Range e Deletando.", Width: 0},
	}

	fmt.Println("\nCapítulo 8: Agrupamento de Dados")
	format.PrintHelpMe(hlp)
}

func executeSection(section string) {
	format.FormatSection(rootDir, section)
}
