package variaveis_valores_tipos

import (
	"fmt"

	"github.com/fabianoflorentino/aprendago/pkg/format"
)

const rootDir = "internal/variaveis_valores_tipos"

func VariaveisValoresETipos() {
	fmt.Printf("02 - Variáveis, Valores & Tipos\n\n")

	executeSection("Go Playground")
	executeSection("Hello, World!")
	executeSection("Operador curto de declaração")
	executeSection("A palavra reservada var")
	executeSection("Explorando tipos")
	executeSection("Valor zero")
	executeSection("O pacote fmt")
	executeSection("Criando seu próprio tipo")
	executeSection("Conversão, não coerção")
}

func MenuVariaveisValoresTipos([]string) []format.MenuOptions {
	return []format.MenuOptions{
		{Options: "--go-playground", ExecFunc: func() { executeSection("Go Playground") }},
		{Options: "--hello-world", ExecFunc: func() { executeSection("Hello, World!") }},
		{Options: "--operador-curto-de-declaracao", ExecFunc: func() { executeSection("Operador curto de declaração") }},
		{Options: "--a-palavra-reservada-var", ExecFunc: func() { executeSection("A palavra reservada var") }},
		{Options: "--explorando-tipos", ExecFunc: func() { executeSection("Explorando tipos") }},
		{Options: "--valor-zero", ExecFunc: func() { executeSection("Valor zero") }},
		{Options: "--o-pacote-fmt", ExecFunc: func() { executeSection("O pacote fmt") }},
		{Options: "--criando-seu-proprio-tipo", ExecFunc: func() { executeSection("Criando seu próprio tipo") }},
		{Options: "--conversao-nao-coercao", ExecFunc: func() { executeSection("Conversão, não coerção") }},
	}
}

func HelpMeVariaveisValoresTipos() {

	hlp := []format.HelpMe{
		{Flag: "--go-playground", Description: "Exibe o Go Playground.", Width: 0},
		{Flag: "--hello-world", Description: "Exibe o Hello World.", Width: 0},
		{Flag: "--operador-curto-de-declaracao", Description: "Exibe o operador curto de declaração.", Width: 0},
		{Flag: "--a-palavra-reservada-var", Description: "Exibe a palavra reservada var.", Width: 0},
		{Flag: "--explorando-tipos", Description: "Exibe como explorar tipos.", Width: 0},
		{Flag: "--valor-zero", Description: "Exibe o valor zero.", Width: 0},
		{Flag: "--o-pacote-fmt", Description: "Exibe o pacote fmt.", Width: 0},
		{Flag: "--criando-seu-proprio-tipo", Description: "Exibe como criar seu próprio tipo.", Width: 0},
		{Flag: "--conversao-nao-coercao", Description: "Exibe a conversão não coerção.", Width: 0},
	}

	fmt.Println("\nCapítulo 2: Variáveis, Valores e Tipos")
	format.PrintHelpMe(hlp)
}

func executeSection(section string) {
	format.FormatSection(rootDir, section)
}
