// Package canais provides functions to demonstrate and execute various concepts related to channels in Go.
// It includes sections on understanding channels, directional channels, range and close, select statement,
// comma ok idiom, convergence, divergence, and context. The package also provides menu options and help
// descriptions for each section.
package canais

import (
	"github.com/fabianoflorentino/aprendago/pkg/topic"
)

func Topics() {
	t := topic.New()
	t.TopicConstructor(rootDir, "Capítulo 21: Canais", topics, t)
}
