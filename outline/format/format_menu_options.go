package outline

import (
	"fmt"
	"strings"
)

// MenuOptions é uma struc que contém as opções de menu
type MenuOptions struct {
	Options  string
	ExecFunc func()
}

const headerMenu = `
Opção inválida

Use a opção --help para obter ajuda
`

// FormatMenuOptions formata as opções de menu e executa a função correspondente
// de acordo com os argumentos passados
func FormatMenuOptions(args []string, options []MenuOptions) {
	// Junta os argumentos separados por espaço em uma string única
	argStr := strings.Join(args, " ")

	// Verifica se a opção passada está contida nas opções disponíveis
	for _, opt := range options {
		if argStr == opt.Options {
			opt.ExecFunc()
			return
		}
	}

	fmt.Print(headerMenu)
}
