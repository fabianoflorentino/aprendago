package cmd

import (
	"fmt"

	"github.com/fabianoflorentino/aprendago/internal/chapter"
)

// runCaps implements the 'caps' subcommand, listing all registered chapters.
func runCaps() error {
	fmt.Println("\nCapítulos do Curso")

	all := chapter.All()
	for _, c := range all {
		fmt.Printf("  --cap=%d --topics    %s\n", c.Number, c.Title)
	}

	return nil
}
