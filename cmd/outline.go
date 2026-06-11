package cmd

import (
	"fmt"

	"github.com/fabianoflorentino/aprendago/internal/chapter"
	"github.com/spf13/cobra"
)

var outlineCmd = &cobra.Command{
	Use:   "outline",
	Short: "Exibe o outline completo do curso",
	RunE: func(cmd *cobra.Command, args []string) error {
		return runOutline()
	},
}

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
