/*
Package agrupamento_de_dados provides functionalities for demonstrating and executing
various data grouping topics in Go. It includes functions to print headers, execute
specific sections related to arrays, slices, and maps, and generate menu options and
help information for these topics. The package leverages the format package to format
and display the sections and help information.
*/
package agrupamento_de_dados

import (
	"fmt"

	"github.com/fabianoflorentino/aprendago/pkg/topic"
)

// AgrupamentoDeDados prints a header and executes a series of sections related to data grouping in Go.
// Each section is executed by calling the executeSection function with a specific topic name.
// The topics covered include arrays, slices (with various operations), and maps.
func Topics() {
	fmt.Printf("\nCapítulo 8: Agrupamento de Dados\n")

	content := topic.New()
	contentsAgrupamentoDeDados(rootDir, content)
}

// contentsAgrupamentoDeDados processes and organizes the contents of the
// "agrupamento de dados" (data grouping) topics. It initializes a new
// topic structure, retrieves the list of topics related to "agrupamento de dados",
// and populates the topics' contents from the specified root directory.
func contentsAgrupamentoDeDados(rootDir string, contents topic.ContentsProvider) {
	contents.TopicsContents(rootDir, listOfTopicsAgrupamentoDeDados())
}

// listOfTopicsAgrupamentoDeDados returns a slice of strings containing
// the topics related to "agrupamento de dados" (data grouping) in Go.
func listOfTopicsAgrupamentoDeDados() []string {
	list := []string{
		array,
		sliceLiteralComposta,
		sliceForRange,
		sliceFatiandoOuDeletando,
		sliceAnexando,
		sliceMake,
		sliceMultiDimensional,
		sliceSurpresaArraySubjacente,
		mapsIntroducao,
		mapsRangeEDeletando,
	}

	return listOfTopics(list)
}

// listOfTopics takes a slice of strings as input and returns a new slice of strings.
// The returned slice is initialized with a capacity of 10 and contains all the elements
// from the input slice.
func listOfTopics(inputList []string) []string {
	outputList := make([]string, 0, 10)
	outputList = append(outputList, inputList...)

	return outputList
}
