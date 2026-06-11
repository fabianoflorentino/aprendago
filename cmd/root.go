package cmd

import (
	"fmt"
	"os"

	"github.com/fabianoflorentino/aprendago/internal/chapter"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "aprendago",
	Short: "CLI para o curso Aprenda Go",
	Run: func(cmd *cobra.Command, args []string) {
		if err := rootRun(cmd, args); err != nil {
			fmt.Fprintf(os.Stderr, "Erro: %v\n", err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.Flags().Int("cap", 0, "Número do capítulo (usar com --topics ou --overview)")
	rootCmd.Flags().Bool("topics", false, "Lista tópicos do capítulo (usar com --cap)")
	rootCmd.Flags().Bool("overview", false, "Mostra conteúdo do capítulo (usar com --cap)")
	rootCmd.Flags().Bool("caps", false, "Lista capítulos disponíveis")
	rootCmd.Flags().Bool("outline", false, "Exibe o outline completo do curso")

	rootCmd.SilenceUsage = true
	rootCmd.SilenceErrors = true
}

func Execute() error {
	return rootCmd.Execute()
}

func rootRun(cmd *cobra.Command, args []string) error {
	if caps, _ := cmd.Flags().GetBool("caps"); caps {
		return runCaps()
	}
	if outline, _ := cmd.Flags().GetBool("outline"); outline {
		return runOutline()
	}
	if capNum, _ := cmd.Flags().GetInt("cap"); capNum > 0 {
		return runCapFlag(capNum, cmd)
	}
	if len(args) > 0 {
		switch args[0] {
		case "caps":
			return runCaps()
		case "cap":
			return runCap(args[1:])
		case "outline":
			return runOutline()
		}
		return fmt.Errorf("comando desconhecido: %s. Use --help para ajuda", args[0])
	}
	return cmd.Help()
}

func runCapFlag(num int, cmd *cobra.Command) error {
	ch := chapter.Get(num)
	if ch == nil {
		return fmt.Errorf("capítulo %d não encontrado", num)
	}
	topics, _ := cmd.Flags().GetBool("topics")
	if topics {
		return printChapterTopics(ch)
	}
	return ch.Overview()
}
