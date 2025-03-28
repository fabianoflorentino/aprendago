// Package canais provides functions to demonstrate and execute various concepts related to channels in Go.
// It includes sections on understanding channels, directional channels, range and close, select statement,
// comma ok idiom, convergence, divergence, and context. The package also provides menu options and help
// descriptions for each section.
package canais

import (
	"fmt"

	"github.com/fabianoflorentino/aprendago/pkg/format"
	"github.com/fabianoflorentino/aprendago/pkg/logger"
	"github.com/fabianoflorentino/aprendago/pkg/section"
	"github.com/fabianoflorentino/aprendago/pkg/topic"
)

// rootDir is a constant that holds the relative path to the "canais" directory
// within the "internal" directory. This path is used as a base directory for
// various operations related to the "canais" module in the project.
const (
	rootDir = "internal/canais"
)

// entendendoCanais represents the topic "Entendendo Canais",
// which introduces the concept of channels in Go, explaining
// their purpose and how they facilitate communication between
// goroutines.
const (
	entendendoCanais                  string = "Entendendo Canais"
	canaisDirecionaisUtilizandoCanais string = "Canais direcionais & utilizando canais"
	rangeEClose                       string = "Range e Close"
	selectStatement                   string = "Select"
	commaOkExpression                 string = "A expressão comma ok"
	convergencia                      string = "Convergência"
	divergencia                       string = "Divergência"
	context                           string = "Context"
)

// listOfTopics is a slice of strings initialized with a capacity of 8.
// It is used to store a list of topics and starts as an empty slice.
var (
	listOfTopics []string = make([]string, 0, 8)
)

// Canais is a function that prints the title of Chapter 21: Channels
// and executes several sections related to the topic of channels in Go.
func Canais() {
	fmt.Printf("\n\nCapítulo 21: Canais\n")
	contentsCanais(rootDir)
}

// MenuCanais returns a slice of format.MenuOptions, each representing a menu option
// for different topics related to Go channels. Each menu option includes a string
// that represents the option and a function to execute when the option is selected.
// The topics covered include understanding channels, directional channels, range and close,
// select statement, the comma ok expression, convergence, divergence, and context.
func MenuCanais([]string) []format.MenuOptions {
	section, err := section.New(rootDir)
	if err != nil {
		logger.New("error to create a new section", err.Error()).Register()
	}

	return []format.MenuOptions{
		{Options: "--entendendo-canais", ExecFunc: func() { section.Format(entendendoCanais) }},
		{Options: "--canais-direcionais-utilizando-canais", ExecFunc: func() { section.Format(canaisDirecionaisUtilizandoCanais) }},
		{Options: "--range-e-close", ExecFunc: func() { section.Format(rangeEClose) }},
		{Options: "--select", ExecFunc: func() { section.Format(selectStatement) }},
		{Options: "--a-expressao-comma-ok", ExecFunc: func() { section.Format(commaOkExpression) }},
		{Options: "--convergencia", ExecFunc: func() { section.Format(convergencia) }},
		{Options: "--convergencia --example", ExecFunc: func() { UsingConverge() }},
		{Options: "--convergencia --example --chan-string", ExecFunc: func() { UsingConvergeString() }},
		{Options: "--divergencia", ExecFunc: func() { section.Format(divergencia) }},
		{Options: "--divergencia --example", ExecFunc: func() { UsingDivergence() }},
		{Options: "--divergencia --example --with-func", ExecFunc: func() { UsingDivergenceWithFunc() }},
		{Options: "--context", ExecFunc: func() { section.Format(context) }},
	}
}

// HelpMeCanais prints a list of topics related to channels in Go.
// It creates a slice of HelpMe structs, each containing a flag and a description
// of a specific topic about channels. The topics include understanding channels,
// directional channels, using channels, range and close, select statement,
// the comma ok idiom, convergence, divergence, and context.
// The function then prints the chapter title and calls PrintHelpMe to display the topics.
func HelpMeCanais() {
	hlp := []format.HelpMe{
		{Flag: "--entendendo-canais", Description: "Entendendo Canais"},
		{Flag: "--canais-direcionais-utilizando-canais", Description: "Canais direcionais & utilizando canais"},
		{Flag: "--range-e-close", Description: "Range e Close"},
		{Flag: "--select", Description: "Select"},
		{Flag: "--a-expressao-comma-ok", Description: "A expressão comma ok"},
		{Flag: "--convergencia", Description: "Convergência"},
		{Flag: "--convergencia --example", Description: "Convergência - Exemplo"},
		{Flag: "--convergencia --example --chan-string", Description: "Convergência de Strings - Exemplo"},
		{Flag: "--divergencia", Description: "Divergência"},
		{Flag: "--divergencia --example", Description: "Divergência - Exemplo"},
		{Flag: "--divergencia --example --with-func", Description: "Divergência - Exemplo com Função"},
		{Flag: "--context", Description: "Context"},
	}

	fmt.Println("\nCapítulo 21: Canais")
	format.PrintHelpMe(hlp)
}

// contentsCanais processes the contents of topics related to "canais".
// It initializes a new topic instance and retrieves the topics' contents
// from the specified root directory using a predefined list of topics.
//
// Parameters:
//   - rootDir: The root directory path where the topics' contents are located.
func contentsCanais(rootDir string) {
	contents := topic.New()
	contents.TopicsContents(rootDir, listOfTopicsCanais())
}

// listOfTopicsCanais returns a slice of strings containing the topics
// related to the "Canais" section. These topics include various concepts
// such as understanding channels, directional channels, range and close,
// select statement, comma-ok expression, convergence, divergence, and context.
func listOfTopicsCanais() []string {
	listOfTopics = append(listOfTopics,
		entendendoCanais,
		canaisDirecionaisUtilizandoCanais,
		rangeEClose,
		selectStatement,
		commaOkExpression,
		convergencia,
		divergencia,
		context,
	)

	return listOfTopics
}
