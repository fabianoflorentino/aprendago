// Package internal provides a collection of functions and modules that cover various topics
// and exercises related to learning the Go programming language. It includes modules for
// basic programming concepts, control flow, data grouping, functions, pointers, concurrency,
// and more. Each module is designed to help users understand and practice different aspects
// of Go programming through a series of exercises and examples.
package internal

import "github.com/fabianoflorentino/aprendago/internal/chapter"

// Outline provides an overview of the course by sequentially calling functions
// from various packages. Each function represents a different topic or set of
// exercises in the course.
func Outline() {
	for _, chapters := range chapter.New() {
		chapters()
	}
}
