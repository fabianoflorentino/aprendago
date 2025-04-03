// Package canais provides functions to demonstrate and execute various concepts related to channels in Go.
// It includes sections on understanding channels, directional channels, range and close, select statement,
// comma ok idiom, convergence, divergence, and context. The package also provides menu options and help
// descriptions for each section.
package canais

import (
	"github.com/fabianoflorentino/aprendago/pkg/topic"
)

// Package canais provides functionality related to the management and construction
// of topics for learning Go (Golang). It is designed to help organize and structure
// content, such as chapters and topics, in a modular way.
//
// Topics initializes and constructs a new topic for "Capítulo 21: Canais" (Chapter 21: Channels).
// It utilizes the topic.New() constructor and sets up the topic with the provided
// root directory, title, and other parameters.
func Topics() {
	t := topic.New()
	t.TopicConstructor(rootDir, "Capítulo 21: Canais", topics, t)
}
