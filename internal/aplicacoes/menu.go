package aplicacoes

import (
	base "github.com/fabianoflorentino/aprendago/pkg/base_content"
	"github.com/fabianoflorentino/aprendago/pkg/format"
)

// MenuAplicacoes returns a slice of format.MenuOptions, each representing a different command-line option
// and its corresponding execution function. The options cover various topics such as JSON documentation,
// JSON marshal/unmarshal, JSON encoder, interface Writer, sort package, custom sorting, and bcrypt.
// Each option is associated with a specific ExecFunc that executes the relevant section or example.
func Menu([]string) []format.MenuOptions {
	m := base.New()

	newFlag := []string{
		flagDocumentacaoJSON,
		flagDocumentacaoJSONExampleJSONMarshal,
		flagDocumentacaoJSONExampleJSONUnmarshal,
		flagDocumentacaoJSONExampleJSONEncoder,
		flagJSONMarshal,
		flagJSONUnmarshal,
		flagInterfaceWriter,
		flagPacoteSort,
		flagPacoteSortExample,
		flagCustomizandoSort,
	}

	newExecFunc := []func(){
		func() { m.SectionFormat(documentacaoJSON, m.SectionFactory(rootDir)) },
		func() { m.SectionFormat(jsonMarshal, m.SectionFactory(rootDir)) },
		func() { m.SectionFormat(jsonUnmarshal, m.SectionFactory(rootDir)) },
		func() { m.SectionFormat(interfaceWriter, m.SectionFactory(rootDir)) },
		func() { m.SectionFormat(pacoteSort, m.SectionFactory(rootDir)) },
		func() { m.SectionFormat(customizandoSort, m.SectionFactory(rootDir)) },
		func() { m.SectionFormat(bcrypt, m.SectionFactory(rootDir)) },
	}

	return m.Menu(newFlag, newExecFunc)
}
