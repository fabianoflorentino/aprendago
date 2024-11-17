package ponteiros

import (
	"fmt"

	"github.com/fabianoflorentino/aprendago/pkg/format"
)

const rootDir = "internal/ponteiros"

func Ponteiros() {
	fmt.Printf("\n\nCapítulo 14: Ponteiros\n")

	executeSection("O que são ponteiros?")
	executeSection("Quando usar ponteiros")
}

func MenuPonteiros([]string) []format.MenuOptions {
	return []format.MenuOptions{
		{Options: "--o-que-sao-ponteiros", ExecFunc: func() { executeSection("O que são ponteiros?") }},
		{Options: "--quando-usar-ponteiros", ExecFunc: func() { executeSection("Quando usar ponteiros") }},
	}
}

func HelpMePonteiros() {
	hlp := []format.HelpMe{
		{Flag: "--o-que-sao-ponteiros", Description: "Descreve o que são ponteiros em Go", Width: 0},
		{Flag: "--quando-usar-ponteiros", Description: "Descreve quando usar ponteiros em Go", Width: 0},
	}

	fmt.Println("\nCapítulo 14: Ponteiros")
	format.PrintHelpMe(hlp)
}

func executeSection(section string) {
	format.FormatSection(rootDir, section)
}
