// Package aplicacoes provides various functionalities and examples related to JSON processing,
// sorting, and encryption in Go. It includes functions to demonstrate JSON marshaling and unmarshaling,
// using the Writer interface, sorting with the sort package, customizing sort operations, and using bcrypt for hashing.
package aplicacoes

import (
	"fmt"

	"github.com/fabianoflorentino/aprendago/pkg/topic"
)

// Topics prints the chapter title and executes a series of sections
// related to various topics such as JSON handling, interfaces, sorting,
// and encryption using bcrypt.
func Topics() {
	fmt.Printf("\nCapítulo 16: Aplicações\n")

	content := topic.New()
	contentsAplicacoes(rootDir, content)
}

// contentsAplicacoes initializes and populates the contents of applications topics.
// It creates a new instance of topics, retrieves the list of topics for applications,
// and processes their contents from the specified root directory.
//
// Parameters:
//   - rootDir: The root directory path where the topics' contents are located.
func contentsAplicacoes(rootDir string, contents topic.ContentsProvider) {
	contents.TopicsContents(rootDir, listOfTopicsAplicacoes())
}

// listOfTopicsAplicacoes returns a slice of strings containing a list of topics
// related to various Go programming concepts and functionalities. The topics
// include JSON handling, interface usage, sorting, and encryption with bcrypt.
func listOfTopicsAplicacoes() []string {
	list := []string{
		DocumentacaoJSON,
		JSONMarshal,
		JSONUnmarshal,
		InterfaceWriter,
		PacoteSort,
		CustomizandoSort,
		Bcrypt,
	}

	return listOfTopics(list)
}

func listOfTopics(inputList []string) []string {
	outputList := make([]string, 0, 7)
	outputList = append(outputList, inputList...)

	return outputList
}
