package canais

import (
	base "github.com/fabianoflorentino/aprendago/pkg/base_content"
	"github.com/fabianoflorentino/aprendago/pkg/format"
)

// MenuCanais returns a slice of format.MenuOptions, each representing a menu option
// for different topics related to Go channels. Each menu option includes a string
// that represents the option and a function to execute when the option is selected.
// The topics covered include understanding channels, directional channels, range and close,
// select statement, the comma ok expression, convergence, divergence, and context.
func Menu([]string) []format.MenuOptions {
	m := base.New()

	newFlag := []string{
		flagEntendendoCanais,
		flagCanaisDirecionaisUtilizandoCanais,
		flagRangeEClose,
		flagSelectStatement,
		flagCommaOkExpression,
		flagConvergencia,
		flagConvergenciaExample,
		flagConvergenciaExampleChanString,
		flagDivergencia,
		flagDivergenciaExample,
		flagDivergenciaExampleWithFunc,
		flagContext,
	}

	newExecFunc := []func(){
		func() { m.SectionFormat(entendendoCanais, m.SectionFactory(rootDir)) },
		func() { m.SectionFormat(canaisDirecionaisUtilizandoCanais, m.SectionFactory(rootDir)) },
		func() { m.SectionFormat(rangeEClose, m.SectionFactory(rootDir)) },
		func() { m.SectionFormat(selectStatement, m.SectionFactory(rootDir)) },
		func() { m.SectionFormat(commaOkExpression, m.SectionFactory(rootDir)) },
		func() { m.SectionFormat(convergencia, m.SectionFactory(rootDir)) },
		func() { UsingConverge() },
		func() { UsingConvergeString() },
		func() { m.SectionFormat(divergencia, m.SectionFactory(rootDir)) },
		func() { UsingDivergence() },
		func() { UsingDivergenceWithFunc() },
		func() { m.SectionFormat(context, m.SectionFactory(rootDir)) },
	}

	return m.Menu(newFlag, newExecFunc)
}
