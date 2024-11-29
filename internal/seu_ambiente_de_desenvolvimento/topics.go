package seu_ambiente_de_desenvolvimento

import "github.com/fabianoflorentino/aprendago/pkg/format"

var rootDir = "internal/seu_ambiente_de_desenvolvimento"

func SeuAmbienteDeDesenvolvimento() {
	executeSection("O terminal")
}

func MenuSeuAmbienteDeDesenvolvimento([]string) []format.MenuOptions {
	return []format.MenuOptions{}
}

func HelpMeSeuAmbienteDeDesenvolvimento() {
	hlp := []format.HelpMe{}

	format.PrintHelpMe(hlp)
}

func executeSection(section string) {
	format.FormatSection(rootDir, section)
}
