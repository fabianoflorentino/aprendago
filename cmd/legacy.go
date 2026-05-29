package cmd

import (
	"github.com/fabianoflorentino/aprendago/internal/compat"
	"github.com/fabianoflorentino/aprendago/internal/menu"
)

// runLegacy delegates to the existing menu system for backwards compatibility.
// During migration, unmigrated chapters and non-cobra flags (--bem-vindo, etc.)
// fall through to this function which uses the original menu.Options() parser.
//
// Phase 2+: Before falling back, the compat router intercepts --cap=N
// flags for chapters already migrated to the new system.
func runLegacy(args []string) {
	// Phase 2: Try compat router first for --cap=N flags
	if compat.Route(args) {
		return
	}

	if len(args) < 1 {
		menu.HelpMe()
		return
	}

	menu.Options(args)
}
