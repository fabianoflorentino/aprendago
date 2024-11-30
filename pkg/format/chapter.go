package format

import (
	"fmt"

	"github.com/fabianoflorentino/aprendago/pkg/reader"
)

// PrintChapterHelpMe imprime o capítulo conforme o número recebido
func PrintChapterHelpMe(chapterNumber int) {
	chapter, error := reader.ReadChapterFile(chapterNumber)
	if error != nil {
		fmt.Printf("Erro ao tentar extrair dados do Capítulo %s\n", chapterNumber)
		return
	}

	helpMes := []HelpMe{}

	for _, topic := range chapter.Topics {
		helpMe := HelpMe{Flag: topic.Flag, Description: topic.HelpMe, Width: 0}
		helpMes = append(helpMes, helpMe)
	}

	fmt.Printf("\nCapitulo %d: %s\n", chapterNumber, chapter.Title)
	PrintHelpMe(helpMes)
}
