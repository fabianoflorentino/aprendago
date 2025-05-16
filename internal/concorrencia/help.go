// Package concorrencia provides utilities and examples related to concurrency in Go.
// It includes functions and constants that demonstrate key concepts such as
// concurrency vs parallelism, goroutines, waitgroups, race conditions, mutexes,
// and atomic operations. This package is designed to support learning and
// understanding of concurrency topics covered in Chapter 18.
package concorrencia

import (
	base "github.com/fabianoflorentino/aprendago/pkg/base_content"
	"github.com/fabianoflorentino/aprendago/pkg/format"
)

// Help provides a summary of concurrency-related topics covered in Chapter 18.
// It creates a list of help items, each containing a flag and its corresponding description,
// and then uses the format package to display the help information in a structured format.
func Help() {
	h := []format.HelpMe{
		{Flag: flagConcorrenciaVsParalelismo, Description: descConcorrenciaVsParalelismo},
		{Flag: flagGoroutinesWaitgroups, Description: descGoroutinesWaitgroups},
		{Flag: flagDiscussaoCondicaoDeCorrida, Description: descDiscussaoCondicaoDeCorrida},
		{Flag: flagCondicaoDeCorrida, Description: descCondicaoDeCorrida},
		{Flag: flagMutex, Description: descMutex},
		{Flag: flagAtomic, Description: descAtomic},
	}

	b := base.New()
	b.HelpMe("Capítulo 18: Concorrência", h)
}
