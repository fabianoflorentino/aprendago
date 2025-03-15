// Package aplicacoes provides various functionalities and examples related to JSON processing,
// sorting, and encryption in Go. It includes functions to demonstrate JSON marshaling and unmarshaling,
// using the Writer interface, sorting with the sort package, customizing sort operations, and using bcrypt for hashing.
package aplicacoes

import (
	"fmt"

	"github.com/fabianoflorentino/aprendago/pkg/format"
)

// rootDir represents the relative path to the "aplicacoes" directory within the internal package.
// It is used as a base directory for various application-specific operations and configurations.
const (
	rootDir = "internal/aplicacoes"
)

// Aplicacoes prints the chapter title and executes a series of sections
// related to various topics such as JSON handling, interfaces, sorting,
// and encryption using bcrypt.
func Aplicacoes() {
	fmt.Printf("\n\nCapítulo 16: Aplicações\n")
	showTopics()
}

// MenuAplicacoes returns a slice of format.MenuOptions, each representing a different command-line option
// and its corresponding execution function. The options cover various topics such as JSON documentation,
// JSON marshal/unmarshal, JSON encoder, interface Writer, sort package, custom sorting, and bcrypt.
// Each option is associated with a specific ExecFunc that executes the relevant section or example.
func MenuAplicacoes([]string) []format.MenuOptions {

	return []format.MenuOptions{
		{Options: "--documentacao-json", ExecFunc: func() { executeSections("Documentação JSON") }},
		{Options: "--documentacao-json --example --json-marshal", ExecFunc: func() { UsingJsonMarshal() }},
		{Options: "--documentacao-json --example --json-unmarshal", ExecFunc: func() { UsingJsonUnmarshal() }},
		{Options: "--documentacao-json --example --json-encoder", ExecFunc: func() { UsingJsonEncoder() }},
		{Options: "--json-marshal", ExecFunc: func() { executeSections("JSON marshal (ordenação)") }},
		{Options: "--json-unmarshal", ExecFunc: func() { executeSections("JSON unmarshal (desornação)") }},
		{Options: "--interface-writer", ExecFunc: func() { executeSections("A interface Writer") }},
		{Options: "--pacote-sort", ExecFunc: func() { executeSections("O pacote sort") }},
		{Options: "--pacote-sort --example", ExecFunc: func() { UsingPackageSort() }},
		{Options: "--customizando-sort", ExecFunc: func() { executeSections("Customizando sort") }},
		{Options: "--customizando-sort --example", ExecFunc: func() { UsingCustomSort() }},
		{Options: "--bcrypt", ExecFunc: func() { executeSections("bcrypt") }},
	}
}

// HelpMeAplicacoes provides a list of help topics related to various Go packages and functionalities.
// It prints out a formatted help guide for each topic, including flags and descriptions.
// The topics cover documentation, JSON handling, interfaces, sorting, and encryption with bcrypt.
func HelpMeAplicacoes() {
	hlp := []format.HelpMe{
		{Flag: "--documentacao-json", Description: "Descreve como documentar um pacote em Go", Width: 0},
		{Flag: "--documentacao-json --example --json-marshal", Description: "Exemplo de como ordenar um JSON", Width: 0},
		{Flag: "--documentacao-json --example --json-unmarshal", Description: "Exemplo de como desordenar um JSON", Width: 0},
		{Flag: "--documentacao-json --example --json-encoder", Description: "Exemplo de como usar o encoder JSON", Width: 0},
		{Flag: "--json-marshal", Description: "Descreve o pacote json.Marshal", Width: 0},
		{Flag: "--json-unmarshal", Description: "Descreve o pacote json.Unmarshal", Width: 0},
		{Flag: "--interface-writer", Description: "Descreve o que é a interface Writer", Width: 0},
		{Flag: "--pacote-sort", Description: "Descreve o pacote sort", Width: 0},
		{Flag: "--pacote-sort --example", Description: "Exemplo de como usar o pacote sort", Width: 0},
		{Flag: "--customizando-sort --example", Description: "Descreve como customizar o pacote sort", Width: 0},
		{Flag: "--bcrypt", Description: "Descreve o pacote bcrypt", Width: 0},
	}

	fmt.Printf("\nCapítulo 16: Aplicações\n")
	format.PrintHelpMe(hlp)
}

// showTopics iterates over a list of topics and executes sections for each topic.
// It retrieves the list of topics from the listOfTopics function and processes
// each topic using the executeSections function.
func showTopics() {
	for _, topic := range listOfTopics() {
		executeSections(topic)
	}
}

// listOfTopics returns a list of topics related to Go programming.
// The topics include JSON documentation, marshaling and unmarshaling,
// the Writer interface, the sort package, custom sorting, and bcrypt.
func listOfTopics() []string {
	return []string{
		"Documentação JSON",
		"JSON marshal (ordenação)",
		"JSON unmarshal (desornação)",
		"A interface Writer",
		"O pacote sort",
		"Customizando sort",
		"bcrypt",
	}
}

// executeSections processes a given section by formatting it.
// It takes a section name as a string parameter and applies the FormatSection
// function from the format package to the section, using the rootDir as the base directory.
//
// Parameters:
//   - section: The name of the section to be formatted.
func executeSections(section string) {
	format.FormatSection(rootDir, section)
}
