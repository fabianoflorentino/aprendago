package format

import (
	"html/template"
	"os"

	"github.com/fabianoflorentino/aprendago/pkg/reader"
)

const documentTemplate = `
{{- .Document.Description.Name}}
{{range .Document.Description.Sections}}
{{- .Title}}

{{.Text -}}
{{end -}}
`

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
