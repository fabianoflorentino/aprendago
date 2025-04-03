// Package concorrencia provides utilities and structures for working with
// concurrency-related topics in Go. It includes functionality for creating
// menus that explore various aspects of concurrency, such as the differences
// between concurrency and parallelism, the use of goroutines and waitgroups,
// handling race conditions, and synchronization mechanisms like mutexes and
// atomic operations. This package is designed to facilitate learning and
// experimentation with concurrency concepts in Go.
package concorrencia

import (
	base "github.com/fabianoflorentino/aprendago/pkg/base_content"
	"github.com/fabianoflorentino/aprendago/pkg/format"
)

// Menu generates a list of menu options for concurrency-related topics.
// It takes a slice of strings as input and returns a slice of formatted menu options.
// The function initializes a base menu object and defines flags and execution functions
// for various concurrency topics such as concurrency vs parallelism, goroutines and waitgroups,
// race conditions, mutexes, and atomic operations. These flags and functions are then passed
// to the base menu object to create the final menu structure.
func Menu([]string) []format.MenuOptions {
	m := base.New()

	newFlag := []string{
		flagConcorrenciaVsParalelismo,
		flagGoroutinesWaitgroups,
		flagDiscussaoCondicaoDeCorrida,
		flagCondicaoDeCorrida,
		flagMutex,
		flagAtomic,
	}

	newExecFunc := []func(){
		func() { m.SectionFormat(concorrenciaVsParalelismo, m.SectionFactory(rootDir)) },
		func() { m.SectionFormat(goroutinesWaitgroups, m.SectionFactory(rootDir)) },
		func() { m.SectionFormat(discussaoCondicaoDeCorrida, m.SectionFactory(rootDir)) },
		func() { m.SectionFormat(condicaoDeCorrida, m.SectionFactory(rootDir)) },
		func() { m.SectionFormat(mutex, m.SectionFactory(rootDir)) },
		func() { m.SectionFormat(atomic, m.SectionFactory(rootDir)) },
	}

	return m.Menu(newFlag, newExecFunc)
}
