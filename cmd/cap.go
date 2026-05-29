package cmd

import (
	"fmt"
	"strconv"

	"github.com/fabianoflorentino/aprendago/internal/chapter"
)

// runCap implements the 'cap' subcommand, accessing a specific chapter.
// Usage: aprendago cap <number> [topics|overview|<topic-name>]
func runCap(args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("uso: aprendago cap <numero> [topics|overview|<topico>]")
	}

	num, err := strconv.Atoi(args[0])
	if err != nil {
		return fmt.Errorf("número do capítulo inválido: %s", args[0])
	}

	ch := chapter.Get(num)
	if ch == nil {
		return fmt.Errorf("capítulo %d não encontrado. Use 'aprendago caps' para listar os capítulos disponíveis", num)
	}

	if len(args) < 2 {
		return ch.Overview()
	}

	switch args[1] {
	case "topics":
		return printChapterTopics(ch)
	case "overview":
		return ch.Overview()
	default:
		return ch.ExecTopic(args[1])
	}
}

// printChapterTopics prints the topic list for a given chapter.
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
