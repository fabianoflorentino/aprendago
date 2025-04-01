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

	content := topic.New()
	content.TopicsContents(rootDir, content.ListOfTopics(list, len(list)))
}
