package cmd

import (
	"fmt"
	"strconv"

	"github.com/fabianoflorentino/aprendago/internal/chapter"
	"github.com/spf13/cobra"
)

var capCmd = &cobra.Command{
	Use:   "cap <numero> [topics|overview|<topico>]",
	Short: "Acessa um capítulo do curso",
	Long: `Acessa um capítulo do curso Aprenda Go.

Exemplos:
  aprendago cap 8 topics
  aprendago cap 8 overview
  aprendago cap 8 "Agregação de fatias utilizando append"`,
	Args: cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		num, err := strconv.Atoi(args[0])
		if err != nil {
			return fmt.Errorf("número do capítulo inválido: %s", args[0])
		}

		ch := chapter.Get(num)
		if ch == nil {
			return fmt.Errorf("capítulo %d não encontrado. Use 'aprendago caps' para listar os capítulos disponíveis", num)
		}

		if len(args) > 1 {
			switch args[1] {
			case "topics":
				return printChapterTopics(ch)
			case "overview":
				return ch.Overview()
			default:
				return ch.ExecTopic(args[1])
			}
		}

		return ch.Overview()
	},
}

func printChapterTopics(ch *chapter.Chapter) error {
	topics, err := ch.Topics()
	if err != nil {
		return fmt.Errorf("erro ao listar tópicos do capítulo %d: %w", ch.Number, err)
	}

	fmt.Printf("\n%s\n", ch.Title)
	for _, t := range topics {
		fmt.Printf("  %s\n", t)
	}

	return nil
}
