// Package documentacao provides functionalities to handle documentation topics
// for the "aprendago" project. It includes functions to display documentation
// sections, generate menu options for documentation topics, and print help
// information related to documentation.
package documentacao

import (
	"github.com/fabianoflorentino/aprendago/pkg/topic"
)

// Topics initializes and constructs the documentation topics for "Capítulo 25: Documentação".
// It creates a new topic instance and sets up the topics using the TopicConstructor function.
func Topics() {
	t := topic.New()
	t.TopicConstructor(rootDir, "Capítulo 25: Documentação", topics, t)
}
