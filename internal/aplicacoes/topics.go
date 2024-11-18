package aplicacoes

import (
	"fmt"

	"github.com/fabianoflorentino/aprendago/pkg/format"
)

const rootDir = "internal/aplicacoes"

func Aplicacoes() {
	executeSections("Documentação JSON")
	executeSections("JSON marshal (ordenação)")
	executeSections("JSON unmarshal (desornação)")
	executeSections("A interface Writer")
	executeSections("O pacote sort")
	executeSections("Customizando sort")
	executeSections("bcrypt")

}

func MenuAplicacoes([]string) []format.MenuOptions {

	return []format.MenuOptions{
		{Options: "--documentacao-json", ExecFunc: func() { executeSections("Documentação JSON") }},
		{Options: "--json-marshal", ExecFunc: func() { executeSections("JSON marshal (ordenação)") }},
		{Options: "--json-unmarshal", ExecFunc: func() { executeSections("JSON unmarshal (desornação)") }},
		{Options: "--interface-writer", ExecFunc: func() { executeSections("A interface Writer") }},
		{Options: "--pacote-sort", ExecFunc: func() { executeSections("O pacote sort") }},
		{Options: "--customizando-sort", ExecFunc: func() { executeSections("Customizando sort") }},
		{Options: "--bcrypt", ExecFunc: func() { executeSections("bcrypt") }},
	}
}

func HelpMeAplicacoes() {
	hlp := []format.HelpMe{
		{Flag: "--documentacao-json", Description: "Descreve como documentar um pacote em Go", Width: 0},
		{Flag: "--json-marshal", Description: "Descreve como ordenar um JSON", Width: 0},
		{Flag: "--json-unmarshal", Description: "Descreve como desordenar um JSON", Width: 0},
		{Flag: "--interface-writer", Description: "Descreve o que é a interface Writer", Width: 0},
		{Flag: "--pacote-sort", Description: "Descreve o pacote sort", Width: 0},
		{Flag: "--customizando-sort", Description: "Descreve como customizar o pacote sort", Width: 0},
		{Flag: "--bcrypt", Description: "Descreve o pacote bcrypt", Width: 0},
	}

	fmt.Printf("\nCapítulo 16: Aplicações\n")
	format.PrintHelpMe(hlp)
}

func executeSections(section string) {
	format.FormatSection(rootDir, section)
}
