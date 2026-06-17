package cmd

import (
	"github.com/fabianoflorentino/aprendago/internal/chapter"
	"github.com/fabianoflorentino/aprendago/internal/tui"
	"github.com/spf13/cobra"
)

var tuiCmd = &cobra.Command{
	Use:   "tui",
	Short: "Modo interativo com navegação visual do curso",
	Long: `Abre uma interface interativa no terminal com split view para navegar
pelos capítulos e tópicos do curso Aprenda Go.

Use as setas ↑↓ para navegar, Enter para selecionar, Esc para voltar
e / para buscar. Pressione q ou Ctrl+C para sair.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		chapters := chapter.All()
		return tui.Run(chapters)
	},
}
