package variaveis_valores_tipos

import (
	"fmt"

	"github.com/fabianoflorentino/aprendago/pkg/format"
)

const rootDir = "internal/variaveis_valores_tipos"

func VariaveisValoresETipos() {
	fmt.Printf("\n\n02 - Variáveis, Valores & Tipos\n")

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

func executeSection(section string) {
	format.FormatSection(rootDir, section)
}
