package format

import (
	"fmt"
	"html/template"
	"os"
	"strings"

	"github.com/fabianoflorentino/aprendago/pkg/reader"
)

const documentDivider = "="

const documentTemplate = `
{{- .Divider}}
{{.Document.Description.Name}}
{{range .Document.Description.Chapter.Sections}}
{{.Title}}

{{.Text -}}
{{end}}
{{.Divider -}}
`

func buildDivider(document reader.Document) string {
	maxLength := len(document.Description.Name)
	lines := []string{document.Description.Name, document.Description.Name}
	for _, section := range document.Description.Chapter.Sections {
		lines = append(lines, section.Title)
		lines = append(lines, strings.Split(section.Text, "\n")...)
	}

	for _, l := range lines {
		if len(l) > maxLength {
			maxLength = len(l)
		}
	}

	return strings.Repeat(documentDivider, maxLength)
}

func FormatDocument(documents []reader.Document) error {
	tmpl, err := template.New("document").Parse(documentTemplate)
	if err != nil {
		return err
	}

	for _, doc := range documents {
		divider := buildDivider(doc)
		context := struct {
			reader.Document
			Divider string
		}{
			Document: doc,
			Divider:  divider,
		}

		err = tmpl.Execute(os.Stdout, context)
		if err != nil {
			return err
		}

		fmt.Println()
	}

	return nil
}
