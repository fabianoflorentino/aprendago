package aplicacoes

import (
	"fmt"

	"github.com/fabianoflorentino/aprendago/pkg/format"
)

// HelpMeAplicacoes provides a list of help topics related to various Go packages and functionalities.
// It prints out a formatted help guide for each topic, including flags and descriptions.
// The topics cover documentation, JSON handling, interfaces, sorting, and encryption with bcrypt.
func Help() {
	hlp := []format.HelpMe{
		{Flag: "--documentacao-json", Description: "Descreve como documentar um pacote em Go", Width: 0},
		{Flag: "--documentacao-json --example --json-marshal", Description: "Exemplo de como ordenar um JSON", Width: 0},
		{Flag: "--documentacao-json --example --json-unmarshal", Description: "Exemplo de como desordenar um JSON", Width: 0},
		{Flag: "--documentacao-json --example --json-encoder", Description: "Exemplo de como usar o encoder JSON", Width: 0},
		{Flag: "--json-marshal", Description: "Descreve o pacote json.Marshal", Width: 0},
		{Flag: "--json-unmarshal", Description: "Descreve o pacote json.Unmarshal", Width: 0},
		{Flag: "--interface-writer", Description: "Descreve o que é a interface Writer", Width: 0},
		{Flag: "--pacote-sort", Description: "Descreve o pacote sort", Width: 0},
		{Flag: "--pacote-sort --example", Description: "Exemplo de como usar o pacote sort", Width: 0},
		{Flag: "--customizando-sort --example", Description: "Descreve como customizar o pacote sort", Width: 0},
		{Flag: "--bcrypt", Description: "Descreve o pacote bcrypt", Width: 0},
	}

	fmt.Printf("\nCapítulo 16: Aplicações\n")
	format.PrintHelpMe(hlp)
}
