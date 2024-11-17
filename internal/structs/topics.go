package structs

import (
	"fmt"

	"github.com/fabianoflorentino/aprendago/pkg/format"
)

const rootDir = "internal/structs"

func TopicStructs() {
	fmt.Print("\n\n10 - Structs\n")

	executeSection("Structs")
	executeSection("Structs embutidos")
	executeSection("Lendo a documentação")
	executeSection("Structs anônimos")
}

func MenuStructs([]string) []format.MenuOptions {
	return []format.MenuOptions{
		{Options: "--structs", ExecFunc: func() { executeSection("Structs") }},
		{Options: "--structs-embutidos", ExecFunc: func() { executeSection("Structs embutidos") }},
		{Options: "--lendo-a-documentacao", ExecFunc: func() { executeSection("Lendo a documentação") }},
		{Options: "--structs-anonimos", ExecFunc: func() { executeSection("Structs anônimos") }},
	}
}

func HelpMeStructs() {
	hlp := []format.HelpMe{
		{Flag: "--structs", Description: "Structs", Width: 0},
		{Flag: "--structs-embutidos", Description: "Structs Embutidos", Width: 0},
		{Flag: "--lendo-a-documentacao", Description: "Lendo a documentação", Width: 0},
		{Flag: "--structs-anonimos", Description: "Structs Anônimos", Width: 0},
	}

	fmt.Println("\n Capítulo 10: Structs")
	format.PrintHelpMe(hlp)
}

func executeSection(section string) {
	format.FormatSection(rootDir, section)
}
