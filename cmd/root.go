package cmd

import (
	"fmt"
	"os"

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
	rootCmd.AddCommand(capsCmd, capCmd, outlineCmd, tuiCmd)

	rootCmd.SilenceUsage = true
	rootCmd.SilenceErrors = true
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Erro: %v\n", err)
		os.Exit(1)
	}
}

func rootRun(cmd *cobra.Command, args []string) error {
	if len(args) > 0 {
		return fmt.Errorf("comando desconhecido: %s. Use --help para ajuda", args[0])
	}
	return cmd.Help()
}
