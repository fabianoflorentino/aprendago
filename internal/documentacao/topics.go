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

}

func MenuDocumentacao([]string) []format.MenuOptions {
	return []format.MenuOptions{}
}

func HelpMeDocumentacao() {
	hlp := []format.HelpMe{}

	fmt.Printf("\nCapítulo 25: Documentação")
	format.PrintHelpMe(hlp)
}

func executeSections(section string) {
	format.FormatSection(rootDir, section)
}
