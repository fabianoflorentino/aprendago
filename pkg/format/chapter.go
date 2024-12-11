package format

import (
	"fmt"

	"github.com/fabianoflorentino/aprendago/pkg/reader"
)

// PrintChapterHelpMe imprime o capítulo conforme o número recebido
func PrintChapterHelpMe(chapterNumber int) {
	chapter, err := reader.ReadChapterFile(chapterNumber)
	if err != nil {
		fmt.Printf("Erro ao tentar extrair dados do Capítulo %d:\n%s", chapterNumber, err)
		return
	}

	helpMes := []HelpMe{}

	for _, topic := range chapter.Topics() {
		for _, helpMe := range chapter.SerializeHelpMes(topic) {
			helpMes = append(helpMes, HelpMe{Flag: helpMe["flag"], Description: helpMe["span"], Width: 0})
		}
	}

	fmt.Printf("\nCapitulo %d: %s\n", chapterNumber, chapter.Title)
	PrintHelpMe(helpMes)
}
