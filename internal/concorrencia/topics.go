package concorrencia

import (
	"fmt"

	"github.com/fabianoflorentino/aprendago/pkg/format"
)

var rootDir = "internal/concorrencia"

func Concorrencia() {
	fmt.Printf("\n\n18 - Concorrência\n")

	executeSection("Concorrência vs Paralelismo")
	executeSection("Goroutines & WaitGroups")
	executeSection("Discussão: Condição de corrida")
	executeSection("Condição de corrida")
	executeSection("Mutex")
	executeSection("Atomic")
}

func MenuConcorrencia([]string) []format.MenuOptions {
	return []format.MenuOptions{
		{Options: "--concorrencia-vs-paralelismo", ExecFunc: func() { executeSection("Concorrência vs Paralelismo") }},
		{Options: "--goroutines-waitgroups", ExecFunc: func() { executeSection("Goroutines & WaitGroups") }},
		{Options: "--discussao-condicao-de-corrida", ExecFunc: func() { executeSection("Discussão: Condição de corrida") }},
		{Options: "--condicao-de-corrida", ExecFunc: func() { executeSection("Condição de corrida") }},
		{Options: "--mutex", ExecFunc: func() { executeSection("Mutex") }},
		{Options: "--atomic", ExecFunc: func() { executeSection("Atomic") }},
	}
}

func HelpMeConcorrencia() {
	hlp := []format.HelpMe{
		{Flag: "--concorrencia-vs-paralelismo", Description: "Apresenta a diferença entre concorrência e paralelismo.", Width: 0},
		{Flag: "--goroutines-waitgroups", Description: "Apresenta o uso de goroutines e waitgroups.", Width: 0},
		{Flag: "--discussao-condicao-de-corrida", Description: "Apresenta uma discussão sobre condição de corrida.", Width: 0},
		{Flag: "--condicao-de-corrida", Description: "Apresenta o conceito de condição de corrida.", Width: 0},
		{Flag: "--mutex", Description: "Apresenta o uso de mutex.", Width: 0},
		{Flag: "--atomic", Description: "Apresenta o uso de atomic.", Width: 0},
	}

	fmt.Println("\nCapítulo 18: Concorrência")
	format.PrintHelpMe(hlp)
}

func executeSection(section string) {
	format.FormatSection(rootDir, section)
}
