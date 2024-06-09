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
		// Separa as opções em um slice de strings
		optComponents := strings.Fields(opt.Options)

		// Verifica se a opção passada contém todas as opções disponíveis
		if contains(argStr, optComponents) {
			opt.ExecFunc()
			return
		}
	}

	fmt.Print(headerMenu)
}

// contains verifica se a string passada contém todas as opções disponíveis
func contains(argStr string, options []string) bool {
	// Verifica se a string passada contém todas as opções disponíveis
	for _, opt := range options {
		// Se a string não contém a opção, retorna falso
		if !strings.Contains(argStr, opt) {
			return false
		}
	}

	// Se a string contém todas as opções, retorna verdadeiro
	return true
}
