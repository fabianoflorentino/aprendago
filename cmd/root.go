package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:                "aprendago",
	Short:              "CLI para o curso Aprenda Go",
	DisableFlagParsing: true,
	Run: func(cmd *cobra.Command, args []string) {
		runLegacy(args)
	},
}

func init() {
	rootCmd.SilenceUsage = true
	rootCmd.SilenceErrors = true
}

func Execute() error {
	return rootCmd.Execute()
}

