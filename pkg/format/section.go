package format

import (
	"log"
	"os"

	"github.com/fabianoflorentino/aprendago/pkg/reader"
)

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
			Chapter: reader.Chapter{
				Sections: []reader.Section{*section},
			},
		},
	}

	err = FormatOverview([]reader.Document{document})
	if err != nil {
		log.Fatalf("Erro ao formatar o documento: %v", err)
	}
}
