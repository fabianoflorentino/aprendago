package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/fabianoflorentino/aprendago/internal/compat"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "aprendago",
	Short: "CLI para o curso Aprenda Go",
	// DisableFlagParsing is required during migration to prevent cobra from
	// rejecting legacy flags (--cap=N, --bem-vindo, etc.) that it doesn't know
	// about. We handle routing manually in rootRun.
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

// Execute is the entry point for the CLI, called from cmd/aprendago/main.go.
func Execute() error {
	return rootCmd.Execute()
}

// rootRun routes the CLI request:
//   - Empty args → legacy help
//   - "caps" → new caps subcommand
//   - "cap" → new cap subcommand
//   - "outline" → new outline subcommand
//   - Legacy flags (--cap=N, --help, etc.) → compat router → legacy fallback
func rootRun(cmd *cobra.Command, args []string) error {
	if len(args) == 0 {
		runLegacy(args)
		return nil
	}

	// Phase 1: New subcommands
	switch args[0] {
	case "caps":
		return runCaps()
	case "cap":
		return runCap(args[1:])
	case "outline":
		return runOutline()
	}

	// Phase 1+2: Legacy flags
	if strings.HasPrefix(args[0], "--") {
		if compat.Route(args) {
			return nil
		}
		runLegacy(args)
		return nil
	}

	return fmt.Errorf("comando desconhecido: %s. Use --help para ajuda", args[0])
}
