package cmd

import (
	"fmt"

	"github.com/fabianoflorentino/aprendago/internal/chapter"
	"github.com/spf13/cobra"
)

var capsCmd = &cobra.Command{
	Use:   "caps",
	Short: "Lista capítulos disponíveis",
	RunE: func(cmd *cobra.Command, args []string) error {
		return runCaps()
	},
}

func runCaps() error {
	fmt.Println("\nCapítulos do Curso")

	all := chapter.All()
	for _, c := range all {
		fmt.Printf("  --cap=%d --topics    %s\n", c.Number, c.Title)
	}

	return nil
}
