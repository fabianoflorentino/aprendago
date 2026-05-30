package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/fabianoflorentino/aprendago/internal/chapter"
	"github.com/fabianoflorentino/aprendago/internal/compat"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "aprendago",
	Short: "CLI para o curso Aprenda Go",
	// DisableFlagParsing is required to prevent cobra from rejecting legacy
	// flags (--cap=N, --help, etc.) that it doesn't know about.
	DisableFlagParsing: true,
	Run: func(cmd *cobra.Command, args []string) {
		if err := rootRun(cmd, args); err != nil {
			fmt.Fprintf(os.Stderr, "Erro: %v\n", err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.SilenceUsage = true
	rootCmd.SilenceErrors = true
}

func Execute() error {
	return rootCmd.Execute()
}

func rootRun(_ *cobra.Command, args []string) error {
	if len(args) == 0 {
		printHelp()
		return nil
	}

	switch args[0] {
	case "caps":
		return runCaps()
	case "cap":
		return runCap(args[1:])
	case "outline":
		return runOutline()
	}

	if strings.HasPrefix(args[0], "--") {
		switch args[0] {
		case "--help":
			printHelp()
			return nil
		case "--caps":
			return runCaps()
		case "--outline":
			return runOutline()
		default:
			if compat.Route(args) {
				return nil
			}
			return fmt.Errorf("opcao desconhecida: %s. Use --help para ajuda", args[0])
		}
	}

	return fmt.Errorf("comando desconhecido: %s. Use --help para ajuda", args[0])
}

func printHelp() {
	fmt.Println(`
Uso: aprendago [comando]

Comandos:
  caps                    Lista capitulos disponiveis
  cap <numero> [acao]     Acessa um capitulo
  outline                 Exibe o outline completo do curso

Acoes para cap:
  topics                  Lista topicos do capitulo
  overview                Mostra conteudo completo do capitulo
  <topico>                Mostra um topico especifico

Exemplos:
  aprendago caps
  aprendago cap 8 topics
  aprendago cap 8 overview
  aprendago cap 8 "Agregacao de fatias utilizando append"

Capitulos do Curso`)

	for _, c := range chapter.All() {
		fmt.Printf("  --cap=%d --topics    %s\n", c.Number, c.Title)
	}

	fmt.Println(`
Outline do Curso por Capitulo`)

	for _, c := range chapter.All() {
		fmt.Printf("  --cap=%d --overview    %s\n", c.Number, c.Title)
	}
}
