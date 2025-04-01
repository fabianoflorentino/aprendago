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

	list := []string{
		documentacaoJSON,
		jsonMarshal,
		jsonUnmarshal,
		interfaceWriter,
		pacoteSort,
		customizandoSort,
		bcrypt,
	}

	c := topic.New()
	c.TopicsContents(rootDir, c.ListOfTopics(list, 7))
}
