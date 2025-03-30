package aplicacoes

import (
	"github.com/fabianoflorentino/aprendago/pkg/format"
	"github.com/fabianoflorentino/aprendago/pkg/logger"
	"github.com/fabianoflorentino/aprendago/pkg/section"
)

// MenuAplicacoes returns a slice of format.MenuOptions, each representing a different command-line option
// and its corresponding execution function. The options cover various topics such as JSON documentation,
// JSON marshal/unmarshal, JSON encoder, interface Writer, sort package, custom sorting, and bcrypt.
// Each option is associated with a specific ExecFunc that executes the relevant section or example.
func Menu([]string) []format.MenuOptions {
	section, err := section.New(rootDir)
	if err != nil {
		logger.New("error to create a new section", err.Error()).Register()
	}

	return []format.MenuOptions{
		{Options: flagDocumentacaoJSON, ExecFunc: func() { sectionsAplicacoes(DocumentacaoJSON, section) }},
		{Options: flagDocumentacaoJSONExampleJSONMarshal, ExecFunc: func() { UsingJsonMarshal() }},
		{Options: flagDocumentacaoJSONExampleJSONUnmarshal, ExecFunc: func() { UsingJsonUnmarshal() }},
		{Options: flagDocumentacaoJSONExampleJSONEncoder, ExecFunc: func() { UsingJsonEncoder() }},
		{Options: flagJSONMarshal, ExecFunc: func() { sectionsAplicacoes(JSONMarshal, section) }},
		{Options: flagJSONUnmarshal, ExecFunc: func() { sectionsAplicacoes(JSONUnmarshal, section) }},
		{Options: flagInterfaceWriter, ExecFunc: func() { sectionsAplicacoes(InterfaceWriter, section) }},
		{Options: flagPacoteSort, ExecFunc: func() { sectionsAplicacoes(PacoteSort, section) }},
		{Options: flagPacoteSortExample, ExecFunc: func() { UsingPackageSort() }},
		{Options: flagCustomizandoSort, ExecFunc: func() { sectionsAplicacoes(CustomizandoSort, section) }},
		{Options: flagCustomizandoSortExample, ExecFunc: func() { UsingCustomSort() }},
		{Options: flagBcrypt, ExecFunc: func() { sectionsAplicacoes(Bcrypt, section) }},
	}
}

// sectionsAplicacoes formats and processes sections for a given title.
// It takes a title as a string and a SectionProvider interface to handle
// the formatting of the sections.
//
// Parameters:
//   - title: The title of the sections to be formatted.
//   - sections: An implementation of the SectionProvider interface
//     that provides the functionality to format the sections.
func sectionsAplicacoes(title string, sections section.SectionProvider) {
	sections.Format(title)
}
