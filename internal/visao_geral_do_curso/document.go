package visao_geral_do_curso

import (
	"log"

	"github.com/fabianoflorentino/aprendago/pkg/format"
	"github.com/fabianoflorentino/aprendago/pkg/reader"
)

const rootDir = "internal/visao_geral_do_curso"

func Document() {
	doc, err := reader.DocumentReader(rootDir)

	if err != nil {
		log.Fatalf("Error to read document: %v", err)
	}

	if len(doc) == 0 {
		log.Fatalf("No document found")
	}

	err = format.FormatDocument(doc)
	if err != nil {
		log.Fatalf("Error to format document: %v", err)
	}
}
