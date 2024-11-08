/*
Package format provides functions to format documents. The documents are printed to os.Stdout.
*/
package format

import (
	"html/template"
	"os"

	"github.com/fabianoflorentino/aprendago/pkg/reader"
)

// documentTemplate is a template for the overview of a document.
// It prints the name of the document and all its sections.
// The sections are printed with their title and text.
const documentTemplate = `
{{- .Document.Description.Name}}
{{range .Document.Description.Sections}}
{{- .Title}}

{{.Text -}}
{{end -}}
`

// FormatOverview prints an overview of the given documents to os.Stdout.
// The overview includes the name of the document and all its sections.
// The sections are printed with their title and text.
func FormatOverview(documents []reader.Document) error {
	tmpl, err := template.New("document").Parse(documentTemplate)
	if err != nil {
		return err
	}

	for _, doc := range documents {
		context := struct {
			reader.Document
			Divider string
		}{
			Document: doc,
		}

		err = tmpl.Execute(os.Stdout, context)
		if err != nil {
			return err
		}
	}

	return nil
}
