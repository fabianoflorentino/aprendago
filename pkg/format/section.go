/*
Package format provides methods to format a document or a section of a document
*/
package format

import (
	"log"
	"os"

	"github.com/fabianoflorentino/aprendago/pkg/reader"
)

// FormatSection formats a section of a document
//   - dir: directory where the document is located
//   - title: title of the section to be formatted
//
// The method reads the section from the file and formats the document overview
// If the section is not found, the method logs an error message
// If an error occurs while reading the section, the method logs an error message
// If an error occurs while formatting the document, the method logs an error message
// and returns the error
func FormatSection(dir string, title string) {
	section, err := reader.ReadSection(dir, title)
	if err != nil {
		if os.IsNotExist(err) {
			log.Fatalf("Seção '%s' não encontrada no diretório '%s'", title, dir)
		} else {
			log.Fatalf("Erro ao ler a seção: %v", err)
		}
	}

	document := reader.Document{
		Description: reader.Description{
			Sections: []reader.Section{*section},
		},
	}

	err = FormatOverview([]reader.Document{document})
	if err != nil {
		log.Fatalf("Erro ao formatar o documento: %v", err)
	}
}
