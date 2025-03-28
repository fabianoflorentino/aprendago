// Package menu provides functionality to display help and outline information
// for the "aprendago" application. It includes a function to show all available
// options and detailed help for each chapter of the course.
package menu

import (
	"fmt"
)

// HEADER is a constant string that provides usage instructions for running the Go program.
// It includes an example command and descriptions of available options such as --outline, --help, and --caps.
const HEADER = `
Uso: go run cmd/aprendago/main.go [opção]

Exemplo:
  go run cmd/aprendago/main.go --bem-vindo

Ajuda:

--outline  Exibe o outline completo do curso.
--help     Exibe a lista de todas as opções disponíveis.
--caps     Exibe a lista de capítulos disponíveis.
`

// HelpMe provides a comprehensive guide through various chapters and exercises
// of the Go programming course. It prints the header and sequentially calls
// helper functions for each chapter and exercise level, offering an overview
// and detailed explanations of variables, control flow, data grouping, structs,
// functions, pointers, concurrency, error handling, and more.
func HelpMe() {
	fmt.Printf("%s\n", HEADER)
	HelpMeCapituloOptions()
	HelpMeCapituloOutline()
}
