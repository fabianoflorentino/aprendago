// Package format provides utilities for formatting and handling menu options.
// It includes structures and functions to manage and execute menu options based on user input.
package format

import (
	"fmt"
	"strings"
)

// MenuOptions represents a menu option with a description and an associated function to execute.
// Options is a string that describes the menu option.
// ExecFunc is a function that will be executed when the menu option is selected.
type MenuOptions struct {
	Options  string
	ExecFunc func()
}

// headerMenu is a constant string that contains a message indicating an invalid option
// and suggests using the --help option to get assistance.
const headerMenu = `
Opção inválida

Use a opção --help para obter ajuda
`

// FormatMenuOptions takes a slice of arguments and a slice of MenuOptions,
// concatenates the arguments into a single string, and checks if this string
// matches any of the available options. If a match is found, it executes the
// corresponding function. If no match is found, it prints the header menu.
//
// args: A slice of strings representing the command-line arguments.
// options: A slice of MenuOptions, each containing an option string and an
// associated function to execute if the option matches.
func FormatMenuOptions(args []string, options []MenuOptions) {
	// Joins the arguments separated by space into a single string
	argStr := strings.Join(args, " ")

	// Checks if the passed option is contained in the available options
	for _, opt := range options {
		if argStr == opt.Options {
			opt.ExecFunc()
			return
		}
	}

	fmt.Print(headerMenu)
}
