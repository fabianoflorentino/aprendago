package fluxo_de_controle

import (
	"fmt"

	helpme "github.com/fabianoflorentino/aprendago/outline/format"
)

func HelpMeFluxoDeControle() {
	hlp := []helpme.HelpMe{
		{Flag: "--entendendo-fluxo-de-controle", Description: "Explica o conceito de fluxo de controle em Go.", Width: 0},
		{Flag: "--loops-inicializacao-condicao-pos", Description: "Detalha o uso de loops com inicialização, condição e pós em Go.", Width: 0},
		{Flag: "--loops-nested-loop", Description: "Explora o conceito de loops aninhados em Go.", Width: 0},
		{Flag: "--loops-a-declaracao-for", Description: "Apresenta a declaração for em Go.", Width: 0},
		{Flag: "--loops-break-e-continue", Description: "Discute as instruções break e continue em loops em Go.", Width: 0},
		{Flag: "--loops-utilizando-ascii", Description: "Desafio surpresa: utilize ASCII para exibir texto em Go.", Width: 0},
		{Flag: "--loops-utilizando-ascii --resolucao", Description: "Desafio surpresa: utilize ASCII para exibir texto em Go.", Width: 0},
		{Flag: "--condicionais-a-declaracao-if", Description: "Apresenta a declaração if em Go.", Width: 0},
		{Flag: "--condicionais-if-else-if-else", Description: "Detalha a declaração if-else-if-else em Go.", Width: 0},
		{Flag: "--condicionais-a-declaracao-switch", Description: "Apresenta a declaração switch em Go.", Width: 0},
		{Flag: "--condicionais-a-declaracao-switch-pt2", Description: "Detalha a declaração switch em Go.", Width: 0},
		{Flag: "--operadores-logicos-condicionais", Description: "Explora os operadores lógicos condicionais em Go.", Width: 0},
	}

	fmt.Println("\nCapípulo 6: Fluxo de Controle")
	helpme.PrintHelpMe(hlp)
}
