// Package aplicacoes provides various functionalities and examples related to JSON processing,
// sorting, and encryption in Go. It includes functions to demonstrate JSON marshaling and unmarshaling,
// using the Writer interface, sorting with the sort package, customizing sort operations, and using bcrypt for hashing.
package aplicacoes

import (
	"github.com/fabianoflorentino/aprendago/pkg/topic"
)

// Topics prints the chapter title and executes a series of sections
// related to various topics such as JSON handling, interfaces, sorting,
// and encryption using bcrypt.
func Topics() {
	t := topic.New()
	t.TopicConstructor(rootDir, "Capítulo 16: Aplicações", topics, t)
}
