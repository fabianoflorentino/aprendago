// Package format provides functionality to format and print overviews of documents.
// It uses HTML templates to generate the output, which includes the name of the document
// and all its sections with their titles and text.
package format

import (
	"html/template"
	"io"
	"os"

	"github.com/fabianoflorentino/aprendago/pkg/reader"
)

// documentTemplate is a constant string that defines a template for formatting
// a document's description and its sections. The template uses the Go text/template
// package syntax to iterate over the sections of the document and format each
// section's title and text. The template expects a Document struct with a Description
// field that contains a Name and a slice of Sections, each having a Title and Text.
const documentTemplate = `
{{- .Document.Description.Name}}
{{range .Document.Description.Sections}}
{{- .Title}}

{{.Text -}}
{{end -}}
`

// FormatOverviewTo formats the given documents using the predefined template and
// writes the output to the provided writer. This allows callers to capture the
// formatted output (e.g. into a bytes.Buffer) instead of printing to stdout.
//
// Parameters:
//
//	w         - the writer to write the formatted output to.
//	documents - a slice of reader.Document to be formatted.
//
// Returns:
//
//	error - an error if there is an issue with template parsing or execution,
//	        otherwise nil.
func FormatOverviewTo(w io.Writer, documents []reader.Document) error {
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

		err = tmpl.Execute(w, context)
		if err != nil {
			return err
		}
	}

	return nil
}

// FormatOverview takes a slice of reader.Document and formats each document
// using a predefined template. It writes the formatted output to the standard
// output. If there is an error during template parsing or execution, it returns
// the error.
//
// Parameters:
//
//	documents - a slice of reader.Document to be formatted.
//
// Returns:
//
//	error - an error if there is an issue with template parsing or execution,
//	        otherwise nil.
func FormatOverview(documents []reader.Document) error {
	return FormatOverviewTo(os.Stdout, documents)
}
