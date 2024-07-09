package fundamentos_da_programacao

import (
	"fmt"

	"github.com/fabianoflorentino/aprendago/pkg/format"
)

func HelpMeFundamentosDaProgramacao() {
	hlp := []format.HelpMe{
		{Flag: "--tipo-booleano", Description: "Explora o tipo de dados booleano em Go.", Width: 0},
		{Flag: "--como-os-computadores-funcionam", Description: "Descreve o funcionamento dos computadores e sua importância para a programação.", Width: 0},
		{Flag: "--tipos-numericos", Description: "Explora os tipos numéricos em Go.", Width: 0},
		{Flag: "--overflow", Description: "Discute o conceito de overflow e como ele pode afetar seu código.", Width: 0},
		{Flag: "--tipo-string", Description: "Explora o tipo de dados string em Go.", Width: 0},
		{Flag: "--sistemas-numericos", Description: "Apresenta os sistemas numéricos e sua importância para a programação.", Width: 0},
		{Flag: "--constantes", Description: "Detalha o uso de constantes em Go.", Width: 0},
		{Flag: "--iota", Description: "Explora o uso do identificador iota em Go.", Width: 0},
		{Flag: "--deslocamento-de-bits", Description: "Discute o conceito de deslocamento de bits em Go.", Width: 0},
	}

	fmt.Println("\nCapítulo 4: Fundamentos da Programação")
	format.PrintHelpMe(hlp)
}
