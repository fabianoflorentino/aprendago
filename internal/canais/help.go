package canais

import (
	base "github.com/fabianoflorentino/aprendago/pkg/base_content"
	"github.com/fabianoflorentino/aprendago/pkg/format"
)

// HelpMeCanais prints a list of topics related to channels in Go.
// It creates a slice of HelpMe structs, each containing a flag and a description
// of a specific topic about channels. The topics include understanding channels,
// directional channels, using channels, range and close, select statement,
// the comma ok idiom, convergence, divergence, and context.
// The function then prints the chapter title and calls PrintHelpMe to display the topics.
func Help() {
	h := []format.HelpMe{
		{Flag: flagEntendendoCanais, Description: descEntendendoCanais},
		{Flag: flagCanaisDirecionaisUtilizandoCanais, Description: descCanaisDirecionaisUtilizandoCanais},
		{Flag: flagRangeEClose, Description: descRangeEClose},
		{Flag: flagSelectStatement, Description: descSelectStatement},
		{Flag: flagCommaOkExpression, Description: descCommaOkExpression},
		{Flag: flagConvergencia, Description: descConvergencia},
		{Flag: flagConvergenciaExample, Description: descConvergenciaExample},
		{Flag: flagConvergenciaExampleChanString, Description: descConvergenciaExampleChanString},
		{Flag: flagDivergencia, Description: descDivergencia},
		{Flag: flagDivergenciaExample, Description: descDivergenciaExample},
		{Flag: flagDivergenciaExampleWithFunc, Description: descDivergenciaExampleWithFunc},
		{Flag: flagContext, Description: descContext},
	}

	b := base.New()
	b.HelpMe("Capítulo 21: Canais", h)
}
