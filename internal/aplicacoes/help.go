package aplicacoes

import (
	base "github.com/fabianoflorentino/aprendago/pkg/base_content"
	"github.com/fabianoflorentino/aprendago/pkg/format"
)

// HelpMeAplicacoes provides a list of help topics related to various Go packages and functionalities.
// It prints out a formatted help guide for each topic, including flags and descriptions.
// The topics cover documentation, JSON handling, interfaces, sorting, and encryption with bcrypt.
func Help() {
	h := []format.HelpMe{
		{Flag: flagDocumentacaoJSON, Description: descDocumentacaoJSON},
		{Flag: flagDocumentacaoJSONExampleJSONMarshal, Description: descDocumentacaoJSONExampleJSONMarshal},
		{Flag: flagDocumentacaoJSONExampleJSONUnmarshal, Description: descDocumentacaoJSONExampleJSONUnmarshal},
		{Flag: flagDocumentacaoJSONExampleJSONEncoder, Description: descDocumentacaoJSONExampleJSONEncoder},
		{Flag: flagJSONMarshal, Description: descJSONMarshal},
		{Flag: flagJSONUnmarshal, Description: descJSONUnmarshal},
		{Flag: flagInterfaceWriter, Description: descInterfaceWriter},
		{Flag: flagPacoteSort, Description: descPacoteSort},
		{Flag: flagPacoteSortExample, Description: descPacoteSortExample},
		{Flag: flagCustomizandoSort, Description: descCustomizandoSort},
		{Flag: flagCustomizandoSortExample, Description: descCustomizandoSortExample},
		{Flag: flagBcrypt, Description: descBcrypt},
	}

	b := base.New()
	b.HelpMe("Capítulo 16: Aplicações", h)
}
