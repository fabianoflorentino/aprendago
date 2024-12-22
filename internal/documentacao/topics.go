package documentacao

import (
	"fmt"

	"github.com/fabianoflorentino/aprendago/pkg/format"
)

const (
	rootDir = "internal/documentacao"
)

func Documentacao() {
	fmt.Printf("\n\nCapítulo 25: Documentação\n")

	executeSections("Introdução")
	executeSections("go doc")
	executeSections("godoc")
	executeSections("https://pkg.go.dev/")
	executeSections("Escrevendo documentação")
}

func MenuDocumentacao([]string) []format.MenuOptions {
	return []format.MenuOptions{
		{Options: "--introducao", ExecFunc: func() { executeSections("Introdução") }},
		{Options: "--go-doc", ExecFunc: func() { executeSections("go doc") }},
		{Options: "--godoc", ExecFunc: func() { executeSections("godoc") }},
		{Options: "--pkg-go-dev", ExecFunc: func() { executeSections("https://pkg.go.dev/") }},
		{Options: "--escrevendo-documentacao", ExecFunc: func() { executeSections("Escrevendo documentação") }},
	}
}

func HelpMeDocumentacao() {
	hlp := []format.HelpMe{
		{Flag: "--introducao", Description: "Introdução"},
		{Flag: "--go-doc", Description: "go doc"},
		{Flag: "--godoc", Description: "godoc"},
		{Flag: "--pkg-go-dev", Description: "https://pkg.go.dev/"},
		{Flag: "--escrevendo-documentacao", Description: "Escrevendo documentação"},
	}

	fmt.Printf("\nCapítulo 25: Documentação")
	format.PrintHelpMe(hlp)
}

func executeSections(section string) {
	format.FormatSection(rootDir, section)
}
