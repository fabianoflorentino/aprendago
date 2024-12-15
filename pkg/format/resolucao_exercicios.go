// Package format provides utilities for formatting and displaying various types of data.
// This package includes functions to format and print exercise resolutions.
package format

import "fmt"

// FormatResolucaoExercicios formats and prints the given resolution.
// It takes any type of resolution as input, converts it to a string using fmt.Sprintf,
// and then prints it to the standard output.
func FormatResolucaoExercicios(resolucao any) {
	fmt.Print("Resolução:\n")
	resolucao = fmt.Sprintf("%v", resolucao)

	fmt.Print(resolucao)
}
