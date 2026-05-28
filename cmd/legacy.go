package cmd

import (
	"github.com/fabianoflorentino/aprendago/internal/menu"
)

// runLegacy delegates to the existing menu system for backwards compatibility.
// During the migration (Phase 0-2), all commands pass through here.
// In later phases, only unmigrated chapters will use this path.
func runLegacy(args []string) {
	if len(args) < 1 {
		menu.HelpMe()
		return
	}
	menu.Options(args)
}
