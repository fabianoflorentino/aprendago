// Package concorrencia provides functions and utilities to demonstrate and explain
// various concepts related to concurrency in Go. It includes sections on the
// differences between concurrency and parallelism, the use of goroutines and
// waitgroups, race conditions, mutexes, and atomic operations. The package also
// provides a menu for selecting different concurrency topics and a help function
// to describe the available options.
package concorrencia

import "github.com/fabianoflorentino/aprendago/pkg/topic"

// Topics initializes a new topic and constructs it with the specified parameters.
// It uses the TopicConstructor method to set up a topic with the root directory,
// a title ("Capítulo 18: Concorrência"), and a list of topics.
// This function is part of the concurrency-related chapter in the project.
func Topics() {
	t := topic.New()
	t.TopicConstructor(rootDir, "Capítulo 18: Concorrência", topics, t)
}
