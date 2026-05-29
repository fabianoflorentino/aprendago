package cmd

import (
	"fmt"

	"github.com/fabianoflorentino/aprendago/internal/chapter"
)

// runOutline implements the 'outline' subcommand, showing all chapters
// and their topics.
func runOutline() error {
	all := chapter.All()
	if len(all) == 0 {
		fmt.Println("Nenhum capítulo disponível.")
		return nil
	}

	fmt.Println()
	for _, c := range all {
		topics, err := c.Topics()
		if err != nil {
			fmt.Printf("%s\n  (erro ao carregar tópicos)\n", c.Title)
			continue
		}
		fmt.Printf("%s\n", c.Title)
		for _, t := range topics {
			fmt.Printf("  %s\n", t)
		}
		fmt.Println()
	}

	return nil
}
