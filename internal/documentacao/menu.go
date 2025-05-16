// Package documentacao provides functionalities to handle documentation topics
// for the "aprendago" project. It includes functions to display documentation
// sections, generate menu options for documentation topics, and print help
// information related to documentation.
package documentacao

import (
	base "github.com/fabianoflorentino/aprendago/pkg/base_content"
	"github.com/fabianoflorentino/aprendago/pkg/format"
)

// Menu creates and returns a slice of format.MenuOptions based on a predefined set of
// documentation-related flags and their corresponding execution functions. It initializes
// a new menu instance, associates each flag with a function that formats a documentation
// section, and returns the constructed menu options.
//
// The input parameter is a slice of strings, which is currently unused.
// The returned value is a slice of format.MenuOptions representing the menu entries.
func Menu([]string) []format.MenuOptions {
	m := base.New()

	newFlag := []string{
		flagIntroducaoDocumentacao,
		flagGoDoc,
		flagCommandGoDoc,
		flagPkgGoDev,
		flagEscrevendoDocumentacao,
	}

	newExecFunc := []func(){
		func() { m.SectionFormat(introducao, m.SectionFactory(rootDir)) },
		func() { m.SectionFormat(goDoc, m.SectionFactory(rootDir)) },
		func() { m.SectionFormat(commandGoDoc, m.SectionFactory(rootDir)) },
		func() { m.SectionFormat(pkgGoDev, m.SectionFactory(rootDir)) },
		func() { m.SectionFormat(escrevendoDoc, m.SectionFactory(rootDir)) },
	}

	return m.Menu(newFlag, newExecFunc)
}
